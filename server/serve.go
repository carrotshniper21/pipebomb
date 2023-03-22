package server

import (
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Serve(r *mux.Router, port string) {
	// setup cors for router
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)
	http.ListenAndServe(":"+port, cors(r))
}
