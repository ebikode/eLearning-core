package course

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload ...
type Payload struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	Image               string `json:"image"`
	DurationPerQuestion int    `json:"duration_per_question"`
	Mode                string `json:"mode"`
	Status              string `json:"status"`
}

// ValidationFields struct to return for validation
type ValidationFields struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	Image               string `json:"image"`
	DurationPerQuestion string `json:"duration_per_question"`
	Mode                string `json:"mode"`
	Status              string `json:"status"`
}

// CourseRepository  provides access to the md.Course storage.
type CourseRepository interface {
	// Get returns the course with given ID.
	Get(uint) *md.Course
	GetSingleByUser(string, uint) *md.Course
	GetByUser(string, int, int) []*md.Course
	// Get returns all salaries.
	GetAll(int, int, string) []*md.Course
	// Store a given course to the repository.
	Store(md.Course) (*md.Course, error)
	// Update a given course in the repository.
	Update(*md.Course) (*md.Course, error)
	// Delete a given course in the repository.
	Delete(md.Course, bool) (bool, error)
}
