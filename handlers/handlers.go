package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"pipebomb/cache"
	"pipebomb/film"
	"sync"
)

var redisCache *cache.RedisCache

func init() {
	redisCache = cache.NewCache("localhost:6379", "", 0)
}

func fetchFilms(query string) (interface{}, error) {
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	var results []*film.FilmStruct

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

func FilmSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	cacheKey := fmt.Sprintf("films_search_%s", query)

	data, err := cache.CacheData(redisCache, cacheKey, fetchFilms, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results := data.([]*film.FilmStruct)

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
