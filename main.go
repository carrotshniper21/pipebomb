// main.go
package main

import (
  "log"
  "fmt"
	"sync"
  "net/http"

  "github.com/gorilla/mux"
	"github.com/gocolly/colly"
)

func CollyInit(query string) string {
	visitedLinks := sync.Map{}
  
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
    outputJson = processLink(elem, &visitedLinks)
	})

	c.Visit("https://flixhq.to/search/" + query)

  return outputJson
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  query := r.URL.Query("q")
  fmt.Fprintf(w, CollyInit(query))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "<html><body><h1>Welcome to the home page. Nyaa~~</h1></body></html>")
}

func main() {
  port := "8001"
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  r.HandleFunc("/api/films/vip/search", SearchHandler)
  fmt.Println("http://localhost:8001/api/films/vip/search?q=puss")
  log.Fatal(http.ListenAndServe(":"+port, r))
}
