// film/scrape.go
package film

import (
	"encoding/json"
	"pipebomb/logging"
	"strconv"
	"strings"
	"sync"
	"log"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

// root is the root URL for the scraper to visit from (VIPStream)
const root = "https://vipstream.tv"

// setRequestCallback sets the film URL to the response struct
func setRequestCallback(c *colly.Collector, film *FilmStruct) {
	c.OnRequest(func(r *colly.Request) {
		film.Href = r.URL.String()
	})
}

// filmScraper scrapes the film page
func filmScraper(filmUrl string) (*FilmStruct, error) {
	var film FilmStruct
	c := colly.NewCollector()

	setPosterCallback(c, &film)
	setRequestCallback(c, &film)
	setIdCallback(filmUrl, &film)
	setDescriptionCallback(c, &film)

	err := c.Visit(filmUrl)
	if err != nil {
		return nil, err
	}

	return &film, nil
}

func setDescriptionCallback(c *colly.Collector, film *FilmStruct) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.description", func(elem *colly.HTMLElement) {
		description := elem.Text
		if description != "" {
			film.Description = description
		} else {
			film.Description = ""
		}
	})
}

// setIdCallback sets the film ID to the response struct
func setIdCallback(filmUrl string, film *FilmStruct) {
	var idpart IdSplit
	idParts := strings.Split(filmUrl, "/")
	if len(idParts) >= 5 {
		film.Id = idParts[4]
		idpart.Type = idParts[3]
		nameAndId := strings.Split(idParts[4], "-")
		if len(nameAndId) > 1 {
			idNum, err := strconv.Atoi(nameAndId[len(nameAndId)-1])
			if err == nil {
				idpart.IdNum = idNum
				nameAndId = nameAndId[:len(nameAndId)-1]
			}
			idpart.Name = strings.Join(nameAndId, "-")
			idpart.Name = strings.TrimPrefix(idpart.Name, "watch-")
		}
	}
	film.IdParts = idpart
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

// ProcessLink processes the link and returns a FilmStruct
func ProcessLink(elem *colly.HTMLElement, visitedLinks *sync.Map) *FilmStruct {
	filmid := elem.Attr("href")
	if strings.Contains(filmid, "/movie/watch") {
		absLink := root + filmid
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

// outputJSON returns a JSON string from a FilmResponse struct
func outputJSON(response FilmResponse) string {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}
