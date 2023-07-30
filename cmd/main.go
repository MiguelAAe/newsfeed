package main

import (
	"fmt"
	"log"
	"net/http"
	"newsfeed/pkg/api"
	"newsfeed/pkg/rsspuller"
	"time"

	_ "newsfeed/docs"
)

func main() {
	bbcFeed := &rsspuller.Feed{
		Name:        "bbc",
		URL:         "http://feeds.bbci.co.uk/news/uk/rss.xml",
		Category:    "news",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.BBCFeedParser,
	}

	bbcFeedTech := &rsspuller.Feed{
		Name:        "bbc",
		URL:         "http://feeds.bbci.co.uk/news/technology/rss.xml",
		Category:    "technology",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.BBCFeedParser,
	}

	skyFeed := &rsspuller.Feed{
		Name:        "sky",
		URL:         "https://feeds.skynews.com/feeds/rss/uk.xml",
		Category:    "news",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.SkyFeedParser,
	}

	skyFeedTech := &rsspuller.Feed{
		Name:        "sky",
		URL:         "https://feeds.skynews.com/feeds/rss/technology.xml",
		Category:    "technology",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.SkyFeedParser,
	}

	skyFeedPolitics := &rsspuller.Feed{
		Name:        "sky",
		URL:         "https://feeds.skynews.com/feeds/rss/politics.xml",
		Category:    "politics",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.SkyFeedParser,
	}

	skyFeedUS := &rsspuller.Feed{
		Name:        "sky",
		URL:         "https://feeds.skynews.com/feeds/rss/us.xml",
		Category:    "us",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.SkyFeedParser,
	}

	skyFeedEntertainment := &rsspuller.Feed{
		Name:        "sky",
		URL:         "https://feeds.skynews.com/feeds/rss/entertainment.xml",
		Category:    "entertainment",
		RefreshTime: 1 * time.Hour,
		Parser:      rsspuller.SkyFeedParser,
	}

	pull := rsspuller.NewPuller()
	defer pull.Close()

	err := pull.Listen(bbcFeed, bbcFeedTech, skyFeed, skyFeedTech, skyFeedPolitics, skyFeedUS, skyFeedEntertainment)
	if err != nil {
		log.Fatal(err)
	}

	router := api.NewAPI(pull)
	log.Printf("Listening on Port:%s\n", "9200")
	err = http.ListenAndServe(fmt.Sprintf(":%s", "9200"), router)
	if err != nil {
		log.Printf("err from router: %v\n", err)
	}
}
