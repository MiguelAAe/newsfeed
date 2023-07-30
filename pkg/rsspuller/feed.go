package rsspuller

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

type Feed struct {
	URL         string
	RefreshTime time.Duration
	done        chan bool
	Name        string
	Category    string
	Parser      FeedParser
	data        FeedContent
	read        chan chan FeedContent
}

type FeedContent struct {
	URLSource string
	Name      string
	Category  string
	Item      []Item
}

type Item struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Media       Media     `json:"media"`
	PubDate     time.Time `json:"pubDate"`
}

type Media struct {
	Link   string `json:"link"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

// ByItem implements sort.Interface for []Item based on
// the PubDate field.
type ByItem []Item

func (a ByItem) Len() int {
	return len(a)
}

func (a ByItem) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByItem) Less(i, j int) bool {
	// we do by after since we want the latest first
	return a[i].PubDate.After(a[j].PubDate)
}

func (f *Feed) Listen() error {
	dataResp, err := f.Parser(f.URL)
	if err != nil {
		return err
	}

	f.data = dataResp
	ticker := time.NewTicker(f.RefreshTime)

	f.read = make(chan chan FeedContent)

	// attempt to make connection
	go func() {
		for {
			select {
			case <-f.done:
				return
			case t := <-ticker.C:
				fmt.Printf("refreshed after %v\n", t)
				dataResp, err = f.Parser(f.URL)
				if err != nil {
					// return err how will I handle this?
				}

				f.data = dataResp

			case r := <-f.read:
				r <- dataResp

			}
		}
	}()

	return nil
}

func (f *Feed) GetFeed() FeedContent {
	// check if channel is nil
	if f.read == nil {
		fmt.Println("empty channel")
		return FeedContent{}
	}

	resp := make(chan FeedContent)

	select {
	// check if channel is closed
	case _, ok := <-f.read:
		if !ok {
			return FeedContent{}
		}
	default:
		f.read <- resp
	}

	return <-resp
}

type FeedParser func(string) (FeedContent, error)

func BBCFeedParser(url string) (FeedContent, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return FeedContent{}, err
	}

	var resp FeedContent
	resp.URLSource = feed.Link
	resp.Name = feed.Title

	itemsResp := make([]Item, len(feed.Items))

	for i, item := range feed.Items {
		itemsResp[i].Description = item.Description
		itemsResp[i].Link = item.Link
		itemsResp[i].Title = item.Title
		itemsResp[i].PubDate = *item.PublishedParsed
	}

	resp.Item = itemsResp

	return resp, nil
}

func SkyFeedParser(url string) (FeedContent, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return FeedContent{}, err
	}

	var resp FeedContent
	resp.URLSource = feed.Link
	resp.Name = feed.Title

	itemsResp := make([]Item, len(feed.Items))

	for i, item := range feed.Items {
		itemsResp[i].Description = item.Description
		itemsResp[i].Link = item.Link
		itemsResp[i].Title = item.Title
		itemsResp[i].PubDate = *item.PublishedParsed
		itemsResp[i].Media = mediaFinder(item)
	}

	resp.Item = itemsResp

	return resp, nil
}

func mediaFinder(item *gofeed.Item) Media {
	maps, ok := item.Extensions["media"]
	if !ok {
		// todo something else
		return Media{}
	}

	thumbnailsList, ok := maps["thumbnail"]
	if !ok {
		// todo something else
		return Media{}
	}

	if len(thumbnailsList) < 1 {
		return Media{}
	}

	attribute := thumbnailsList[0].Attrs

	var media Media

	media.Link = attribute["url"]
	media.Width = attribute["width"]
	media.Height = attribute["height"]

	return media
}
