package question

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// QuestionService  provides question operations
type QuestionService interface {
	GetQuestion(string) *md.Question
	GetQuestionsByCourse(uint) []*md.PubQuestion
	GetQuestions(int, int) []*md.Question
	CreateQuestion(md.Question) (*md.Question, tr.TParam, error)
	UpdateQuestion(*md.Question) (*md.Question, tr.TParam, error)
}

type service struct {
	qRepo QuestionRepository
}

// NewService creates a question service with the necessary dependencies
func NewService(
	qRepo QuestionRepository,
) QuestionService {
	return &service{qRepo}
}

// Get a question
func (s *service) GetQuestion(id string) *md.Question {
	return s.qRepo.Get(id)
}

// GetQuestions Get all questions from DB
//
// @userType == admin | customer
func (s *service) GetQuestions(page, limit int) []*md.Question {
	return s.qRepo.GetAll(page, limit)
}

func (s *service) GetQuestionsByCourse(courseID uint) []*md.PubQuestion {
	return s.qRepo.GetByCourse(courseID)
}

// CreateQuestion Creates New question
func (s *service) CreateQuestion(c md.Question) (*md.Question, tr.TParam, error) {

	question, err := s.qRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return question, tParam, err
	}

	return question, tr.TParam{}, nil

}

// UpdateQuestion update existing question
func (s *service) UpdateQuestion(c *md.Question) (*md.Question, tr.TParam, error) {
	question, err := s.qRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return question, tParam, err
	}

	return question, tr.TParam{}, nil

}

// Validate Function for validating question input
func Validate(question Payload, r *http.Request) error {
	return validation.ValidateStruct(&question,
		validation.Field(&question.CourseID, ut.IDRule(r)...),
		validation.Field(&question.Question, ut.RequiredRule(r, "general.question")...),
		validation.Field(&question.OptionA, ut.RequiredRule(r, "general.pension")...),
		validation.Field(&question.OptionB, ut.RequiredRule(r, "general.optionB")...),
		validation.Field(&question.OptionC, ut.RequiredRule(r, "general.option_c")...),
		validation.Field(&question.Answer, ut.RequiredRule(r, "general.answer")...),
	)
}

// ValidateUpdates Function for validating question update input
func ValidateUpdates(question Payload, r *http.Request) error {
	return validation.ValidateStruct(&question,
		validation.Field(&question.Question, ut.MoneyRule(r)...),
		validation.Field(&question.Question, ut.RequiredRule(r, "general.question")...),
		validation.Field(&question.OptionA, ut.RequiredRule(r, "general.pension")...),
		validation.Field(&question.OptionB, ut.RequiredRule(r, "general.optionB")...),
		validation.Field(&question.OptionC, ut.RequiredRule(r, "general.option_c")...),
		validation.Field(&question.Answer, ut.RequiredRule(r, "general.answer")...),
	)
}
