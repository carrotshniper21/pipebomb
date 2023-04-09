// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"pipebomb/film"
	"sync"
)

// FilmSearch searches for a film and returns a JSON response
// @Summary Search for a film
// @Description Searches for a film and returns a JSON response
// @Tags films
// @Accept  json
// @Produce  json
// @Param   q query string true "Search Query"
// @Success 200 {object} map[string]interface{}
// @Router /films/vip/search [get]
func FilmSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	var results []*film.FilmStruct

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		film := film.ProcessLink(elem, &visitedLinks)
		if film != nil {
			results = append(results, film)
		}
	})

	c.OnScraped(func(response *colly.Response) {
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
		w.Write(responseBytes)
	})

	c.Visit("https://flixhq.to/search/" + query)
}

// Home is the home page
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
}
