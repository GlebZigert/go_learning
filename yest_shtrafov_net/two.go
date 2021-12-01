package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
	"os"
)

func main() {

    webPage := os.Args[1]
    resp, err := http.Get(webPage)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    title := doc.Find("title").Text()
    fmt.Println(title)
}