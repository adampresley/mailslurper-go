package listener

import (
	"fmt"
	"net/http"

	"github.com/adampresley/mailslurper/mailslurperservice/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func NewHttpListener(address string, port int) *http.Server {
	router := setupHttpRouter()

	server := alice.New(
		middleware.AccessControl,
		middleware.OptionsHandler,
		middleware.Logger
	).Then(router)

	listener := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", address, port),
		Handler: server,
	}

	return listener
}

func setupHttpRouter() http.ServerMux {
	router := mux.NewRouter()

	router.HandleFunc("/version", versionController.GetVersion).Methods("GET", "OPTIONS")

	return router
}

func StartHttpListener(httpListener *http.Server) error {
	return httpListener.ListenAndServe()
}
