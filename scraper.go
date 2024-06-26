package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/rpstvs/BlogAggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timebetweenRequests time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timebetweenRequests, concurrency)

	ticker := time.NewTicker(timebetweenRequests)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Couldn't get next feeds to fetch", err)
			continue
		}

		log.Printf("Found %v feeds to fetch!", len(feeds))

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Couldnt mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(feed.Url)

	if err != nil {
		log.Printf("Couldn't Collect feed %s: %v", feed.Url, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		log.Println("Found post.", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubdate"`
}

func fetchFeed(url string) (*RSSFeed, error) {

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil

}
