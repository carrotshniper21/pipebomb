// pipebomb/handlers/handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(sources)
	util.LoggingError("(film) sources", err)
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

	if err != nil {
		http.Error(w, "Error fetching (show) servers", http.StatusInternalServerError)
		return
	}

	jsonResponse := map[string]interface{}{
		"servers": servers,
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, s := w.Write(responseBytes)
	if s != nil {
		fmt.Println("Error writing response for (show) servers: ", s)
		return
	}
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

	if err != nil {
		http.Error(w, "Error fetching (show) sources", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(sources)
	if err != nil {
		fmt.Println("Error writing response for (show) sources: ", err)
	}
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles.Users)
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
	w.Header().Set("Content-Type", "application/json")
	var user profiles.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profiles.Users = append(profiles.Users, user)
	profiles.SaveUsers()
	json.NewEncoder(w).Encode(user)
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	for _, user := range profiles.Users {
		if user.Username == username {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]
	for i, user := range profiles.Users {
		if user.Username == username {
			var updatedUser profiles.User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Merge the updated fields with the existing user data
			if updatedUser.Username != "" {
				user.Username = updatedUser.Username
			}
			if updatedUser.Avatar != "" {
				user.Avatar = updatedUser.Avatar
			}
			profile := &user.Profile
			updatedProfile := &updatedUser.Profile
			if updatedProfile.Name != "" {
				profile.Name = updatedProfile.Name
			}
			if updatedProfile.Image != "" {
				profile.Image = updatedProfile.Image
			}
			if updatedProfile.Bio != "" {
				profile.Bio = updatedProfile.Bio
			}
			if updatedProfile.Philosophy != "" {
				profile.Philosophy = updatedProfile.Philosophy
			}

			profiles.Users[i] = user
			profiles.SaveUsers()
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
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
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	color.Red("DELETE request received to delete a user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]

	if username != "Space Mommy" {
		for i, user := range profiles.Users {
			if user.Username == username {
				profiles.Users = append(profiles.Users[:i], profiles.Users[i+1:]...)
				profiles.SaveUsers()
				json.NewEncoder(w).Encode(user)
				return
			}
		}
	}

	http.NotFound(w, r)
}

// ######## Home ########

// Home just a place for the user to go when they need to go poo-poo pee-pee
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
	if err != nil {
		fmt.Println("error writing home page: ", err)
		return
	}
}
