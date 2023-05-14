// pipebomb/server/routing.go
package server

import (
	"pipebomb/handlers"
	"pipebomb/logging"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitRouter initializes the router
func InitRouter() *mux.Router {
	color.Blue("Starting server on port http://127.0.0.1:8001")
	r := mux.NewRouter()
	r.Use(logging.LoggingMiddleware)
	r.HandleFunc("/", handlers.Home)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/api/films/vip/search", handlers.FilmSearch)
	r.HandleFunc("/api/films/vip/servers", handlers.FetchFilms)
	r.HandleFunc("/api/films/vip/sources", handlers.FetchFilmSources)
	r.HandleFunc("/api/series/vip/search", handlers.ShowSearch)
	r.HandleFunc("/api/series/vip/seasons", handlers.ShowSeason)
	r.HandleFunc("/api/series/vip/servers", handlers.FetchShows)
	r.HandleFunc("/api/series/vip/sources", handlers.FetchShowSources)
	r.HandleFunc("/api/anime/all/search", handlers.AnimeSearch)
	r.HandleFunc("/api/anime/all/sources", handlers.FetchAnimeSources)
  r.HandleFunc("/api/profiles/users", handlers.GetUsers).Methods("GET")
  r.HandleFunc("/api/profiles/users", handlers.CreateUser).Methods("POST")
  r.HandleFunc("/api/profiles/users/{username}", handlers.GetUser).Methods("GET")
  r.HandleFunc("/api/profiles/users/{username}", handlers.UpdateUser).Methods("PUT")
  r.HandleFunc("/api/profiles/users/{username}", handlers.DeleteUser).Methods("DELETE")

	return r
}
