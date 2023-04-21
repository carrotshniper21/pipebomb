// pipebomb/anime/animesearch.go
package anime

import (
	"github.com/gocolly/colly"
)

const root = "https://zoro.to"

func animeSearcher() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		elem.Attr("href")
	})

	c.Visit("https://zoro.to/search?keyword=panty+stocking")
}
