package grade

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload Request data
type Payload struct {
	Month         int    `json:"month,omitempty"`
	Year          int    `json:"year,omitempty"`
	Status        string `json:"status,omitempty"`
	PaymentStatus string `json:"payment_status,omitempty"`
}

// ValidationFields  Error response data
type ValidationFields struct {
	Month         string `json:"month,omitempty"`
	Year          string `json:"year,omitempty"`
	Status        string `json:"status,omitempty"`
	PaymentStatus string `json:"account_id,omitempty"`
}

// GradeRepository provides access to the Grade storage.
type GradeRepository interface {
	GetReports(int, int) []*md.GradeReport
	// Get returns the grade with given ID.
	Get(uint) *md.Grade
	// returns all grades set with page and limit.
	GetAll(int, int) []*md.Grade
	// returns the grades with given userID.
	GetByUser(string) []*md.Grade
	GetByApplication(int) *md.Grade
	GetByCourse(string, int, int, int) []*md.Grade
	// Store a given user grade to the repository.
	Store(md.Grade) (*md.Grade, error)
	// Update a given grade in the repository.
	Update(*md.Grade) (*md.Grade, error)
	// Delete a given grade in the repository.
	// Delete(*md.Grade, bool) (bool, error)
}
