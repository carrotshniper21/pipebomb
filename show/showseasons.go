// pipebomb/show/showseasons.go
package show

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetShowSeason(showid string) ([]ShowSeason, error) {
	c := colly.NewCollector()

	return getSeasonDataid(c, showid)
}


func getSeasonDataid(c *colly.Collector, showid string) ([]ShowSeason, error) {
	servers := []ShowSeason{}

	c.OnHTML("a[data-id]", func(e *colly.HTMLElement) {
		SeasonID := e.Attr("data-id")
		seasonName := strings.TrimSpace(e.Text)
		seasonName = strings.ReplaceAll(seasonName, "\n", "")
		servers = append(servers, ShowSeason{SeasonName: seasonName, SeasonID: SeasonID})
	})

	err := c.Visit("https://vipstream.tv/ajax/v2/tv/seasons/" + showid)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
