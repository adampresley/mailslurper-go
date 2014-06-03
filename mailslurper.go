// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adampresley/mailslurper/admin/controllers"
	"github.com/adampresley/mailslurper/settings"
	"github.com/adampresley/mailslurper/smtp"
	"github.com/gorilla/mux"
)

func main() {
	settings.Config = settings.Configuration{
		WWW:         "www/",
		WWWPort:     8080,
		SmtpAddress: "127.0.0.1",
		SmtpPort:    8000,
	}

	settings.Config.LoadHeader("header")
	settings.Config.LoadFooter("footer")

	err := settings.Config.LoadSettings("config.json")
	if err != nil {
		log.Println("There was an error reading your config.json settings file: ", err)
		return
	}

	wwwAbs, _ := filepath.Abs(settings.Config.WWW)
	settings.Config.WWWAbs = wwwAbs
	staticPath := filepath.Join(settings.Config.WWWAbs, "resources")

	/*
	 * Setup global database connection handle
	 */
	setupGlobalDatabaseConnection()
	defer smtp.Storage.Disconnect()

	/*
	 * Setup the SMTP listener
	 */
	smtpServer := smtp.Server{Address: fmt.Sprintf("%s:%d", settings.Config.SmtpAddress, int(settings.Config.SmtpPort))}
	defer smtpServer.Close()

	/*
	 * Start up the SMTP server and serve requests
	 * out of a goroutine.
	 */
	smtpServer.Connect()
	go smtpServer.ProcessRequests()

	/*
	 * Setup web server for the administrator
	 */
	requestRouter := mux.NewRouter()

	// Home
	requestRouter.HandleFunc("/", controllers.Home).Methods("GET")

	// Mail items
	requestRouter.HandleFunc("/mail", controllers.GetMailItem).Methods("GET")
	requestRouter.HandleFunc("/mails", controllers.GetMailCollection).Methods("GET")
	requestRouter.HandleFunc("/attachment", controllers.DownloadAttachment).Methods("GET")

	// Configuration
	requestRouter.HandleFunc("/configuration", controllers.Config).Methods("GET")
	requestRouter.HandleFunc("/config", controllers.GetConfig).Methods("GET")
	requestRouter.HandleFunc("/config", controllers.SaveConfig).Methods("PUT")

	// Web-sockets
	requestRouter.HandleFunc("/ws", smtp.WebsocketHandler)

	// Static requests
	requestRouter.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir(staticPath))))

	log.Printf("MailSlurper administrator started on %s (%s)\n\n", settings.Config.GetFullListenAddress(), settings.Config.WWW)
	http.ListenAndServe(settings.Config.GetFullListenAddress(), requestRouter)
}

func setupGlobalDatabaseConnection() {
	var engine int
	var host string
	var port string
	var database string
	var userName string
	var password string

	switch settings.Config.DBEngine {
	case "sqlite":
		engine = smtp.ENGINE_SQLITE

	case "mysql":
		engine = smtp.ENGINE_MYSQL
		host = settings.Config.DBHost
		port = settings.Config.DBPort
		database = settings.Config.DBDatabase
		userName = settings.Config.DBUserName
		password = settings.Config.DBPassword
	}

	smtp.Storage = smtp.MailStorage{
		Engine:   engine,
		Host:     host,
		Port:     port,
		Database: database,
		UserName: userName,
		Password: password,
	}

	err := smtp.Storage.Connect()

	if err != nil {
		panic("Unable to connect to database")
	}
}
