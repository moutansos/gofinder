package main

import (
	"time"
	"log"
	"github.com/mmcdole/gofeed"
)

type listing struct {
	Title string
	Url string
}

func runFinder(itemsChannel chan<- listing) {
	for true {
		go openFeed(itemsChannel)
		time.Sleep(5 * time.Second)
	}
}

func openFeed(itemsChannel chan<- listing) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://spokane.craigslist.org/search/apa?format=rss")
	if err != nil {
		log.Println("Error parsing feed: ", err)
	}

	listings := feed.Items

	for _, item := range listings {
		newListing := *item
		log.Printf("%s: %s", newListing.Title, newListing.Link)
		itemsChannel <- listing{
			Title: item.Title,
			Url: item.Link,
		}
	}
}