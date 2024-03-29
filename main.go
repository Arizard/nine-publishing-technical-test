package main

import (
	"github.com/arizard/nine-publishing-technical-test/handlers"
	"github.com/arizard/nine-publishing-technical-test/infrastructure"
	"github.com/arizard/nine-publishing-technical-test/presenters"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Printf("setting up")

	jsonPresenter := presenters.JSONPresenter{}

	r := mux.NewRouter().StrictSlash(false)

	JSONHandler := handlers.Handler{
		ContentType:       "application/json",
		Presenter:         jsonPresenter,
		ArticleRepository: infrastructure.NewInMemoryArticleRepository(),
	}

	r.NotFoundHandler = http.HandlerFunc(JSONHandler.NotFoundHandler)

	r.HandleFunc("/articles", JSONHandler.CORSWrapper(JSONHandler.SubmitArticleHandler)).Methods("POST", "OPTIONS")
	r.HandleFunc("/articles/{id}", JSONHandler.CORSWrapper(JSONHandler.GetArticleHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/tags/{tagName}/{date}", JSONHandler.CORSWrapper(JSONHandler.GetArticlesByTagHandler)).Methods("GET", "OPTIONS")

	log.Printf("starting server")

	http.ListenAndServe(":8080", r)

}
