// pipebomb/film/filmservers.go
package film

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetFilmServer(filmId string) ([]FilmServer, error) {
	c := colly.NewCollector()

	return getServerDataid(c, filmId)
}
func getServerDataid(c *colly.Collector, filmId string) ([]FilmServer, error) {
	servers := []FilmServer{}

	c.OnHTML("a[data-linkid]", func(e *colly.HTMLElement) {
		linkID := e.Attr("data-linkid")
		serverName := strings.TrimSpace(e.Text)
		serverName = strings.ReplaceAll(serverName, "\n", "")
		servers = append(servers, FilmServer{ServerName: serverName, LinkID: linkID})
	})

	if err := c.Visit(root + "/ajax/movie/episodes/" + filmId); err != nil {
		return nil, err
	}

	return servers, nil
}
