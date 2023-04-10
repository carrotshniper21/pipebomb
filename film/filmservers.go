// pipebomb/film/filmsource.go
package film

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetFilmServer(filmid string) ([]FilmServer, error) {
	c := colly.NewCollector()

	return getServerDataid(c, filmid)
}
func getServerDataid(c *colly.Collector, filmid string) ([]FilmServer, error) {
	servers := []FilmServer{}

	c.OnHTML("a[data-linkid]", func(e *colly.HTMLElement) {
		linkID := e.Attr("data-linkid")
		serverName := strings.TrimSpace(e.Text)
		serverName = strings.ReplaceAll(serverName, "\n", "")
		servers = append(servers, FilmServer{ServerName: serverName, LinkID: linkID})
	})

	err := c.Visit("https://vipstream.tv/ajax/movie/episodes/" + filmid)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
