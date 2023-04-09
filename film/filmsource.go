package film

import (
	"github.com/gocolly/colly"
)

func GetFilmSource(filmid string) {
	c := colly.NewCollector()

	getFilmData(c, filmid)
}

func getFilmData(c *colly.Collector, filmid string) string {
	c.OnHTML("a[data-linkid]", func(e *colly.HTMLElement) {
		linkID := e.Attr("data-linkid")
	})

	// Visit the URL to be scraped
	err = c.Visit("https://vipstream.tv/ajax/movie/episodes/" + filmid)
	if err != nil {
		return nil, err
	}
	return linkID
}
