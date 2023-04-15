// pipebomb/show/showservers.go

package show


/*
import (
	"strings"

	"github.com/gocolly/colly"
)

func GetShowServer(episodeid string) ([]ShowServer, error) {
	c := colly.NewCollector()
}

func getServerDataid(c *colly.Collector, episodeid string) ([]ShowServer, error) {
	servers := []ShowServer{}

	c.OnHTML("a[data-linkid]", func(e *colly.HTMLElement) {
		linkID := e.Attr("data-linkid")
		serverName := strings.TrimSpace(e.Text)
		serverName = strings.ReplaceAll(serverName, "\n", "")
		servers = append(servers, ShowServer{ServerName: serverName, LinkID: linkID})
	})

	err := c.Visit("https://vipstream.tv/ajax/
}
*/
