package main

import (
	"fmt"
)

func main() {
	var result []string = Scrape("https://zenn.dev/books/explore")
	fmt.Println(result)
}