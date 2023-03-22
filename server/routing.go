package server

import (
	"github.com/gorilla/mux"
	"pipebomb/handlers"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/films", handlers.GetFilm).Methods("GET")
	r.HandleFunc("/films/filmID", handlers.GetFilmSources).Methods("GET")

	return r
}
