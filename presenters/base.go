package presenters

import (
	// "github.com/arizard/script-engine-server/entities"
)

// Presenter defines the contract for presenters, either html or json.
type Presenter interface {
	Deserialize([]byte) map[string]string
	NotFound() string
	InternalServerError() string
	Forbidden() string
	Index() string
	SubmitArticle() string
	GetArticle() string
	GetArticleByTag() string
}
