package assessment

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload Request data
type Payload struct {
	ApplicationID  uint   `json:"application_id,omitempty"`
	QuestionID     string `json:"question_id,omitempty"`
	SelectedAnswer string `json:"selected_answer,omitempty"`
}

// ValidationFields  Error response data
type ValidationFields struct {
	ApplicationID  string `json:"application_id,omitempty"`
	QuestionID     string `json:"question_id,omitempty"`
	SelectedAnswer string `json:"selected_answer,omitempty"`
}

// AssessmentRepository provides access to the Assessment storage.
type AssessmentRepository interface {
	Get(string, string) *md.Assessment
	GetLastAssessment() *md.Assessment
	// returns all assessments set with page and limit.
	GetAll(int, int) []*md.Assessment
	// returns the assessments with given userID.
	GetByUser(string, int, int) []*md.Assessment
	GetByCourse(int) []*md.Assessment
	GetSingleByCourse(int) *md.Assessment
	// Store a given user assessment to the repository.
	Store(md.Assessment) (*md.Assessment, error)
	// Update a given assessment in the repository.
	Update(*md.Assessment) (*md.Assessment, error)
	// Delete a given assessment in the repository.
	Delete(*md.Assessment, bool) (bool, error)
}
