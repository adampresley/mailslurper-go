// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"

	"github.com/adampresley/mailslurper/admin"
	"github.com/adampresley/mailslurper/data"
	"github.com/adampresley/mailslurper/settings"
	"github.com/adampresley/mailslurper/smtp"
)

func main() {
	settings.Config = settings.Configuration{
		WWW:         "www/",
		WWWPort:     8080,
		SmtpAddress: "127.0.0.1",
		SmtpPort:    8000,
	}

	err := settings.Config.LoadSettings("config.json")
	if err != nil {
		fmt.Printf("There was an error reading your config.json settings file: %s", err)
		return
	}

	/*
	 * Setup global database connection handle
	 */
	setupGlobalDatabaseConnection()
	defer data.Storage.Disconnect()

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
	setupAdminHandlers()
	fmt.Printf("MailSlurper administrator started on 0.0.0.0:%d (%s)\n\n", int(settings.Config.WWWPort), settings.Config.WWW)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", int(settings.Config.WWWPort)), nil)
}

func setupAdminHandlers() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(settings.Config.WWW))))
	http.HandleFunc("/mails", admin.GetMailCollection)
	http.HandleFunc("/ws", admin.WebsocketHandler)
}

func setupGlobalDatabaseConnection() {
	data.Storage = data.MailStorage{}
	err := data.Storage.Connect("./mail.db")

	if err != nil {
		panic("Unable to create database")
	}
}
