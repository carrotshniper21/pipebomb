// pipebomb/server/server.go
package server

import (
	handies "github.com/gorilla/handlers"
	"log"
	"net/http"
	_ "pipebomb/docs"

	"github.com/gorilla/mux"
)

// Server sets up the server
func Server(r *mux.Router, port string) {
	// setup cors for router
	cors := handies.CORS(
		handies.AllowedOrigins([]string{"*"}),
		handies.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handies.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)
	log.Fatal(http.ListenAndServe(":"+port, cors(r)))
}

func StartServer(port string) {
	r := InitRouter()
	Server(r, port)
}
