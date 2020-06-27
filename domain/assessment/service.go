package assessment

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// AssessmentService provides assessment operations
type AssessmentService interface {
	GetAssessment(string, string) *md.Assessment
	GetLastAssessment() *md.Assessment
	GetAssessments(int, int) []*md.Assessment
	GetAssessmentsByCourse(int) []*md.Assessment
	GetSingleAssessmentByCourse(int) *md.Assessment
	GetUserApplicationAssessments(string, string, int, int) []*md.Assessment
	CreateAssessment(md.Assessment) (*md.Assessment, tr.TParam, error)
	UpdateAssessment(*md.Assessment) (*md.Assessment, tr.TParam, error)
}

type service struct {
	pRepo AssessmentRepository
}

// NewService creates assessment service with the necessary dependencies
func NewService(
	pRepo AssessmentRepository,
) AssessmentService {
	return &service{pRepo}
}

/*
* Get single assessment of a user
* @param userID => the ID of the user whose assessment is needed
* @param assessmentID => the ID of the assessment requested.
 */
func (s *service) GetAssessment(userID, assessmentID string) *md.Assessment {
	return s.pRepo.Get(userID, assessmentID)
}

/*
* GetLastAssessment
 */
func (s *service) GetLastAssessment() *md.Assessment {
	return s.pRepo.GetLastAssessment()
}

/*
* Get all assessments
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetAssessments(page, limit int) []*md.Assessment {
	return s.pRepo.GetAll(page, limit)
}

func (s *service) GetAssessmentsByCourse(courseID int) []*md.Assessment {
	return s.pRepo.GetByCourse(courseID)
}

func (s *service) GetSingleAssessmentByCourse(courseID int) *md.Assessment {
	return s.pRepo.GetSingleByCourse(courseID)
}

/*
* Get all assessments of a user
* @param applicationID => the ID of the application whose assessment is needed
* @param userID => the ID of the user whose application assessment is needed
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetUserApplicationAssessments(userID, applicationID string, page, limit int) []*md.Assessment {
	return s.pRepo.GetByApplication(userID, applicationID, page, limit)
}

// Create New assessment
func (s *service) CreateAssessment(p md.Assessment) (*md.Assessment, tr.TParam, error) {
	// Generate ID
	pID := ut.RandomBase64String(8, "EL")
	p.ID = pID

	assessment, err := s.pRepo.Store(p)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return assessment, tParam, err
	}

	return assessment, tr.TParam{}, nil
}

// update existing assessment
func (s *service) UpdateAssessment(p *md.Assessment) (*md.Assessment, tr.TParam, error) {
	assessment, err := s.pRepo.Update(p)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return assessment, tParam, err
	}
	return assessment, tr.TParam{}, nil
}

// Validate Function for validating assessment input
func Validate(assessment Payload, r *http.Request) error {
	return validation.ValidateStruct(&assessment,
		validation.Field(&assessment.ApplicationID, ut.IDRule(r)...),
		validation.Field(&assessment.QuestionID, ut.IDRule(r)...),
		validation.Field(&assessment.SelectedAnswer, ut.RequiredRule(r, "general.answer")...),
	)
}

// ValidateUpdates Function for validating assessment update input
func ValidateUpdates(assessment Payload, r *http.Request) error {
	return validation.ValidateStruct(&assessment)
}
