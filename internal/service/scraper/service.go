package scraper

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/gocolly/colly"
)

const (
	sberCat    = "https://megamarket.ru/catalog/myaso-i-ptica/"
	sber       = "https://megamarket.ru/"
	pyterkaCat = "https://5ka.ru/special_offers"
	pyterka    = "https://5ka.ru/"
	lentaCat   = "https://lenta.com/catalog/myaso-ptica-kolbasa/"
	lenta      = "https://lenta.com/"
)

type Good struct {
	Url   string
	Name  string
	Image string
	Price string
}

type service struct {
	c *colly.Collector
}

func NewScraper(cl *colly.Collector) *service {
	return &service{
		c: cl,
	}
}

func (s *service) Scrap() error {

	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
		setLentaHeaders(r)
	})

	s.c.OnHTML(".catalog-items-list .catalog-item-mobile", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText(".item-title a"))
	})

	err := s.c.Visit(sberCat)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ScrapLenta() error {
	goods := make([]Good, 0)

	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
		setLentaHeaders(r)
	})

	s.c.OnHTML(".CatalogContainer_skusList__uo1wv .Card_card__1Xuhi", func(h *colly.HTMLElement) {
		url := h.Attr("href")
		name := h.ChildText(".Text_text__B7QH_")
		price := h.ChildText(".PriceText_priceText_discont__IYTcg")
		img := h.ChildAttr(".Image_image__BJqQX", "src")

		good := Good{
			Url:   url,
			Name:  name,
			Image: img,
			Price: price,
		}
		goods = append(goods, good)
	})

	err := s.c.Visit(lentaCat)

	if err != nil {
		return err
	}

	err = s.Save(goods, "./bin/lenta.json")
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Scrap5ka() error {

	goods := make([]Good, 0)

	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
		setPytorochkaHeaders(r)
	})

	s.c.OnHTML(".items-list .item", func(e *colly.HTMLElement) {

		url := e.ChildAttr("product-card item", "href")
		img := e.ChildAttr(".image-cont img", "src")
		name := e.ChildText(".item-name")
		price := e.ChildText(".price-regular")

		good := Good{
			Url:   path.Join(pyterka, url),
			Name:  name,
			Image: img,
			Price: price,
		}
		goods = append(goods, good)
	})

	err := s.c.Visit(pyterkaCat)
	if err != nil {
		return err
	}

	err = s.Save(goods, "./bin/5ka.json")
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Save(p []Good, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func setPytorochkaHeaders(r *colly.Request) {
	r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	r.Headers.Add("Accept-Language", "en-GB,en;q=0.5")
	r.Headers.Add("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Add("Connection", "keep-alive")
	r.Headers.Add("Referer", "https://www.google.com/")
	r.Headers.Add("Upgrade-Insecure-Requests", "1")
	r.Headers.Add("Sec-Fetch-Dest", "document")
	r.Headers.Add("Sec-Fetch-Mode", "navigate")
	r.Headers.Add("Sec-Fetch-Site", "cross-site")
	r.Headers.Add("TE", "trailers")
}

func setLentaHeaders(r *colly.Request) {
	r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	r.Headers.Add("Accept-Language", "en-GB,en;q=0.5")
	r.Headers.Add("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Add("Connection", "keep-alive")
	r.Headers.Add("Cookie", "_utm_referrer=undefined; _userID=undefined; _selectedStoreID=0006; qrator_jsid=1715629691.014.P9izB2g8UEJLhlW4-9j7hisrj58ath82j5q7el5675co2cv3p; .ASPXANONYMOUS=Sx9y35Myuy9HKnAtexbzT3t7tHCWQwPjw617HNIrz7XOdnV9yttsGTnb9YCbMWhmMxVwRxpdmA6-xxXrG4ODowA9JOurkWPOuBSFvXZNP8ejBAr2rR1ShNoy7pCD0VXb48VoPQ2; ASP.NET_SessionId=01npyzphrhdqjqsq1oh0tqbv; CustomerId=858b30e8156043e2a81f89f6f8c057d4; ValidationToken=201754529640977f4620894d104e9002; LNCatalogRoot=true; LNProduct=true; LNStart=true; LNCatalogNodes=true; GACustomerId=ea93a36a2ef3467b8cc5c620df8b8934; ShouldSetDeliveryOptions=True; DontShowCookieNotification=False; cookiesession1=678B28701387C0F30C908FE65598A48E; LastUpdate=2024-05-13; _ymab_param=vAjXKLVxsp9FyZ3lBwabI0XsA8jOrMXo-Rq_hp0zvgWgIEEfbHs_7zYT9OV9ny4SEUiOct1uNsXqizwAP9ilbc9xlWc; _ga=GA1.1.1286963123.1715629694; _gid=GA1.2.719558669.1715629694; _ym_uid=171562969451285772; _ym_d=1715629694; _ym_isad=2; tmr_lvid=ef2e72466ded2d915742164134c8aa68; tmr_lvidTS=1715629695038; domain_sid=M9HUX9f8H1tQzPHCssSbG%3A1715629696074; tmr_detect=0%7C1715629697300; _ga_QB4J0GGLMG=GS1.1.1715632229.2.0.1715632229.0.0.0; _ga_QB4J0GGLM=GS1.1.1715632229.2.0.1715632229.0.0.1016671327")
	r.Headers.Add("Referer", "https://lenta.com/catalog/myaso-ptica-kolbasa/")
	r.Headers.Add("Sec-Fetch-Dest", "empty")
	r.Headers.Add("Sec-Fetch-Mode", "cors")
	r.Headers.Add("Sec-Fetch-Site", "same-origin")
}
