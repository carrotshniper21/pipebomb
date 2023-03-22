// @title Pip API
// @description Pip API is a RESTful API designed for films and shows. 🎥
// @description You can retrieve film/show data by making API calls to endpoints.
// @description Users will be able to fetch show and film data, as well as stream sources.
// @version 0.0.1
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Pip API
// @description Pip API is a RESTful API designed for films and shows. 🎥
// @description You can retrieve film/show data by making API calls to endpoints.
// @description Users will be able to fetch show and film data, as well as stream sources.
// @version 0.0.1
// @basePath /
func main() {
	r := mux.NewRouter()

	// Enable CORS middleware
	c := cors.Default().Handler(r)

	// Define routes
	r.HandleFunc("/", getRoutes).Methods("GET")
	r.HandleFunc("/anime/search_anime", searchAnime).Methods("POST")
	r.HandleFunc("/anime/get_episodes/", getEpisodes).Methods("POST")
	r.HandleFunc("/anime/get_episode_url/", getEpisodeUrl).Methods("POST")
	r.HandleFunc("/films", getFilms).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/films/{filmID}", getSources).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/shows", getShows).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/shows/{showID}", getShowInfo).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/shows/episodeID", getEpisodeId).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/lightnovel/search_novels", searchNovels).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/lightnovel/novel_info", getNovelInfo).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/lightnovel/chapters", chooseNovelChapter).Methods("GET").Queries("q", "{q}")
	r.HandleFunc("/lightnovel/chapter_content", getChapter).Methods("GET").Queries("q", "{q}")

	// Serve Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", c))
}

// @Summary Get Routes
// @Description Get all available API routes
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func getRoutes(w http.ResponseWriter, r *http.Request) {
	routes := mux.Vars(r)
	json.NewEncoder(w).Encode(routes)
}

// searchAnime accepts a search query and returns a list of anime
// @Summary Search Anime
// @Description Search for anime
// @Accept json
// @Produce json
// @Param query body AnimeSearch true "Anime Search Request Body"
// @Success 200 {array} AnimeSearchResult
// @Failure 422 {object} HTTPValidationError
// @Router /anime/search_anime [post]
func searchAnime(w http.ResponseWriter, r *http.Request) {

}

// @Summary Get Episodes
// @Description Get episodes for anime
// @Accept json
// @Produce json
// @Param body body AnimeId true "Anime Id Request Body"
// @Success 200 {object} Episodes
// @Failure 422 {object} HTTPValidationError
// @Router /anime/get_episodes/ [post]
func getEpisodes(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Episode Url
// @Description Get streaming URL for episode
// @Accept json
// @Produce json
// @Param body body EpisodeInfo true "Episode Info Request Body"
// @Success 200 {array} AnimeEpResult
// @Failure 422 {object} HTTPValidationError
// @Router /anime/get_episode_url/ [post]
func getEpisodeUrl(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Film
// @Description Get film information
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} FilmModel
// @Failure 422 {object} HTTPValidationError
// @Router /films [get]
func getFilms(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Sources
// @Description Get sources for a film
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /films/{filmID} [get]
func getSources(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Shows
// @Description Get shows information
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} FilmModel
// @Failure 422 {object} HTTPValidationError
// @Router /shows [get]
func getShows(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Show Info
// @Description Get show information
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /shows/{showID} [get]
func getShowInfo(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Episode Id
// @Description Get episode id for show
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /shows/episodeID [get]
func getEpisodeId(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Search Novels
// @Description Search for novels
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /lightnovel/search_novels [get]
func searchNovels(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Novel Info
// @Description Get information for a novel
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /lightnovel/novel_info [get]
func getNovelInfo(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Choose Novel Chapter
// @Description Choose a chapter for a novel
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /lightnovel/chapters [get]
func chooseNovelChapter(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// @Summary Get Chapter
// @Description Get content for a chapter
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} map[string]interface{}
// @Failure 422 {object} HTTPValidationError
// @Router /lightnovel/chapter_content [get]
func getChapter(w http.ResponseWriter, r *http.Request) {
	// handle request
}

// Initialize router and routes
func initializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getRoutes).Methods("GET")
	router.HandleFunc("/anime/search_anime", searchAnime).Methods("POST")
	router.HandleFunc("/anime/get_episodes/", getEpisodes).Methods("POST")
	router.HandleFunc("/anime/get_episode_url/", getEpisodeUrl).Methods("POST")
	router.HandleFunc("/films", getFilms).Methods("GET")
	router.HandleFunc("/films/{filmID}", getSources).Methods("GET")
	router.HandleFunc("/shows", getShows).Methods("GET")
	router.HandleFunc("/shows/{showID}", getShowInfo).Methods("GET")
	router.HandleFunc("/shows/episodeID", getEpisodeId).Methods("GET")
	router.HandleFunc("/light-novel/search_novels", searchNovels).Methods("GET")
	router.HandleFunc("/light-novel/novel_info", getNovelInfo).Methods("GET")
	router.HandleFunc("/light-novel/chapters", chooseNovelChapter).Methods("GET")
	router.HandleFunc("/light-novel/chapter_content", getChapter).Methods("GET")

	// Swagger docs
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}
