package scrape

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestScrape(t *testing.T) {
    go func() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            html := `
                <!DOCTYPE html>
                <html>
                <head>
                    <title>Test Page</title>
                </head>
                <body>
                    <article>
                        <a>
                            <h3>Test Book</h3>
                        </a>
                        <div>
                            <a>Test Author</a>
                        </div>
                    </article>
                </body>
                </html>
            `
            w.Write([]byte(html))
        })

        http.ListenAndServe(":8080", nil)
    }()

    // 少し待ってからスクレイピングを開始
    time.Sleep(time.Second)

	var result[] string = Scrape("http://localhost:8080")
	var expected string = "1 著者: Test Author / タイトル: Test Book\n"
    fmt.Println(result)
	if result[0] != expected {
		t.Errorf("got %v\nwant %v", result, expected)
	}
}