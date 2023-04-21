package server

// InitRouter initializes the router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/api/films/vip/search", handlers.FilmSearch)
	r.HandleFunc("/api/films/vip/servers", handlers.FetchFilms)
	r.HandleFunc("/api/films/vip/sources", handlers.FetchFilmSources)
	r.HandleFunc("/api/shows/vip/search", handlers.ShowSearch)
	r.HandleFunc("/api/shows/vip/seasons", handlers.ShowSeason)
	r.HandleFunc("/api/shows/vip/servers", handlers.FetchShows)
	r.HandleFunc("/api/shows/vip/sources", handlers.FetchShowSources)

	return r
}
