package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func parseHTML(html io.Reader) ([]Product, error) {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, err
	}

	products := []Product{}

	doc.Find(".ui-search-layout__item").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find("h3.poly-component__title-wrapper a").Attr("href")
		if !ok {
			link = ""
		}
		product := Product{
			Title:    s.Find("a.poly-component__title").Text(),
			Price:    s.Find("div.poly-price__current span.andes-money-amount__fraction").First().Text(),
			Discount: s.Find("span.andes-money-amount__discount").First().Text(),
			Link:     link,
		}
		products = append(products, product)
	})

	return products, nil
}
