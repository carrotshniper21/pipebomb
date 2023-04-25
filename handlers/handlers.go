// pipebomb/handlers/handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"pipebomb/profiles"
	"pipebomb/film"
	"pipebomb/show"
)

// @Summary Fetch show sources
// @Description Fetch show servers by server ID
// @Tags Series
// @Accept json
// @Produce json
// @Param id query string true "Server ID"
// @Success 200 {array} show.ShowSourcesEncrypted
// @Router /series/vip/sources [get]
func FetchShowSources(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("id")

	sources, err := show.GetShowSources(serverID)
	if err != nil {
		http.Error(w, "Error fetching show sources", http.StatusInternalServerError)
		return }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(sources)
	if err != nil {
		fmt.Println("error writing response for show sources: ", err)
	}
}

// @Summary Fetch film sources
// @Description Fetch film servers by server ID
// @Tags Films
// @Accept json
// @Produce json
// @Param id query string true "Server ID"
// @Success 200 {array} film.FilmSourcesEncrypted
// @Router /films/vip/sources [get]
func FetchFilmSources(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("id")

	sources, err := film.GetFilmSources(serverID)
	if err != nil {
		http.Error(w, "Error fetching film sources", http.StatusInternalServerError)
		return }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(sources)
	if err != nil {
		fmt.Println("error writing response for film sources: ", err)
	}
}

// @Summary Fetch film servers
// @Description Fetch film servers by film ID
// @Tags Films
// @Accept  json
// @Produce  json
// @Param   id query string true "Film ID"
// @Success 200 {array} film.FilmServer
// @Router /films/vip/servers [get]
func FetchFilms(w http.ResponseWriter, r *http.Request) {
	filmID := r.URL.Query().Get("id")
	servers, err := film.GetFilmServer(filmID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(servers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, s := w.Write(responseBytes)
	if s != nil {
		fmt.Println("error writing response for film servers: ", s)
		return
	}
}

func searchFilms(query string) (interface{}, error) {
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	var results []*film.FilmSearch

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		film := film.ProcessLink(elem, &visitedLinks)
		if film != nil {
			results = append(results, film)
		}
	})

	err := c.Visit("https://flixhq.to/search/" + query)
	if err != nil {
		return nil, err
	}

	return results, nil
}


// @Summary Search for films
// @Description Search for films by query
// @Tags Films
// @Accept  json
// @Produce  json
// @Param   q query string true "Search Query"
// @Success 200 {object} film.FilmSearch
// @Router /films/vip/search [get]
func FilmSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	results, _ := searchFilms(query)

	jsonResponse := map[string]interface{}{
		"results": results,
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
		fmt.Println("error writing response for film search: ", s)
		return
	}
}


func searchShows(query string) (interface{}, error) {
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	var results []*show.ShowSearch

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		show := show.ProcessLink(elem, &visitedLinks)
		if show != nil {
			results = append(results, show)
		}
	})

	err := c.Visit("https://flixhq.to/search/" + query)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// @Summary Search for shows
// @Description Search for shows by query
// @Tags Series
// @Accept  json
// @Produce  json
// @Param   q query string true "Search Query"
// @Success 200 {object} show.ShowSearch
// @Router /series/vip/search [get]
func ShowSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	
	results, _ := searchShows(query)

	jsonResponse := map[string]interface{}{
		"results": results,
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
		fmt.Println("Error writing response for film search: ", s)
		return
	}
}

func showSeasons(query string) (map[string]show.ShowSeason, error) {
	response, err := show.GetShowSeason(query)
	if err != nil {
		return nil, err
	}
	seasonsMap := make(map[string]show.ShowSeason)
	for _, season := range response {
		seasonsMap[season.SeasonName] = season
	}

	return seasonsMap, nil
}


// @Summary Fetch show seasons and episodes
// @Description Fetch show seasons and episodes by id
// @Tags Series
// @Accept json
// @Produce json
// @Param id query string true "Search Query"
// @Success 200 {array} show.ShowSeason
// @Router /series/vip/seasons [get]
func ShowSeason(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")

	results, _ := showSeasons(query)

	responseBytes, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, s := w.Write(responseBytes)
	if s != nil {
		fmt.Println("Error writing repsonse for show servers: ", s)
		return
	}
}

// @Summary Fetch show servers
// @Description Fetch show servers by episode ID
// @Tags Series
// @Accept  json
// @Produce  json
// @Param   id query string true "Episode ID"
// @Success 200 {array} show.ShowServer
// @Router /series/vip/servers [get]
func FetchShows(w http.ResponseWriter, r *http.Request) {
	filmID := r.URL.Query().Get("id")

	servers, err := show.GetShowServer(filmID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(servers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, s := w.Write(responseBytes)
	if s != nil {
		fmt.Println("error writing response for film servers: ", s)
		return
	}
}


// @Summary     Get all users
// @Description Retrieve a list of all users
// @ID              get-users
// @Tags Users
// @Produce     json
// @Success     200 {array} profiles.User
// @Router          /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
    color.Green("GET request received for all users")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(profiles.Users)
}

// @Summary     Create a new user
// @Description Create a new user with the given data
// @ID              create-user
// @Tags Users
// @Accept          json
// @Produce     json
// @Param           user    body        profiles.User    true    "User to be created"
// @Success     200     {object}    profiles.User
// @Router          /users [post]
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

// @Summary     Get a specific user
// @Description Retrieve a user by their username
// @ID              get-user
// @Tags Users
// @Produce     json
// @Param           username    path        string  true    "Username of the user to be fetched"
// @Success     200         {object}    profiles.User
// @Failure     404         "User not found"
// @Router          /users/{username} [get]
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

// @Summary     Update a user
// @Description Update a user's data by their username
// @ID              update-user
// @Tags Users
// @Accept          json
// @Produce     json
// @Param           username    path        string  true    "Username of the user to be updated"
// @Param           updatedUser body        profiles.User    true    "Updated user data"
// @Success     200         {object}    profiles.User
// @Failure     404         "User not found"
// @Router          /users/{username} [put]
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

// @Summary     Delete a user
// @Description Delete a user by their username
// @Tags Users
// @ID              delete-user
// @Produce     json
// @Param           username    path        string  true    "Username of the user to be deleted"
// @Success     200         {object}    profiles.User
// @Failure     404         "User not found"
// @Router          /users/{username} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    color.Red("DELETE request received to delete a user")
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    username := params["username"]

    if username != "Space Mommy" || username != "Space%20Mommy" {

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


func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
	if err != nil {
		fmt.Println("error writing home page: ", err)
		return
	}
}
