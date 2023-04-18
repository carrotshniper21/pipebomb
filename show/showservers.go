package show

import (
	"fmt"
	"strings"

	"pipebomb/logging"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

// TODO Fetch Film servers via id
// TODO Add paramters for logging

func GetShowServer(episodeid, reqType, remoteAddress, reqPath, reqQueryParams string) ([]ShowServer, error) {
	c := colly.NewCollector()
	fmt.Println(color.GreenString(logging.HttpLogger()[0]+":"), color.HiWhiteString(" %s - '%s %s?%s'", remoteAddress, reqType, reqPath, reqQueryParams))

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
