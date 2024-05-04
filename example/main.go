package main

import (
	"fmt"

	scrape "github.com/ei-sugimoto/webscrape"
)

func main() {
    var result []string = scrape.Scrape("https://zenn.dev/books/explore")
    fmt.Println(result)
}