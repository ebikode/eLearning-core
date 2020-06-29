package journal

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

// JournalRepository  provides access to the md.Journal storage.
type JournalRepository interface {
	// Get returns the journal with given ID.
	Get(uint) *md.Journal
	GetByCourse(int) []*md.Journal
	GetByUser(string, int, int) []*md.Journal
	// Get returns all journals.
	GetAll(int, int) []*md.Journal
	// Store a given journal to the repository.
	Store(md.Journal) (*md.Journal, error)
	// Update a given journal in the repository.
	Update(*md.Journal) (*md.Journal, error)
	// Delete a given journal in the repository.
	Delete(md.Journal, bool) (bool, error)
}
