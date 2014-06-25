// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"

	"bitbucket.org/kardianos/service"
	"github.com/adampresley/mailslurper/admin/controllers"
	"github.com/adampresley/mailslurper/profiling"
	"github.com/adampresley/mailslurper/settings"
	"github.com/adampresley/mailslurper/smtp"
	"github.com/gorilla/mux"
)

func main() {
	var err error

	profiling.Initialize()
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
	 * Prepare SIGINT handler (CTRL+C)
	 */
	profiling.Timer.Step("Initializing SIGINT channel")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		for _ = range done {
/*
			if *cpuprofile != "" {
				pprof.StopCPUProfile();
			}
*/

			log.Println("Shutting down...")

/*
			if *memprofile != "" {
				memFile, err := os.Create(*memprofile)
				if err != nil {
					log.Fatal(err)
				}

				pprof.WriteHeapProfile(memFile)
				memFile.Close()
			}
*/

			os.Exit(0)
		}
	}()

	profiling.Timer.Step("Loading settings")

	settings.Config = settings.Configuration{
		WWW:         "www/",
		WWWPort:     8080,
		SmtpAddress: "127.0.0.1",
		SmtpPort:    8000,
	}

	settings.Config.LoadHeader("header")
	settings.Config.LoadFooter("footer")

	err = settings.Config.LoadSettings("config.json")
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
	profiling.Timer.Step("Setup database storage and write-listener")

	setupGlobalDatabaseConnection()
	defer smtp.Storage.Disconnect()

	/*
	 * Setup the SMTP listener
	 */
	profiling.Timer.Step("Setup SMTP server")
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
	profiling.Timer.Step("Setup HTTP administrator")
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

	profiling.Timer.Step("Serving requests")
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

	case "mssql":
		engine = smtp.ENGINE_MSSQL
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
		log.Panic("Unable to connect to database: ", err)
	}
}
