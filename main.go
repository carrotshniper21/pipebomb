// main.go
package main

import (
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	visitedLinks := sync.Map{}
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(elem *colly.HTMLElement) {
		processLink(elem, &visitedLinks)
	})

	c.Visit("https://flixhq.to/search/boob")
}
