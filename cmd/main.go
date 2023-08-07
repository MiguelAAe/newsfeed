package main

import (
	"context"
	"log"
	"net/http"
	"newsfeed/pkg/api"
	"newsfeed/pkg/rsspuller"
	"os"
	"os/signal"
	"syscall"
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

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":9200",
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for a signal to shutdown the server
	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	log.Printf("Shutting down server gracefully")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}
	log.Printf("Successful server shutdown")
}
