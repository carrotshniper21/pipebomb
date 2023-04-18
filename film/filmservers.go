// pipebomb/film/filmsource.go
package film

import (
	"fmt"
	"strings"

	"pipebomb/logging"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

func GetFilmServer(filmid, reqType, remoteAddress, reqPath, reqQueryParams string) ([]FilmServer, error) {
	c := colly.NewCollector()
	fmt.Println(color.GreenString(logging.HttpLogger()[0]+":"), color.HiWhiteString(" %s - '%s %s?%s'", remoteAddress, reqType, reqPath, reqQueryParams))

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

	if err := c.Visit(root + "/ajax/movie/episodes/" + filmid); err != nil {
		return nil, err
	}

	return servers, nil
}
