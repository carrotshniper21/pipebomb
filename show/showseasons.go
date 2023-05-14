// pipebomb/show/showseasons.go
package show

import (
	"strings"
	"log"

	"github.com/gocolly/colly"
)

func GetShowSeasons(showid string) (map[string]ShowSeason, error) {
	c := colly.NewCollector() 

	response, _ := getSeasonData(c, showid)

	seasonsMap := make(map[string]ShowSeason)
	for _, season := range response {
		seasonsMap[season.SeasonName] = season
	}

	return seasonsMap, nil
}

func getSeasonData(c *colly.Collector, showid string) ([]ShowSeason, error) {
	seasons := []ShowSeason{}

	c.OnHTML("div.dropdown-menu > a[data-id]", func(e *colly.HTMLElement) {
		seasonID := e.Attr("data-id")
		episodes, err := getEpisodeData(c, seasonID)
		if err != nil {
			log.Println(err)
		}
		seasonName := strings.TrimSpace(e.Text)
		seasonName = strings.ReplaceAll(seasonName, "\n", "")
		season := ShowSeason{
			SeasonName: seasonName, 
			SeasonID: seasonID,
			Episodes: episodes,
		}
		seasons = append(seasons, season)
	})

	if err := c.Visit(root + "/ajax/v2/tv/seasons/" + showid); err != nil {
		return nil, err
	}

	return seasons, nil
}

func getEpisodeData(c *colly.Collector, seasonID string) ([]Episode, error) {
	episodes := []Episode{}

	c.OnHTML("a.eps-item", func(elem *colly.HTMLElement) {
		title := strings.TrimSpace(elem.Attr("title"))
		episodeID := elem.Attr("data-id")
		episode := Episode{
			Title: title,
			EpisodeID: episodeID,
		}
		episodes = append(episodes, episode)
	})

	if err := c.Visit(root + "/ajax/v2/season/episodes/" + seasonID); err != nil {
		return nil, err
	}

	return episodes, nil
}
