package usecases

import (
	"log"
	// "github.com/arizard/nine-publishing-technical-test/entities"
)

// SubmitArticle is a use case which performs the submission of an article.
// Dependencies are injected when the struct is initialized.
type SubmitArticle struct {
	ArticleData map[string]string
	Response *ResponseCollector
}

// Execute is a method to action the use case using the injected dependencies.
func (uc SubmitArticle) Execute() {
	log.Printf("use case SubmitArticle executed.")

	resp := Response{
		Body: map[string]string{
			"test": "Hello World!",
		},
	}

	uc.Response.SetResponse(&resp)

}