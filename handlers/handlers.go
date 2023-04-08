package handlers

import (
  "fmt"
  "sync"
  "net/http"
  "encoding/json"
  "pipebomb/film"
  "github.com/gocolly/colly"
)

func FilmSearch(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    visitedLinks := sync.Map{}
    c := colly.NewCollector()

    results := []*film.FilmStruct{}

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

func Home(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
}
