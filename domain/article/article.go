package article

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload ...
type Payload struct {
	CourseID uint   `json:"course_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Status   string `json:"status"`
}

// ValidationFields struct to return for validation
type ValidationFields struct {
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Status   string `json:"status"`
}

// ArticleRepository  provides access to the md.Article storage.
type ArticleRepository interface {
	// Get returns the article with given ID.
	Get(uint) *md.Article
	GetByCourse(int) *md.Article
	// Get returns all articles.
	GetAll(int, int) []*md.Article
	// Store a given article to the repository.
	Store(md.Article) (*md.Article, error)
	// Update a given article in the repository.
	Update(*md.Article) (*md.Article, error)
	// Delete a given article in the repository.
	Delete(md.Article, bool) (bool, error)
}
