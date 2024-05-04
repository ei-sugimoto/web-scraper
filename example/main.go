package main

import (
	"fmt"

	scrape "github.com/ei-sugimoto/webscraper"
)

func main() {
    var result []string = scrape.Scrape("https://zenn.dev/books/explore")
    fmt.Println(result)
}