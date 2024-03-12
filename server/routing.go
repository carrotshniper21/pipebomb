// pipebomb/server/routing.go
package server

import (
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"pipebomb/cache"
	"pipebomb/handlers"
	"pipebomb/logging"
)

func InitRouter(redisCache *cache.RedisCache) *mux.Router {
	color.Blue("Starting server on port http://127.0.0.1:8001")
	r := mux.NewRouter()
	r.Use(logging.LoggingMiddleware)
	r.HandleFunc("/", handlers.Home)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/api/films/vip/search", handlers.FilmSearch(redisCache)).Methods("GET") 
	r.HandleFunc("/api/films/vip/servers", handlers.FetchFilmServers(redisCache)).Methods("GET")
	r.HandleFunc("/api/films/vip/sources", handlers.FetchFilmSources(redisCache)).Methods("GET")
	r.HandleFunc("/api/series/vip/search", handlers.ShowSearch(redisCache)).Methods("GET")
	r.HandleFunc("/api/series/vip/seasons", handlers.ShowSeason(redisCache)).Methods("GET")
	r.HandleFunc("/api/series/vip/servers", handlers.FetchShowServers(redisCache)).Methods("GET")
	r.HandleFunc("/api/series/vip/sources", handlers.FetchShowSources(redisCache)).Methods("GET")
	r.HandleFunc("/api/anime/all/search", handlers.AnimeSearch(redisCache)).Methods("GET")
	r.HandleFunc("/api/anime/all/sources", handlers.FetchAnimeSources(redisCache)).Methods("GET")
	r.HandleFunc("/api/novels/rln/search", handlers.NovelSearch(redisCache)).Methods("GET")

	return r
}
