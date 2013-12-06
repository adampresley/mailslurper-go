// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/adampresley/mailslurper/admin"
	"github.com/adampresley/mailslurper/data"
	"github.com/adampresley/mailslurper/smtp"
)

var flagWWW = flag.String("www", "www/", "Path to the web administrator directory. Defaults to 'www/'")
var flagSmtpPort = flag.String("smtpport", "8000", "Port number to bind to for SMTP server. Defaults to 8000")
var flagWwwPort = flag.String("wwwport", "8080", "Port number to bind to for WWW administrator. Defaults to 8080")

func main() {
	flag.Parse()

	/*
	 * Setup global database connection handle
	 */
	setupGlobalDatabaseConnection()
	defer data.Storage.Disconnect()

	/*
	 * Setup the SMTP listener
	 */
	smtpServer := smtp.Server{Address: "127.0.0.1:" + *flagSmtpPort}
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
	fmt.Printf("MailSlurper administrator started on localhost:%s (%s)\n\n", *flagWwwPort, *flagWWW)
	http.ListenAndServe("0.0.0.0:"+*flagWwwPort, nil)
}

func setupAdminHandlers() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*flagWWW))))
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
