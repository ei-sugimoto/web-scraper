# web-scraper
[![test](https://github.com/ei-sugimoto/webscraper/actions/workflows/test.yml/badge.svg)](https://github.com/ei-sugimoto/webscraper/actions/workflows/test.yml) [![codecov](https://codecov.io/github/ei-sugimoto/webscraper/graph/badge.svg?token=XHRUJ4OVA8)](https://codecov.io/github/ei-sugimoto/webscraper) [![.github/workflows/lint.yml](https://github.com/ei-sugimoto/webscraper/actions/workflows/lint.yml/badge.svg)](https://github.com/ei-sugimoto/webscraper/actions/workflows/lint.yml)
## install

```bash
go get "github.com/ei-sugimoto/webscraper"
```
## example
```go
package main

import (
	"fmt"

	scrape "github.com/ei-sugimoto/webscraper"
)

func main() {
    var result []string = scrape.Scrape("https://zenn.dev/books/explore")
    fmt.Println(result)
}
```
```bash:output
go run main.go
[1 著者: y_ta@大学生Webエンジニア0から始めるFirebase入門¥027人がいいね / タイトル: 0から始めるFirebase入門
 2 著者: ytaisei（たいせー）TanStack Query　〜プロダクトで採用するための勘所〜¥0192人がいいね / タイトル: TanStack Query　〜プロダクトで採用するための勘所〜
 3 著者: PADAone🐕TypeScriptの代数的部分型模型¥0230人がいいね / タイトル: TypeScriptの代数的部分型模型
 4 著者: 中條 剛（ちゅーやん）仕組みから理解する Riverpod¥1,200143人がいいね / タイトル: 仕組みから理解する Riverpod
...
]
```