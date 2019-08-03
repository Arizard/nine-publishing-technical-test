package main

import (
	"github.com/arizard/nine-publishing-technical-test/presenters"
	"github.com/arizard/nine-publishing-technical-test/handlers"
	"github.com/arizard/nine-publishing-technical-test/infrastructure"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	log.Printf("setting up")

	jsonPresenter := presenters.JSONPresenter{}

	r := mux.NewRouter().StrictSlash(false)

	JSONHandler := handlers.Handler{
		ContentType: "application/json",
		Presenter: jsonPresenter,
		ArticleRepository: infrastructure.NewInMemoryArticleRepository(),
	}

	r.NotFoundHandler = http.HandlerFunc(JSONHandler.NotFoundHandler)

	r.HandleFunc("/articles", JSONHandler.CORSWrapper(JSONHandler.SubmitArticleHandler)).Methods("POST", "OPTIONS")
	r.HandleFunc("/article/{id}", JSONHandler.CORSWrapper(JSONHandler.GetArticleHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/tag/{tagName}/{date}", JSONHandler.CORSWrapper(JSONHandler.GetArticleByTagHandler)).Methods("GET", "OPTIONS")

	log.Printf("starting server")

	http.ListenAndServe(":8080", r)


}