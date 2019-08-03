package entities

// Article is an entity which models the Article
type Article struct {
	id string
	title string
	date string
	body string
	tags []string
}

// NewArticle constructs a new article instance.
func NewArticle(id string, title string, date string, body string, tags []string) Article {
	return Article{
		id,
		title,
		date,
		body,
		tags,
	}
}

func (a Article) GetID() string {
	return a.id
}

// ArticleRepository implements the repository model for Article
type ArticleRepository interface {
	Add(Article) error
	Get(Article) (Article, error)
	Find(date string, tagName string) ([]Article, error)
}