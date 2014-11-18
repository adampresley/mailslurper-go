package listener

import (
	"fmt"
	"net/http"

	"github.com/adampresley/mailslurper/mailslurperservice/controllers/versionController"
	"github.com/adampresley/mailslurper/mailslurperservice/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

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

func setupHttpRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/version", versionController.GetVersion).Methods("GET", "OPTIONS")

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
	return router
}

func StartHttpListener(httpListener *http.Server) error {
	return httpListener.ListenAndServe()
}
