package usecases

import (
	"fmt"
	"github.com/arizard/nine-publishing-technical-test/entities"
	"log"
)

// GetArticle is a use case which retrieves an article.
// Dependencies are injected when the struct is initialized.
type GetArticle struct {
	ArticleRepository entities.ArticleRepository
	ArticleID         string
	Response          *ResponseCollector
}

// Execute is a method to action the use case using the injected dependencies.
func (uc GetArticle) Execute() {
	defer panicHandler(uc.Response)

	res, err := uc.ArticleRepository.Get(uc.ArticleID)

	if err != nil {
		resp := ResponseError{
			Name:        "ERROR_REPOSITORY_GET",
			Description: fmt.Sprintf("%s", err),
		}

		uc.Response.SetError(&resp)

		return
	}

	log.Printf("use case GetArticle executed. %v", res)

	resp := Response{
		Body: map[string]interface{}{
			"article": res,
		},
	}

	uc.Response.SetResponse(&resp)

}
