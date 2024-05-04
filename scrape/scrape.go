package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Scrape(url string) []string {
	c := colly.NewCollector()
    i := 0
    var results []string

    c.OnHTML("article", func(e *colly.HTMLElement) {
        i++
        book := e.DOM.Find("a > h3").Text()
        author := e.DOM.Find("div > a").Text()
        if author == "" {
            author = e.ChildText(".BookLink_userName__avtjq")
        }

        results = append(results, fmt.Sprintf("%d 著者: %s / タイトル: %s\n", i, author, book))
    })

    c.Visit(url)

    return results
}