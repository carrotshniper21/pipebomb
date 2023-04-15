package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"pipebomb/logging"
	"pipebomb/film"
	"pipebomb/show"
)

// @Summary Fetch film sources
// @Description Fetch film servers by server ID
// @Tags films
// @Accept json
// @Produce json
// @Param id query string true "Server ID"
// @Success 200 {array} film.FilmSources
// @Router /films/vip/sources [get]
func FetchFilmSources(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("id")
	reqType := r.Method
	remoteAddress := r.RemoteAddr
	reqUrl := r.URL

	sources := film.GetFilmSources(serverID, reqType, remoteAddress, reqUrl.Path, reqUrl.RawQuery)
	if sources == nil {
		http.Error(w, "Error fetching film sources", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(sources)
	if err != nil {
		fmt.Println("error writing response for film sources: ", err)
	}
}

// @Summary Fetch film servers
// @Description Fetch film servers by film ID
// @Tags films
// @Accept  json
// @Produce  json
// @Param   id query string true "Film ID"
// @Success 200 {array} film.FilmServer
// @Router /films/vip/servers [get]
func FetchFilms(w http.ResponseWriter, r *http.Request) {
	filmID := r.URL.Query().Get("id")
	reqType := r.Method
	remoteAddress := r.RemoteAddr
	reqUrl := r.URL

	servers, err := film.GetFilmServer(filmID, reqType, remoteAddress, reqUrl.Path, reqUrl.RawQuery)

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

func searchFilms(query, reqType, remoteAddress, reqPath, reqQueryParams string) (interface{}, error) {
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	var results []*film.FilmSearch

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		film := film.ProcessLink(elem, &visitedLinks)
		if film != nil {
			results = append(results, film)
		}
	})

	fmt.Println(color.GreenString(logging.HttpLogger()[0] + ":"), color.HiWhiteString(" '%s - %s %s?%s'", remoteAddress, reqType, reqPath, reqQueryParams))
	err := c.Visit("https://flixhq.to/search/" + query)
	if err != nil {
		return nil, err
	}

	return results, nil
}


// @Summary Search for films
// @Description Search for films by query
// @Tags films
// @Accept  json
// @Produce  json
// @Param   q query string true "Search Query"
// @Success 200 {object} film.FilmSearch
// @Router /films/vip/search [get]
func FilmSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	reqType := r.Method
	remoteAddress := r.RemoteAddr
	reqUrl := r.URL

	results, _ := searchFilms(query, reqType, remoteAddress, reqUrl.Path, reqUrl.RawQuery)

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


func searchShows(query, reqType, remoteAddress, reqPath, reqQueryParams string) (interface{}, error) {
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	var results []*show.ShowSearch

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		show := show.ProcessLink(elem, &visitedLinks)
		if show != nil {
			results = append(results, show)
		}
	})

	fmt.Println(color.GreenString(logging.HttpLogger()[0] + ":"), color.HiWhiteString(" '%s - %s %s?%s'", remoteAddress, reqType, reqPath, reqQueryParams))
	err := c.Visit("https://flixhq.to/search/" + query)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// @Summary Search for shows
// @Description Search for shows by query
// @Tags shows
// @Accept  json
// @Produce  json
// @Param   q query string true "Search Query"
// @Success 200 {object} show.ShowSearch
// @Router /shows/vip/search [get]
func ShowSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	reqType := r.Method
	remoteAddress := r.RemoteAddr
	reqUrl := r.URL

	results, _ := searchShows(query, reqType, remoteAddress, reqUrl.Path, reqUrl.RawQuery)

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

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
	if err != nil {
		fmt.Println("error writing home page: ", err)
		return
	}
}
