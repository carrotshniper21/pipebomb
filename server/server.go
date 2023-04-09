package server

import (
  "net/http"
  "pipebomb/handlers"
	handies "github.com/gorilla/handlers"

  "github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/", handlers.Home)
  r.HandleFunc("/api/films/vip/search", handlers.FilmSearch)
	// r.HandleFunc("/api/films/vip/source", handlers.FilmSource)
  return r
}

func Server(r *mux.Router, port string) {
	// setup cors for router
	cors := handies.CORS(
		handies.AllowedOrigins([]string{"*"}),
		handies.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
    handies.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)
	http.ListenAndServe(":"+port, cors(r))
}

func StartServer(port string) {
  r := InitRouter() 
  Server(r, port)
}
