package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func PacktPubScraper() {
	doc, err := goquery.NewDocument("https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".dotd-title h2").Each(func(i int, s *goquery.Selection) {
		free_book := s.Text()
		fmt.Printf("%s\n", strings.TrimSpace(free_book))
	})
}

func main() {
	PacktPubScraper()
}
