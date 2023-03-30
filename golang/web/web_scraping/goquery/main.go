package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// Request the HTML page.
	res, err := http.Get("https://ediblelandscaping.com/buyPlants.php?func=view&id=156")
	if err != nil {
		fmt.Println("http.Get error:", err)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("http.Get error:", err)
		return
	}

	// Find the review items
	doc.Find("#saleBar").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the name.
		name := s.Find("p.node").Text()
		fmt.Println(i, name)
	})

}
