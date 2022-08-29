package main

import (
	"github.com/go-chi/chi"
	"github.com/vfunin/elastic/handler"
	"github.com/vfunin/elastic/l"
	"github.com/vfunin/elastic/store"
)

// Переписать не на Martini
func main() {
	//Sentry error handler
	//sentry.Init(sentry.Client(os.Getenv("SENTRY_DSN")))
	//Initialize Stores
	articleStore, err := store.NewArticleStore()
	parseErr(err)
	//Initialize Handlers
	articleHandler := handler.NewArticleHandler(articleStore)
	//Initialize Router

	r := chi.NewRouter()
	//
	//r.Get("/article/id/:id", articleHandler.Id)
	//r.Post("/article/add", articleHandler.Add)
	//r.Post("/article/search", articleHandler.Search)
	//	panicHandler := handler.PanicHandler{}
	//r.Get("/panic", panicHandler.Handle)
	//r.Post("/log/add", panicHandler.Log)
}

func parseErr(err error) {
	if err != nil {
		l.F(err)
	}
	l.Log.Log("Application started")
}
