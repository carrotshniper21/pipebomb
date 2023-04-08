// scrape.go
package film

import (
	"encoding/json"
  "pipebomb/logging"
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

// scrape.go
func ProcessLink(elem *colly.HTMLElement, visitedLinks *sync.Map) *FilmStruct {
    link := elem.Attr("href")
    if strings.Contains(link, "/movie/watch") {
        absLink := root + link
        if _, visited := visitedLinks.Load(absLink); visited {
            return nil
        }
        visitedLinks.Store(absLink, struct{}{})

        film, err := filmScraper(absLink)
        if err != nil {
            return nil
        }

        color.Cyan(logging.Random007Phrase())

        return film
    }

    return nil
}

func outputJSON(response FilmResponse) string {
    jsonBytes, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
    }
    return string(jsonBytes)
}
