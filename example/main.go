package main

import (
	"fmt"
    "github.com/ei-sugimoto/webscraper"
)

func main() {
    var result []string = .Scrape("https://zenn.dev/books/explore")
    fmt.Println(result)
}