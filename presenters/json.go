package presenters

import (
	// "github.com/arizard/nine-publishing-technical-test/entities"
	// "fmt"
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
	return "{ message: \"Not Found\" }"
}


// Forbidden presents the 403 output.
func (JSONPresenter) Forbidden() string {
	return "{ message: \"Forbidden\" }"
}

// InternalServerError displays the 500 output.
func (JSONPresenter) InternalServerError() string {
	return "{ message: \"Internal Server Error.\" }"
}

// Index displays the index output.
func (JSONPresenter) Index() string {
	return "{ message: \"Index\" }"
}

// SubmitArticle is a presenter which renders the input arguments as a JSON
// string.
func (JSONPresenter) SubmitArticle() string {
	return "SubmitArticle"
}

// GetArticle is a presenter which renders the input arguments as a JSON
// string.
func (JSONPresenter) GetArticle() string {
	return "GetArticle"
}

// GetArticleByTag is a presenter which renders the input arguments as a JSON
// string.
func (JSONPresenter) GetArticleByTag() string {
	return "GetArticleByTags"
}

