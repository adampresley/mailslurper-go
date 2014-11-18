// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// LOOK ITO USING https://github.com/GeertJohan/go.rice
// TO EMBED ASSETS IN A SINGLE EXE!!!
package main

import (
//	"fmt"
	"log"
//	"net/http"
	"os"
//	"os/signal"
//	"path/filepath"
	"runtime"

	"github.com/adampresley/mailslurper/libmailslurper/configuration"
	"github.com/adampresley/mailslurper/libmailslurper/storage"
	"github.com/adampresley/mailslurper/mailslurperservice/listener"

	"github.com/adampresley/sigint"
)

func main() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
	 * Prepare SIGINT handler (CTRL+C)
	 */
	sigint.ListenForSIGINT(func() {
		log.Println("Shutting down...")
		os.Exit(0)
	})

	/*
	 * Load configuration
	 */
	config, err := configuration.LoadConfigurationFromFile(configuration.CONFIGURATION_FILE_NAME)
	if err != nil {
		log.Println("ERROR - There was an error reading your configuration file: ", err)
		os.Exit(0)
	}

	/*
	 * Setup global database connection handle
	 */
	databaseConnection := config.GetDatabaseConfiguration()
	err = storage.ConnectToStorage(databaseConnection)
	if err != nil {
		log.Println("ERROR - There was an error connecting to your data storage: ", err)
		os.Exit(0)
	}

	defer storage.DisconnectFromStorage()

	/*
	 * Setup the SMTP listener
	smtpServer := smtp.Server{Address: fmt.Sprintf("%s:%d", settings.Config.SmtpAddress, int(settings.Config.SmtpPort))}
	defer smtpServer.Close()
	 */

	/*
	 * Start up the SMTP server and serve requests
	 * out of a goroutine.
	smtpServer.Connect()
	go smtpServer.ProcessRequests()
	 */

	/*
	 * Setup web server for the administrator
	 */
/*	profiling.Timer.Step("Setup HTTP administrator")
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
*/

	/*
	 * Start the services server
	 */
	log.Println("MailSlurper started")
	listener.StartHttpListener(listener.NewHttpListener("0.0.0.0", 8085))
}

