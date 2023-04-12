// pipebomb/show/showseasons.go
package main

import (
    "encoding/json"
    "fmt"
    "strings"

    "github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector()

    showid := "39480"

    seasons, err := getSeasonDataid(c, showid)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    seasonsMap := make(map[string]ShowSeason)
    for _, season := range seasons {
        seasonsMap[season.SeasonName] = season
    }

    jsonData, err := json.Marshal(seasonsMap)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonData))
}

func getSeasonDataid(c *colly.Collector, showid string) ([]ShowSeason, error) {
    seasons := []ShowSeason{}

    c.OnHTML("a[data-id]", func(e *colly.HTMLElement) {
        seasonID := e.Attr("data-id")
        episodes, err := getEpisodeDataid(c, seasonID)
        if err != nil {
            fmt.Println(err)
        }
        seasonName := strings.TrimSpace(e.Text)
        seasonName = strings.ReplaceAll(seasonName, "\n", "")
        season := ShowSeason{
            SeasonName: seasonName,
            SeasonID:   seasonID,
            Episodes:   episodes,
        }
				fmt.Println(season)
        seasons = append(seasons, season)
    })

    err := c.Visit("https://vipstream.tv/ajax/v2/tv/seasons/" + showid)
    if err != nil {
        return nil, err
    }

    return seasons, nil
}

func getEpisodeDataid(c *colly.Collector, seasonID string) ([]Episode, error) {
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

    err := c.Visit("https://vipstream.tv/ajax/v2/season/episodes/" + seasonID)
    if err != nil {
        return nil, err
    }

    return episodes, nil
}
