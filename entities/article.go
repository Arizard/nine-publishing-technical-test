package entities

// Article is an entity which models the Article
type Article struct {
	id string
	title string
	date string
	body string
	tags []string
}

// ArticleRepository implements the repository model for Article
type ArticleRepository interface {
	Add(Article) repositoryError
	Get(Article) (Article, repositoryError)
	Find(Article) ([]Article, repositoryError)
}