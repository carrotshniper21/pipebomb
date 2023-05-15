// pipebomb/handlers/handlers.go
package handlers

import (
	"encoding/json"
	"net/http"
	"pipebomb/util"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"pipebomb/anime"
	"pipebomb/film"
	"pipebomb/novel"
	"pipebomb/profiles"
	"pipebomb/show"
)

// ######## Films ########

// FilmSearch
// @Summary		Search for films
// @Description	Search for films by query
// @Tags			Films
// @Accept			json
// @Produce		json
// @Param			q	query		string	true	"Search Query"
// @Success		200	{object}	film.FilmSearch
// @Router			/films/vip/search [get]
func FilmSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results, err := film.ProcessQuery(query)

	util.HandleError(w, err, "Error handling (film) query", http.StatusInternalServerError)

	jsonResponse := map[string]interface{}{
		"results": results,
	}

	util.WriteJSONResponse(w, jsonResponse)
}

// FetchFilmServers
// @Summary		Fetch film servers
// @Description	Fetch film servers by film ID
// @Tags			Films
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"Film ID"
// @Success		200	{array}	film.FilmServer
// @Router			/films/vip/servers [get]
func FetchFilmServers(w http.ResponseWriter, r *http.Request) {
	filmID := r.URL.Query().Get("id")
	servers, err := film.GetFilmServer(filmID)

	util.HandleError(w, err, "Error fetching (film) servers", http.StatusInternalServerError)

	jsonResponse := map[string]interface{}{
		"servers": servers,
	}

	util.WriteJSONResponse(w, jsonResponse)
}

// FetchFilmSources
// @Summary		Fetch film sources
// @Description	Fetch film servers by server ID
// @Tags			Films
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"Server ID"
// @Success		200	{array}	film.FilmSourcesEncrypted
// @Router			/films/vip/sources [get]
func FetchFilmSources(w http.ResponseWriter, r *http.Request) {
	serverId := r.URL.Query().Get("id")
	sources, err := film.GetFilmSources(serverId)

	util.HandleError(w, err, "Error fetching (film) sources", http.StatusInternalServerError)

	util.WriteJSONResponse(w, sources)
}

// ######## Anime ########

// AnimeSearch
// @Summary		Search for anime
// @Description	Search for anime by query
// @Tags			Anime
// @Accept			json
// @Produce		json
// @Param			q	query		string	true	"Search Query"
// @Success		200	{object}	anime.AnimeSearch
// @Router			/anime/all/search [get]
func AnimeSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results, err := anime.ProcessQuery(query)

	util.HandleError(w, err, "Error handling (anime) query", http.StatusInternalServerError)

	jsonResponse := map[string]interface{}{
		"results": results,
	}

	util.WriteJSONResponse(w, jsonResponse)
}

// FetchAnimeSources
// @Summary		Fetch anime sources
// @Description	Fetch anime sources by show ID
// @Tags			Anime
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"anime ID"
// @Param			tt	query	string	true	"translation Type"
// @Param			e	query	string	true	"episode Number"
// @Success		200	{array}	anime.AnimeSource
// @Router			/anime/all/sources [get]
func FetchAnimeSources(w http.ResponseWriter, r *http.Request) {
	animeId := r.URL.Query().Get("id")
	translationType := r.URL.Query().Get("tt")
	episodeString := r.URL.Query().Get("e")
	anime, err := anime.ProcessSources(animeId, translationType, episodeString)

	util.HandleError(w, err, "Error fetching (anime) sources", http.StatusInternalServerError)

	util.WriteJSONResponse(w, anime)
}

// ######## Shows ########

// ShowSearch
// @Summary		Search for shows
// @Description	Search for shows by query
// @Tags			Series
// @Accept			json
// @Produce		json
// @Param			q	query		string	true	"Search Query"
// @Success		200	{object}	show.ShowSearch
// @Router			/series/vip/search [get]
func ShowSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results, err := show.ProcessQuery(query)

	util.HandleError(w, err, "Error handling (show) query", http.StatusInternalServerError)

	jsonResponse := map[string]interface{}{
		"results": results,
	}

	util.WriteJSONResponse(w, jsonResponse)
}

// ShowSeason
// @Summary		Fetch show seasons and episodes
// @Description	Fetch show seasons and episodes by show ID
// @Tags			Series
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"Show ID"
// @Success		200	{array}	show.ShowSeason
// @Router			/series/vip/seasons [get]
func ShowSeason(w http.ResponseWriter, r *http.Request) {
	showId := r.URL.Query().Get("id")
	results, err := show.GetShowSeasons(showId)

	util.HandleError(w, err, "Error fetching (show) seasons & episodes", http.StatusInternalServerError)

	util.WriteJSONResponse(w, results)
}

// FetchShowServers
// @Summary		Fetch show servers
// @Description	Fetch show servers by episode ID
// @Tags			Series
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"Episode ID"
// @Success		200	{array}	show.ShowServer
// @Router			/series/vip/servers [get]
func FetchShowServers(w http.ResponseWriter, r *http.Request) {
	episodeId := r.URL.Query().Get("id")
	servers, err := show.GetShowServer(episodeId)

	util.HandleError(w, err, "Error fetching (show) servers", http.StatusInternalServerError)

	jsonResponse := map[string]interface{}{
		"servers": servers,
	}

	util.WriteJSONResponse(w, jsonResponse)
}

// FetchShowSources
// @Summary		Fetch show sources
// @Description	Fetch show servers by server ID
// @Tags			Series
// @Accept			json
// @Produce		json
// @Param			id	query	string	true	"Server ID"
// @Success		200	{array}	show.ShowSourcesEncrypted
// @Router			/series/vip/sources [get]
func FetchShowSources(w http.ResponseWriter, r *http.Request) {
	serverId := r.URL.Query().Get("id")
	sources, err := show.GetShowSources(serverId)

	util.HandleError(w, err, "Error fetching (show) sources", http.StatusInternalServerError)

	util.WriteJSONResponse(w, sources)
}

// ######## Novels ########

// NovelSearch
// @Summary		Search for novels
// @Description	Search for novels by query
// @Tags			Novels
// @Accept			json
// @Produce		json
// @Param			q	query		string	true	"Search Query"
// @Success		200	{object}	novel.NovelSearch
// @Router			/novels/rln/search [get]
func NovelSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results, err := novel.ProcessQuery(query)

	util.HandleError(w, err, "Error handling (novel) query", http.StatusInternalServerError)

	jsonResponse := map[string]interface{}{
		"results": results,
	}

	util.WriteJSONResponse(w, jsonResponse)
}

// ######## Users ########

// GetUsers
// @Summary		Get all users
// @Description	Retrieve a list of all users
// @ID				get-users
// @Tags			Users
// @Produce		json
// @Success		200	{array}	profiles.User
// @Router			/api/profiles/users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	color.Green("GET request received for all users")
	util.WriteJSONResponse(w, profiles.Users)
}

// CreateUser
// @Summary		Create a new user
// @Description	Create a new user with the given data
// @ID				create-user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			user	body		profiles.User	true	"User to be created"
// @Success		200		{object}	profiles.User
// @Router			/api/profiles/users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	color.Cyan("POST request received to create a new user")
	var user profiles.User
	err := json.NewDecoder(r.Body).Decode(&user)
	util.HandleError(w, err, err.Error(), http.StatusBadRequest)
	profiles.Users = append(profiles.Users, user)
	profiles.SaveUsers()
	util.WriteJSONResponse(w, user)
}

// GetUser
// @Summary		Get a specific user
// @Description	Retrieve a user by their username
// @ID				get-user
// @Tags			Users
// @Produce		json
// @Param			username	path		string	true	"Username of the user to be fetched"
// @Success		200			{object}	profiles.User
// @Failure		404			"User not found"
// @Router			/api/profiles/users/{username} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	color.Yellow("GET request received for a specific user")
	params := mux.Vars(r)
	username := params["username"]
	user, err := profiles.FindUserByUsername(username)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	util.WriteJSONResponse(w, user)
}

// UpdateUser
// @Summary		Update a user
// @Description	Update a user's data by their username
// @ID				update-user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			username	path		string			true	"Username of the user to be updated"
// @Param			updatedUser	body		profiles.User	true	"Updated user data"
// @Success		200			{object}	profiles.User
// @Failure		404			"User not found"
// @Router			/api/profiles/users/{username} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	color.Magenta("PUT request received to update a user")
	params := mux.Vars(r)
	username := params["username"]

	// Decode the updated user data from the request body
	var updatedUser profiles.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	util.HandleError(w, err, err.Error(), http.StatusBadRequest)

	// Find the existing user
	user, err := profiles.FindUserByUsername(username)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Merge the updated fields with the existing user data
	profiles.MergeUpdated(user, updatedUser)

	// Save the updated users list
	profiles.SaveUsers()

	util.WriteJSONResponse(w, user)
}

// DeleteUser
// @Summary		Delete a user
// @Description	Delete a user by their username
// @Tags			Users
// @ID				delete-user
// @Produce		json
// @Param			username	path		string	true	"Username of the user to be deleted"
// @Success		200			{object}	profiles.User
// @Failure		404			"User not found"
// @Router			/api/profiles/users/{username} [delete]
// DeleteUser deletes a user by their username and returns the deleted user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	color.Red("DELETE request received to delete a user")
	params := mux.Vars(r)
	username := params["username"]

	user, err := profiles.DeleteUser(username)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	util.WriteJSONResponse(w, user)
}

// ######## Home ########

// Home where the user goes when they need to poo-poo pee-pee
func Home(w http.ResponseWriter, r *http.Request) {
	color.Blue("GET request received for home page")
	response := map[string]string{"message": "Welcome to the Profiles API!"}
	util.WriteJSONResponse(w, response)
}
