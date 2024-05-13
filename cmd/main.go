package main

import (
	"fmt"

	"github.com/drizzleent/scraper/internal/service"
	"github.com/drizzleent/scraper/internal/service/scraper"
	"github.com/gocolly/colly"
)

const (
	userAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0"
)

func main() {
	var srv service.Scraper

	c := colly.NewCollector(colly.UserAgent(userAgent))
	srv = scraper.NewScraper(c)

	err := srv.Scrap()
	if err != nil {
		fmt.Println(err)
	}

}
