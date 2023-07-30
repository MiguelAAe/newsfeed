package api

import (
	"net/http"
	"newsfeed/pkg/api/feed"
	"newsfeed/pkg/rsspuller"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewAPI(puller *rsspuller.Puller) *chi.Mux {
	//  set up router
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	handler := feed.RSSFeed{FeedService: puller}

	r.Route("/newsFeed", func(r chi.Router) {
		r.Get("/", handler.GetAllFeed)
		r.Get("/categories", handler.GetFeedCategories)
		r.Get("/names", handler.GetFeedNames)
	})

	r.Get("/swagger/*", httpSwagger.Handler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world, welcome to a news feed"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("all very healthy here"))
	})

	return r
}
