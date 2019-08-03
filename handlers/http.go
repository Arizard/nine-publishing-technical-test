package handlers

import (
	"io/ioutil"
	"log"
	"github.com/arizard/nine-publishing-technical-test/usecases"
	"github.com/arizard/nine-publishing-technical-test/presenters"
	"github.com/arizard/nine-publishing-technical-test/entities"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

// The handler should handle HTTP headers and status codes, and execute the
// use cases, then run the output through the presenter.

// Handler is a struct which implements methods that take the 
// ResponseWriter and Request objects as arguments, such as from an
// HTTP request. It is used to decouple the Drivers layer from the
// Controllers and Presenters.
type Handler struct {
	ContentType string
	Presenter presenters.Presenter
	ArticleRepository entities.ArticleRepository
}

// CORSWrapper is used to allow cross-origin requests on a handler by wrapping
// the handler function.
func (handler Handler) CORSWrapper(hf func (http.ResponseWriter, *http.Request)) func (http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", handler.ContentType)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cache-Control")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if (r.Method == "OPTIONS") {
			log.Printf("incoming CORS request")
		} else {
			log.Printf("incoming request")
			w.Header().Set("Content-Type", handler.ContentType)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
			hf(w, r)
		}
	}
	
}


// NotFoundHandler handles 404s
func (handler Handler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, handler.Presenter.NotFound())
}

// InternalServerErrorHandler handles 500s
func (handler Handler) InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintf(w, handler.Presenter.InternalServerError())
}

// SubmitArticleHandler handles the POST request to submit an article.
func (handler Handler) SubmitArticleHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	parsedBody := handler.Presenter.Deserialize(body)

	log.Printf("request parsed: %s", parsedBody)

	resp := usecases.ResponseCollector{}

	uc := usecases.SubmitArticle{
		ArticleRepository: handler.ArticleRepository,
		ArticleData: parsedBody,
		Response: &resp,
	}

	uc.Execute()
	
	log.Printf("response collector received %s", resp.Response.Body)

	fmt.Fprint(w, handler.Presenter.SubmitArticle())
}

// GetArticleHandler handles the GET request to retrieve an article.
func (handler Handler) GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501) // Not implemented
	
	id := mux.Vars(r)["id"]

	log.Printf("handler GetArticleHandler var id: %s", id)
}

// GetArticleByTagHandler handles the GET request for retrieving articles with 
// the tagName and date parameter.
func (handler Handler) GetArticleByTagHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501) // Not implemented

	tagName := mux.Vars(r)["tagName"]
	date := mux.Vars(r)["date"]

	log.Printf("handler GetArticleHandler var tagName: %s, date: %s", tagName, date)
}