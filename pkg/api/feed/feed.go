package feed

import (
	"encoding/json"
	"io"
	"net/http"
	"newsfeed/pkg/rsspuller"
)

type RSSFeed struct {
	FeedService *rsspuller.Puller
}

type feedQuery struct {
	Query map[string][]string
}

// GetAllFeed
//
//	@Summary		Gets a list of all feeds
//	@Description	Gets a list of all feeds
//	@Accept			json
//	@Produce		json
//	@Param			query		body		feedQuery	false	"a map of sources to list of categories"
//	@Param			source		query		string		false	"source				of the feed"
//	@Param			category	query		string		false	"category	of	the	feed"
//	@Success		200			{object}	feedQuery	"a list of feeds"
//	@Failure		404
//	@Router			/newsFeed [get]
func (ws *RSSFeed) GetAllFeed(w http.ResponseWriter, r *http.Request) {
	var content feedQuery

	err := json.NewDecoder(r.Body).Decode(&content)
	if err != io.EOF && err != nil {
		// todo: log this on console
		return
	}

	if content.Query != nil {
		resp := ws.FeedService.SearchFeeds(content.Query)

		if resp == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			// log error
			return
		}
		return
	}

	source := r.URL.Query().Get("source")
	category := r.URL.Query().Get("category")

	if source == "" && category == "" {
		resp := ws.FeedService.GetAllFeeds()

		if resp == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			// log error
			return
		}
		return
	} else if source != "" && category != "" {
		condition := make(map[string][]string)
		condition[source] = []string{category}

		resp := ws.FeedService.SearchFeeds(condition)
		if resp == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			// log error
			return
		}
		return
	}

	if source != "" {
		resp := ws.FeedService.GetFeedsByName(source)
		if resp == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			// log error
			return
		}
		return
	}

	if category != "" {
		resp := ws.FeedService.GetFeedsByCategory(category)
		if resp == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			// log error
			return
		}
		return
	}
}

type category struct {
	Categories []string `json:"categories"`
}

// GetFeedCategories
//
//	@Summary		Gets a list of all feed categories
//	@Description	Gets a list of all feed categories
//	@Produce		json
//	@Success		200	{object}	category	"a list of categories"
//	@Router			/newsFeed/categories [get]
func (ws *RSSFeed) GetFeedCategories(w http.ResponseWriter, r *http.Request) {
	resp := ws.FeedService.ListFeedCategories()

	cat := make([]string, len(resp))
	for i, c := range resp {
		cat[i] = c
	}

	err := json.NewEncoder(w).Encode(&category{Categories: cat})
	if err != nil {
		// log error
		return
	}
}

type names struct {
	Names []string `json:"names"`
}

// GetFeedNames
//
//	@Summary		Gets a list of all feed names
//	@Description	Gets a list of all feed names
//	@Produce		json
//	@Success		200	{object}	names	"a list of names"
//	@Router			/newsFeed/names [get]
func (ws *RSSFeed) GetFeedNames(w http.ResponseWriter, r *http.Request) {
	resp := ws.FeedService.ListFeedNames()

	names := make([]string, len(resp))
	for i, name := range resp {
		names[i] = name
	}

	err := json.NewEncoder(w).Encode(&category{Categories: names})
	if err != nil {
		// log error
		return
	}
}
