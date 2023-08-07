package rsspuller

import (
	"fmt"
	"sort"
)

type Puller struct {
	feedSources map[string]map[string]*Feed //map of source to map of category feed
	initialised bool
	stopped     bool
	done        chan bool
}

func NewPuller() *Puller {
	return &Puller{
		done:        make(chan bool),
		feedSources: make(map[string]map[string]*Feed),
	}
}

func (p *Puller) Listen(feeds ...*Feed) error {
	p.initialised = true
	for _, feed := range feeds {
		// populate the cancel channel
		feed.done = p.done

		err := feed.Listen()
		if err != nil {
			return fmt.Errorf("failed to listen feed: %s, %w", feed.Name, err)
		}

		// pull struct work to stay up to date
		// p.feeds = append(p.feeds, &feed)

		// check if the name of the feed is already in

		cat, ok := p.feedSources[feed.Name]
		if !ok {
			// pupulate map of maps if it doesn't exist
			catMap := make(map[string]*Feed)
			catMap[feed.Category] = feed

			p.feedSources[feed.Name] = catMap
			continue
		}

		// check if the category already exists
		_, ok = p.feedSources[feed.Name][feed.Category]
		if !ok {
			// add feed to database
			cat[feed.Category] = feed
			p.feedSources[feed.Name] = cat
		}

		// if the feed is repeated do nothing
		// if we update it we would need to serialise access to its share variables
		continue
	}
	return nil
}

func (p *Puller) GetAllFeeds() []Item {
	// todo: opportunity to cache

	var resp []Item

	for _, categoryFeedMap := range p.feedSources {
		for _, feed := range categoryFeedMap {
			resp = append(resp, feed.GetFeed().Item...)
		}
	}

	// short feeds
	sort.Sort(ByItem(resp))
	return resp
}

func (p *Puller) GetFeedsByCategory(category ...string) []Item {
	// todo: opportunity to cache

	var resp []Item

	for _, categoryFeedMap := range p.feedSources {

		for _, catName := range category {
			feed, ok := categoryFeedMap[catName]
			if !ok {
				continue
			}

			resp = append(resp, feed.GetFeed().Item...)
		}
	}
	// short feeds
	sort.Sort(ByItem(resp))
	return resp
}

func (p *Puller) GetFeedsByName(name ...string) []Item {
	// todo: opportunity to cache

	var resp []Item

	for _, nameToLook := range name {
		feedCat, ok := p.feedSources[nameToLook]
		if !ok {
			continue
		}

		for _, feed := range feedCat {
			resp = append(resp, feed.GetFeed().Item...)
		}
	}

	// short feeds
	sort.Sort(ByItem(resp))
	return resp
}

func (p *Puller) SearchFeeds(condition map[string][]string) []Item {
	// todo: opportunity to cache

	// iterate condition map
	// find source
	// iterate through cat and map sources
	// repeat until the end

	var resp []Item

	for sourceName, categories := range condition {
		categoryFeedMap, ok := p.feedSources[sourceName]
		if !ok {
			continue
		}

		for _, catName := range categories {
			feed, ok := categoryFeedMap[catName]
			if !ok {
				continue
			}

			resp = append(resp, feed.GetFeed().Item...)
		}

	}

	// short feeds
	sort.Sort(ByItem(resp))
	return resp
}

func (p *Puller) ListFeedCategories() []string {
	categoriesMap := make(map[string]bool)

	for _, catMap := range p.feedSources {
		for keyName, _ := range catMap {
			categoriesMap[keyName] = true

		}
	}

	categoriesResp := make([]string, len(categoriesMap))
	i := 0
	for keyName, _ := range categoriesMap {
		categoriesResp[i] = keyName
		i++
	}

	return categoriesResp
}

func (p *Puller) ListFeedNames() []string {
	sourceNames := make([]string, len(p.feedSources))

	i := 0
	for keyName, _ := range p.feedSources {
		sourceNames[i] = keyName
		i++
	}

	return sourceNames
}

func (p *Puller) Close() {
	if p.initialised == false || p.stopped == false {
		return
	}
	p.done <- true
}
