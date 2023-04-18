// pipebomb/film/filmsearch.go
package film

import (
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// root is the root URL for the scraper to visit from (VIPStream)
const root = "https://vipstream.tv"

// filmScraper scrapes the film page
func filmSearcher(filmUrl string) (*FilmSearch, error) {
	var film FilmSearch
	c := colly.NewCollector()

	setHrefCallback(c, &film)
	setPosterCallback(c, &film)
	setTitleCallback(c, &film)
	setIdCallback(filmUrl, &film)
	setDescriptionCallback(c, &film)
	setReleasedCallback(c, &film)
	setCastCallback(c, &film)
	setGenreCallback(c, &film)
	setDurationCallback(c, &film)
	setCountryCallback(c, &film)
	setProductionCallback(c, &film)

	err := c.Visit(filmUrl)
	if err != nil {
		return nil, err
	}

	return &film, nil
}

// setHrefCallback sets the film URL to the response struct
func setHrefCallback(c *colly.Collector, film *FilmSearch) {
	c.OnRequest(func(r *colly.Request) {
		film.Href = r.URL.String()
	})
}

// setPosterCallback sets the film poster to the response struct or a default image
func setPosterCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML(".dp-i-c-poster .film-poster-img", func(elem *colly.HTMLElement) {
		poster := elem.Attr("src")

		if poster != "" {
			film.Poster = poster
		} else {
			film.Poster = "https://i.imgur.com/3ZQZ9Zm.png"
		}
	})
}

func setTitleCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > h2 > a", func(elem *colly.HTMLElement) {
		film.Title = elem.Text
	})
}

// setIdCallback sets the film ID to the response struct
func setIdCallback(filmUrl string, film *FilmSearch) {
	var idpart IdSplit
	idParts := strings.Split(filmUrl, "/")
	if len(idParts) >= 5 {
		film.Id = idParts[3] + "/" + idParts[4]
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

func setDescriptionCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.description", func(elem *colly.HTMLElement) {
		description := elem.Text
		if description != "" {
			film.Description = description
		} else {
			film.Description = ""
		}
	})
}

func setReleasedCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-7.col-lg-7.col-md-8.col-sm-12 > div:nth-child(1)", func(elem *colly.HTMLElement) {
		released := elem.Text
		releasedParts := strings.Split(released, ":")
		film.Released = strings.ReplaceAll(releasedParts[1], " ", "")
	})
}

func setCastCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-7.col-lg-7.col-md-8.col-sm-12 > div:nth-child(3)", func(elem *colly.HTMLElement) {
		casts := strings.Replace(elem.Text, "Casts:", "", 1)
		castsParts := strings.Split(casts, ",")
		film.Cast = make([]string, len(castsParts))
		for i, cast := range castsParts {
			film.Cast[i] = strings.TrimSpace(cast)
		}
	})
}

func setGenreCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-7.col-lg-7.col-md-8.col-sm-12 > div:nth-child(2)", func(elem *colly.HTMLElement) {
		genres := strings.Replace(elem.Text, "Genre:", "", 1)
		genresParts := strings.Split(genres, ",")
		film.Genres = make([]string, len(genresParts))
		for i, genre := range genresParts {
			film.Genres[i] = strings.TrimSpace(genre)
		}
	})
}

func setDurationCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-5.col-lg-5.col-md-4.col-sm-12 > div:nth-child(1)", func(elem *colly.HTMLElement) {
		duration := strings.TrimSpace(strings.Replace(elem.Text, "Duration:", "", 1))
		duration = strings.Replace(duration, "min", "", 1)
		duration = strings.Replace(duration, "\n", "", 1)
		duration = strings.ReplaceAll(duration, " ", "")
		if strings.Contains(duration, "N/A") {
			film.Duration = "N/A"
		} else {
			film.Duration = duration + " min"
		}
	})
}

func setCountryCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-5.col-lg-5.col-md-4.col-sm-12 > div:nth-child(2)", func(elem *colly.HTMLElement) {
		country := strings.Replace(elem.Text, "Country:", "", 1)
		if strings.Contains(country, ",") {
			countryParts := strings.Split(country, ",")
			film.Country = make([]string, len(countryParts))
			for i, country := range countryParts {
				film.Country[i] = strings.TrimSpace(country)
			}
		} else {
			country = strings.TrimSpace(country)
			film.Country = []string{country}
		}
	})
}

func setProductionCallback(c *colly.Collector, film *FilmSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-5.col-lg-5.col-md-4.col-sm-12 > div:nth-child(3)", func(elem *colly.HTMLElement) {
		production := strings.Replace(elem.Text, "Production:", "", 1)
		if strings.Contains(production, ",") {
			productionParts := strings.Split(production, ",")
			film.Production = make([]string, len(productionParts))
			for i, production := range productionParts {
				film.Production[i] = strings.TrimSpace(production)
			}
		} else {
			production = strings.TrimSpace(production)
			film.Production = []string{production}
		}

	})
}

// ProcessLink processes the link and returns a FilmStruct
func ProcessLink(elem *colly.HTMLElement, visitedLinks *sync.Map) *FilmSearch {
	filmid := elem.Attr("href")
	if strings.Contains(filmid, "/movie/watch") {
		absLink := root + filmid
		if _, visited := visitedLinks.Load(absLink); visited {
			return nil
		}
		visitedLinks.Store(absLink, struct{}{})

		film, err := filmSearcher(absLink)
		if err != nil {
			return nil
		}

		return film
	}

	return nil
}
