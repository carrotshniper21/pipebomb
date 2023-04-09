package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"pipebomb/film"

	"github.com/gocolly/colly"
)

// @Summary Fetch film sources
// @Description Fetch film sources by film ID
// @Tags film
// @Accept  json
// @Produce  json
// @Param   id query string true "Film ID"
// @Success 200 {array} film.FilmSource
// @Router /fetch-films [get]
func FetchFilms(w http.ResponseWriter, r *http.Request) {
	filmID := r.URL.Query().Get("id")

	sources, err := film.GetFilmSource(filmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(sources)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, s := w.Write(responseBytes)
	if s != nil {
		fmt.Println("error writing response for film sources: ", s)
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
// @Tags film
// @Accept  json
// @Produce  json
// @Param   q query string true "Search Query"
// @Success 200 {object} map[string]interface{}
// @Router /search-films [get]
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

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
	if err != nil {
		fmt.Println("error writing home page: ", err)
		return
	}
}
