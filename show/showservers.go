// pipebomb/show/showservers.go
package show

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetShowServer(episodeid string) ([]ShowServer, error) {
	c := colly.NewCollector()

	return getServerDataid(c, episodeid)
}

func getServerDataid(c *colly.Collector, episodeid string) ([]ShowServer, error) {
	servers := []ShowServer{}

	c.OnHTML("a[data-id]", func(e *colly.HTMLElement) {
		linkID := e.Attr("data-id")
		serverName := strings.TrimSpace(e.Text)
		serverName = strings.ReplaceAll(serverName, "\n", "")
		servers = append(servers, ShowServer{ServerName: serverName, LinkID: linkID})
	})

	if err := c.Visit(root + "/ajax/v2/episode/servers/" + episodeid + "/#servers-list"); err != nil {
		return nil, err
	}

	return servers, nil
}
