package novel

import (
    "fmt"
    "strings"

    "github.com/gocolly/colly"
)

const root = "https://readlightnovels.net"

func novelSearcher(query string) ([]*NovelSearch, error) {
	var novels []*NovelSearch
	c := colly.NewCollector()

	c.OnHTML("div.row", func(e *colly.HTMLElement) {
		e.ForEach("div.col-md-3", func(_ int, el *colly.HTMLElement) {
			href := el.ChildAttr("a", "href")
			title := el.ChildAttr("a", "title")
			image := el.ChildAttr("img", "src")

			info, err := novelInfo(href)
			if err != nil {
				fmt.Printf("Failed to retrieve novel information for %s: %s\n", title, err)
				return
			}

			novel := &NovelSearch{
				Title:       title,
				Href:        href,
				Image:       image,
				Author:      info.Author,
				Genres:      info.Genres,
				Status:      info.Status,
				Views:       info.Views,
				Description: info.Description,
			}
			novels = append(novels, novel)
		})
	})

	if err := c.Visit(root + "?s=" + strings.ReplaceAll(query, " ", "+")); err != nil {
		return nil, err
	}

	return novels, nil
}

func novelInfo(novelLink string) (*NovelSearch, error) {
    c := colly.NewCollector()
    var info *NovelSearch

    c.OnHTML("div.col-xs-12.col-info-desc", func(e *colly.HTMLElement) {
        description := e.ChildText("div.desc-text")

        infoText := e.ChildText("div.info")
        infoLines := strings.Split(infoText, "\n")

        var parsedInfo [][]string
        for _, line := range infoLines {
            line = strings.TrimSpace(line)
            if line != "" {
                parts := strings.Split(line, ":")
                parsedInfo = append(parsedInfo, parts)
            }
        }

        if len(parsedInfo) >= 4 {
            info = &NovelSearch{
                Author:      strings.TrimSpace(parsedInfo[0][1]),
                Genres:      strings.TrimSpace(parsedInfo[1][1]),
                Status:      strings.TrimSpace(parsedInfo[2][1]),
                Views:       strings.TrimSpace(parsedInfo[3][1]),
                Description: strings.TrimSpace(description),
            }
        }
    })

    if err := c.Visit(novelLink); err != nil {
        return nil, err
    }

    return info, nil
}


func ProcessQuery(query string) ([]*NovelSearch, error) {
	results, err := novelSearcher(query)
	if err != nil {
		return nil, err
	}

	return results, nil
}
