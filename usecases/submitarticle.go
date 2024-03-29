package usecases

import (
	"fmt"
	"github.com/arizard/nine-publishing-technical-test/entities"
	"log"
	"time"
)

// SubmitArticle is a use case which performs the submission of an article.
// Dependencies are injected when the struct is initialized.
type SubmitArticle struct {
	ArticleRepository entities.ArticleRepository
	ArticleData       map[string]interface{}
	Response          *ResponseCollector
}

// Execute is a method to action the use case using the injected dependencies.
func (uc SubmitArticle) Execute() {
	defer panicHandler(uc.Response)

	tags := make([]string, 0)

	for _, val := range uc.ArticleData["tags"].([]interface{}) {
		tags = append(tags, fmt.Sprintf("%s", val))
	}

	newArticle := entities.NewArticle(
		uc.ArticleData["id"].(string),
		uc.ArticleData["title"].(string),
		uc.ArticleData["date"].(string),
		uc.ArticleData["body"].(string),
		tags,
		time.Now().UTC(),
	)

	addErr := uc.ArticleRepository.Add(newArticle)

	if addErr != nil {
		log.Printf("Error during SubmitArticle use case: %s", addErr)

		resp := ResponseError{
			Name:        "ERROR_REPOSITORY_ADD",
			Description: fmt.Sprintf("%s", addErr),
		}

		uc.Response.SetError(&resp)
	}

	log.Printf("use case SubmitArticle executed.")

	resp := Response{
		Body: map[string]interface{}{
			"success": "Submitted Article",
		},
	}

	uc.Response.SetResponse(&resp)
}
