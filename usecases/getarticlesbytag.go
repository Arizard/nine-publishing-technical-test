package usecases

import (
	"fmt"
	"github.com/arizard/nine-publishing-technical-test/entities"
	"log"
)

// GetArticlesByTag is a use case which retrieves an article.
// Dependencies are injected when the struct is initialized.
type GetArticlesByTag struct {
	ArticleRepository entities.ArticleRepository
	TagName           string
	Date              string
	Response          *ResponseCollector
}

// Execute is a method to action the use case using the injected dependencies.
func (uc GetArticlesByTag) Execute() {
	defer panicHandler(uc.Response)

	res, err := uc.ArticleRepository.Find(uc.TagName, uc.Date, 0)

	if err != nil {
		resp := ResponseError{
			Name:        "ERROR_REPOSITORY_FIND",
			Description: fmt.Sprintf("%s", err),
		}

		uc.Response.SetError(&resp)

		return
	}

	log.Printf("use case GetArticlesByTag executed. %v", res)

	resp := Response{
		Body: map[string]interface{}{
			"articles": res,
		},
	}

	uc.Response.SetResponse(&resp)

}
