package entities

// Article is an entity which models the Article
type Article struct {
	ID    string   `json:"id,omitempty"`
	Title string   `json:"title,omitempty"`
	Date  string   `json:"date,omitempty"`
	Body  string   `json:"body,omitempty"`
	Tags  []string `json:"tags,omitempty"`
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

// GetID returns the article ID.
func (a Article) GetID() string {
	return a.ID
}

// GetDate returns the article date.
func (a Article) GetDate() string {
	return a.Date
}

// GetTags returns the article tags.
func (a Article) GetTags() []string {
	return a.Tags
}

// HasTag returns true when the Tags property contains the tagName argument.
func (a Article) HasTag(tagName string) bool {
	for _, t := range a.Tags {
		if t == tagName {
			return true
		}
	}
	return false
}

// ArticleRepository implements the repository model for Article
type ArticleRepository interface {
	Add(Article) error
	Get(string) (Article, error)
	Find(date string, tagName string, limit int) ([]Article, error)
}
