package presenters

import (
	"github.com/arizard/nine-publishing-technical-test/entities"
)

// Presenter defines the contract for presenters, either html or json.
type Presenter interface {
	Deserialize([]byte) map[string]interface{}
	NotFound() string
	InternalServerError() string
	Forbidden() string
	SubmitArticle(map[string]interface{}) string
	GetArticle(entities.Article) string
	GetArticlesByTag([]entities.Article) string
}
