package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"log"
	"github.com/gocolly/colly"
)

type stock struct{
	company,price,change string
}

func main() {
	ticker := []string{
		"MSFT",
		"GOOG",
		"AMZN",
		"FB",
		"AAPL",
		"INTC",
		"CSCO",
		"CMCSA",
		"PEP",
		"ADBE",
		"NFLX",
		"PYPL",
		"COST",
		"NVDA",
		"AMGN",
		"AVGO",
		"TXN",
		"CHTR",
	}

	// Initialize stocks slice only once
	stocks := []stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong :(", err)
	})
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {

		s := stock{}

		s.company = e.ChildText("h1")
		fmt.Println("Company: ", s.company)
		s.price = e.ChildText("fin-streamer[data-test=qsp-price]")
		fmt.Println("Price: ", s.price)
		s.change = e.ChildText("fin-streamer[data-field=regularMarketChangePercent]")
		fmt.Println("Change: ", s.change)

		stocks = append(stocks, s)
	})

	c.Wait()

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create csv file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	headers := []string{
		"Company",
		"Price",
		"Change",
	}
	writer.Write(headers)
	for _, s := range stocks {
		record := []string{
			s.company,
			s.price,
			s.change,
		}
		writer.Write(record)
	}
	defer writer.Flush()
}