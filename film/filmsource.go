// pipebomb/film/filmsource.go
package film

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetFilmSource(filmid string) ([]FilmSource, error) {
	c := colly.NewCollector()

	return getFilmData(c, filmid)
}
func getFilmData(c *colly.Collector, filmid string) ([]FilmSource, error) {
	sources := []FilmSource{}

	c.OnHTML("a[data-linkid]", func(e *colly.HTMLElement) {
		linkID := e.Attr("data-linkid")
		serverName := strings.TrimSpace(e.Text)
		serverName = strings.ReplaceAll(serverName, "\n", "")
		sources = append(sources, FilmSource{ServerName: serverName, LinkID: linkID})
	})

	err := c.Visit("https://vipstream.tv/ajax/movie/episodes/" + filmid)
	if err != nil {
		return nil, err
	}

	return sources, nil
}
