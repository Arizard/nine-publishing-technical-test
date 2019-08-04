package presenters

import (
	"github.com/arizard/nine-publishing-technical-test/entities"
	"log"
	"encoding/json"
)

// All the JSON presenter should do is take a struct as an argument, and then
// output the correct JSON formatted string result.

// JSONPresenter presents data in browser-renderable html.
type JSONPresenter struct {}

// Deserialize converts a string body to a mapping of string to string.
func (JSONPresenter) Deserialize(body []byte) map[string]interface{} {
	deserial := map[string]interface{}{}
	err := json.Unmarshal([]byte(body), &deserial)

	if err != nil {
		log.Printf("failed to deserialize: \"%s\"", err)
	}

	return deserial
}

// NotFound presents the 404 output.
func (JSONPresenter) NotFound() string {
	return "{ \"error\": \"Not Found\" }"
}


// Forbidden presents the 403 output.
func (JSONPresenter) Forbidden() string {
	return "{ \"error\": \"Forbidden\" }"
}

// InternalServerError displays the 500 output.
func (JSONPresenter) InternalServerError() string {
	return "{ \"error\": \"Internal Server Error.\" }"
}

// SubmitArticle is a presenter which renders the input arguments as a JSON
// string.
func (JSONPresenter) SubmitArticle(info map[string]interface{}) string {
	json, err := json.Marshal(info)

	if err != nil {
		log.Printf("error serializing info: %s", err)
		return ""
	}

	return string(json)
}

// GetArticle is a presenter which renders the input arguments as a JSON
// string.
func (JSONPresenter) GetArticle(article entities.Article) string {
	json, err := json.Marshal(article)

	if err != nil {
		log.Printf("error serializing article: %s", err)
		return ""
	}

	return string(json)
}

// GetArticlesByTag is a presenter which renders the input arguments as a JSON
// string.
func (JSONPresenter) GetArticlesByTag(articles []entities.Article) string {
	json, err := json.Marshal(articles)

	if err != nil {
		log.Printf("error serializing articles: %s", err)
		return ""
	}

	log.Printf("presenter %s", json)

	return string(json)
}

