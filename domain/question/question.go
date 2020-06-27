package question

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload ...
type Payload struct {
	CourseID uint   `json:"course_id"`
	Question string `json:"question"`
	OptionA  string `json:"option_a"`
	OptionB  string `json:"option_b"`
	OptionC  string `json:"option_c"`
	Answer   string `json:"answer"`
	Solution string `json:"solution"`
	Status   string `json:"status"`
}

// ValidationFields struct to return for validation
type ValidationFields struct {
	CourseID string `json:"course_id,omitempty"`
	Question string `json:"question,omitempty"`
	OptionA  string `json:"option_a"`
	OptionB  string `json:"option_b"`
	OptionC  string `json:"option_c"`
	Answer   string `json:"answer"`
	Solution string `json:"solution"`
	Status   string `json:"status"`
}

// QuestionRepository  provides access to the md.Question storage.
type QuestionRepository interface {
	// Get returns the question with given ID.
	Get(string) *md.Question
	GetByCourse(uint) []*md.PubQuestion
	// Get returns all questions.
	GetAll(int, int) []*md.Question
	// Store a given question to the repository.
	Store(md.Question) (*md.Question, error)
	// Update a given question in the repository.
	Update(*md.Question) (*md.Question, error)
	// Delete a given question in the repository.
	Delete(md.Question, bool) (bool, error)
}
