package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Company string `csv:"company"`
	Price   string `csv:"price"`
	Change  string `csv:"change"`
}

func Scrape(c *colly.Collector, ticker string, stocks *[]Stock) {
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {

		var stock Stock

		stock.Company = e.ChildText("h1")
		fmt.Println("company : ", stock.Company)

		stock.Price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("price : ", stock.Price)

		stock.Change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("change : ", stock.Change)

		*stocks = append(*stocks, stock)

	})

	c.Visit("https://finance.yahoo.com/quote/" + ticker + "?p=" + ticker)

}

func main() {
	c := colly.NewCollector()

	TICKER := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"TSLA",
		"MCD",
		"COIN",
		"MCD",
		"PLTR",
		"EL",
		"UBER",
		"CAT",
		"GRRR",
		"NVTA",
		"YTEN",
	}

	stocks := []Stock{}
	// Find and visit all links

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for _, ticker := range TICKER {
		Scrape(c, ticker, &stocks)
	}
	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"company",
		"price",
		"change",
	}

	writer.Write(headers)

	for _, stock := range stocks {

		recorder := []string{
			stock.Company,
			stock.Price,
			stock.Change,
		}

		writer.Write(recorder)
	}

	defer writer.Flush()

}
