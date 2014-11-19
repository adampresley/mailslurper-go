// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package listener

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adampresley/mailslurper/mailslurperservice/controllers/mailController"
	"github.com/adampresley/mailslurper/mailslurperservice/controllers/versionController"
	"github.com/adampresley/mailslurper/mailslurperservice/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func buildUrl(version int, endpoint string) string {
	return fmt.Sprintf("/v%d%s", version, endpoint)
}

/*
Sets up and returns an HTTP server structure. This configures
middleware for logging and access control.
*/
func NewHttpListener(address string, port int) *http.Server {
	router := setupHttpRouter()

	server := alice.New(
		middleware.AccessControl,
		middleware.OptionsHandler,
		middleware.Logger).Then(router)

	listener := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", address, port),
		Handler: server,
	}

	return listener
}

/*
Sets the HTTP routes
*/
func setupHttpRouter() http.Handler {
	router := mux.NewRouter()

	/* Version */
	router.HandleFunc(buildUrl(1, "/version"), versionController.GetVersion_v1).Methods("GET", "OPTIONS")

	/* Mail and attachments */
	router.HandleFunc(buildUrl(1, "/mails/{mailId}"), mailController.GetMail_v1).Methods("GET", "OPTIONS")
	router.HandleFunc(buildUrl(1, "/mails/page/{pageNumber}"), mailController.GetMailCollection_v1).Methods("GET", "OPTIONS")
	router.HandleFunc(buildUrl(1, "/mails/{mailId}/attachments/{attachmentId}"), mailController.DownloadAttachment_v1).Methods("GET", "OPTIONS")

/*

	// Home
	requestRouter.HandleFunc("/", controllers.Home).Methods("GET")

	// Configuration
	requestRouter.HandleFunc("/configuration", controllers.Config).Methods("GET")
	requestRouter.HandleFunc("/config", controllers.GetConfig).Methods("GET")
	requestRouter.HandleFunc("/config", controllers.SaveConfig).Methods("PUT")

	// Web-sockets
	requestRouter.HandleFunc("/ws", smtp.WebsocketHandler)

	// Static requests
	requestRouter.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir(staticPath))))
*/
	return router
}

/*
Starts the HTTP listener and serves Service requests
*/
func StartHttpListener(httpListener *http.Server) error {
	log.Println("INFO - HTTP listener started on", httpListener.Addr)
	return httpListener.ListenAndServe()
}
