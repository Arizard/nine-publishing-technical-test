package infrastructure

import (
	"fmt"
	"github.com/arizard/nine-publishing-technical-test/entities"
)

// InMemoryArticleRepository implements a concrete repository model which
// stores data in memory.
type InMemoryArticleRepository struct {
	articles map[string]entities.Article
}

// NewInMemoryArticleRepository initializes a new concrete ArticleRepository
// which stores articles in memory.
func NewInMemoryArticleRepository() InMemoryArticleRepository {
	return InMemoryArticleRepository{
		articles: map[string]entities.Article{},
	}
}

// Add inserts a new Article into the repository.
func (s InMemoryArticleRepository) Add(article entities.Article) error {
	id := article.GetID()

	if _, ok := s.articles[id]; ok == true {
		err := fmt.Errorf("this Article ID already exists: %s", id)
		return err
	}

	s.articles[id] = article

	return nil
}

// Get retrieves an existing article from the repository.
// If the article does not exist, it returns an error.
func (s InMemoryArticleRepository) Get(id string) (entities.Article, error) {
	for _, val := range s.articles {
		if (val.GetID() == id) == true {
			return val, nil
		}
	}

	return entities.Article{}, fmt.Errorf("article not found with id %s", id)
}

// Find matches the arguments against the elements in the repository.
// It always returns a slice of entities.Article
func (s InMemoryArticleRepository) Find(tagName string, date string, limit int) ([]entities.Article, error) {
	results := make([]entities.Article, 0)

	for _, val := range s.articles {
		if len(results) >= limit && limit > 0 {
			break
		}
		if val.HasTag(tagName) {
			if val.GetDate() == date {
				results = append(results, val)
			}
		}
	}

	return results, nil
}
