// scrape.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

const root = "https://vipstream.tv"

// setRequestCallback sets the film URL to the response struct
func setRequestCallback(c *colly.Collector, film *FilmStruct) {
	c.OnRequest(func(r *colly.Request) {
		film.Href = r.URL.String()
	})
}

// FilmResponse is the response struct for the film scraper
func filmScraper(filmUrl string) (*FilmStruct, error) {
	var film FilmStruct
	c := colly.NewCollector()

	setPosterCallback(c, &film)
	setRequestCallback(c, &film)

	err := c.Visit(filmUrl)
	if err != nil {
		return nil, err
	}

	return &film, nil
}

// setPosterCallback sets the film poster to the response struct or a default image
func setPosterCallback(c *colly.Collector, film *FilmStruct) {
	c.OnHTML(".dp-i-c-poster .film-poster-img", func(elem *colly.HTMLElement) {
		poster := elem.Attr("src")

		if poster != "" {
			film.Poster = poster
		} else {
			film.Poster = "https://i.imgur.com/3ZQZ9Zm.png"
		}
	})
}

func processLink(elem *colly.HTMLElement, visitedLinks *sync.Map) {
	link := elem.Attr("href")
	if strings.Contains(link, "/movie/watch") {
		absLink := root + link
		if _, visited := visitedLinks.Load(absLink); visited {
			return
		}
		visitedLinks.Store(absLink, struct{}{})

		film, err := filmScraper(absLink)
		response := FilmResponse{}
		if err != nil {
			response.Status = "error"
		} else {
			response.Status = "ok"
			response.Film = film
		}

		outputJSON(response)
		color.Cyan(random007Phrase())
	}
}

func outputJSON(response FilmResponse) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}
