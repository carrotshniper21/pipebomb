// pipebomb/show/showsearch.go
package show

import (
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// root is the root URL for the scraper to visit from (VIPStream)
const root = "https://vipstream.tv"

// showScraper scrapes the film page
func showSearcher(showUrl string) (*ShowSearch, error) {
	var show ShowSearch
	c := colly.NewCollector()

	setHrefCallback(c, &show)
	setPosterCallback(c, &show)
	setTitleCallback(c, &show)
	setIdCallback(showUrl, &show)
	setDescriptionCallback(c, &show)
	setReleasedCallback(c, &show)
	setCastCallback(c, &show)
	setGenreCallback(c, &show)
	setDurationCallback(c, &show)
	setCountryCallback(c, &show)
	setProductionCallback(c, &show)

	err := c.Visit(showUrl)
	if err != nil {
		return nil, err
	}

	return &show, nil
}

// setHrefCallback sets the show URL to the response struct
func setHrefCallback(c *colly.Collector, show *ShowSearch) {
	c.OnRequest(func(r *colly.Request) {
		show.Href = r.URL.String()
	})
}

// setPosterCallback sets the show poster to the response struct or a default image
func setPosterCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML(".dp-i-c-poster .film-poster-img", func(elem *colly.HTMLElement) {
		poster := elem.Attr("src")

		if poster != "" {
			show.Poster = poster
		} else {
			show.Poster = "https://i.imgur.com/3ZQZ9Zm.png"
		}
	})
}

func setTitleCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > h2 > a", func(elem *colly.HTMLElement) {
		show.Title = elem.Text
	})
}

// setIdCallback sets the show ID to the response struct
func setIdCallback(showUrl string, show *ShowSearch) {
	var idpart IdSplit
	idParts := strings.Split(showUrl, "/")
	if len(idParts) >= 5 {
		show.Id = idParts[3] + "/" + idParts[4]
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
	show.IdParts = idpart
}

func setDescriptionCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.description", func(elem *colly.HTMLElement) {
		description := elem.Text
		if description != "" {
			show.Description = description
		} else {
			show.Description = ""
		}
	})
}

func setReleasedCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-7.col-lg-7.col-md-8.col-sm-12 > div:nth-child(1)", func(elem *colly.HTMLElement) {
		released := elem.Text
		releasedParts := strings.Split(released, ":")
		show.Released = strings.ReplaceAll(releasedParts[1], " ", "")
	})
}

func setCastCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-7.col-lg-7.col-md-8.col-sm-12 > div:nth-child(3)", func(elem *colly.HTMLElement) {
		casts := strings.Replace(elem.Text, "Casts:", "", 1)
		castsParts := strings.Split(casts, ",")
		show.Cast = make([]string, len(castsParts))
		for i, cast := range castsParts {
			show.Cast[i] = strings.TrimSpace(cast)
		}
	})
}

func setGenreCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-7.col-lg-7.col-md-8.col-sm-12 > div:nth-child(2)", func(elem *colly.HTMLElement) {
		genres := strings.Replace(elem.Text, "Genre:", "", 1)
		genresParts := strings.Split(genres, ",")
		show.Genres = make([]string, len(genresParts))
		for i, genre := range genresParts {
			show.Genres[i] = strings.TrimSpace(genre)
		}
	})
}

func setDurationCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-5.col-lg-5.col-md-4.col-sm-12 > div:nth-child(1)", func(elem *colly.HTMLElement) {
		duration := strings.TrimSpace(strings.Replace(elem.Text, "Duration:", "", 1))
		duration = strings.Replace(duration, "min", "", 1)
		duration = strings.Replace(duration, "\n", "", 1)
		duration = strings.ReplaceAll(duration, " ", "")
		if strings.Contains(duration, "N/A") {
			show.Duration = "N/A"
		} else {
			show.Duration = duration + " min"
		}
	})
}

func setCountryCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-5.col-lg-5.col-md-4.col-sm-12 > div:nth-child(2)", func(elem *colly.HTMLElement) {
		country := strings.Replace(elem.Text, "Country:", "", 1)
		if strings.Contains(country, ",") {
			countryParts := strings.Split(country, ",")
			show.Country = make([]string, len(countryParts))
			for i, country := range countryParts {
				show.Country[i] = strings.TrimSpace(country)
			}
		} else {
			country = strings.TrimSpace(country)
			show.Country = []string{country}
		}
	})
}

func setProductionCallback(c *colly.Collector, show *ShowSearch) {
	c.OnHTML("#main-wrapper > div.detail_page.detail_page-style > div > div.detail_page-watch > div.detail_page-infor.border-bottom-block > div > div.dp-i-c-right > div.elements > div > div.col-xl-5.col-lg-5.col-md-4.col-sm-12 > div:nth-child(3)", func(elem *colly.HTMLElement) {
		production := strings.Replace(elem.Text, "Production:", "", 1)
		if strings.Contains(production, ",") {
			productionParts := strings.Split(production, ",")
			show.Production = make([]string, len(productionParts))
			for i, production := range productionParts {
				show.Production[i] = strings.TrimSpace(production)
			}
		} else {
			production = strings.TrimSpace(production)
			show.Production = []string{production}
		}

	})
}

// ProcessLink processes the link and returns a ShowStruct
func processLink(elem *colly.HTMLElement, visitedLinks *sync.Map) *ShowSearch {
	showid := elem.Attr("href")
	if strings.Contains(showid, "/tv/watch") {
		absLink := root + showid
		if _, visited := visitedLinks.Load(absLink); visited {
			return nil
		}
		visitedLinks.Store(absLink, struct{}{})

		show, err := showSearcher(absLink)
		if err != nil {
			return nil
		}

		return show
	}

	return nil
}

func ProcessQuery(query string) (interface{}, error) {
    visitedLinks := sync.Map{}
    c := colly.NewCollector()

    var results []*ShowSearch

    c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
        show := processLink(elem, &visitedLinks)
        if show != nil {
            results = append(results, show)
        }
    })

    err := c.Visit("https://flixhq.to/search/" + query)
    if err != nil {
        return nil, err
    }

    return results, nil
}
