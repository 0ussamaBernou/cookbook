package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf(" %s took %v\n", name, time.Since(start))
	}
}
func main() {
	defer timer("main")()
	c := colly.NewCollector()

	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	c.OnRequest(func(req *colly.Request) {
		fmt.Println("visiting: ", req.URL)
	})

	c.Visit("https://books.toscrape.com/catalogue/page-1.html")

}
