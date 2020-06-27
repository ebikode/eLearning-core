package application

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload ...
type Payload struct {
	CourseID uint   `json:"course_id"`
	Status   string `json:"status"`
}

// ValidationFields struct to return for validation
type ValidationFields struct {
	CourseID string `json:"course_id"`
	Status   string `json:"status"`
}

// ApplicationRepository  provides access to the md.Application storage.
type ApplicationRepository interface {
	// Get returns the application with given ID.
	Get(uint) *md.Application
	GetByUser(string) []*md.Application
	GetByCourse(int) []*md.Application
	// Get returns all applications.
	GetAll(int, int) []*md.Application
	// Store a given application to the repository.
	Store(md.Application) (*md.Application, error)
	// Update a given application in the repository.
	Update(*md.Application) (*md.Application, error)
	// Delete a given application in the repository.
	Delete(md.Application, bool) (bool, error)
}
