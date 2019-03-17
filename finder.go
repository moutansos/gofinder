package main

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

func runFinder(itemsChannel chan<- ListingRecord, db Database) {
	for true {
		go openFeed(itemsChannel, db)
		time.Sleep(10 * time.Minute)
	}
}

func openFeed(itemsChannel chan<- ListingRecord, db Database) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://spokane.craigslist.org/search/apa?availabilityMode=0&format=rss&max_bedrooms=2&max_price=800&min_bedrooms=1")
	if err != nil {
		log.Println("Error parsing feed: ", err)
	}

	listings := feed.Items

	for _, item := range listings {
		if db.RecordExists(item.Link) {
			continue
		}

		newListing := *item
		record := ListingRecord{
			Title: item.Title,
			Url:   item.Link,
		}

		db.AddRecord(record)

		log.Printf("%s: %s", newListing.Title, newListing.Link)
		itemsChannel <- record
	}
}
