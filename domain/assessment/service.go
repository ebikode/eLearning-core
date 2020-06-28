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
	GetAssessment(uint, string) *md.Assessment
	GetLastAssessment() *md.Assessment
	GetAssessments(int, int) []*md.Assessment
	GetAssessmentsByCourse(int) []*md.Assessment
	GetSingleAssessmentByCourse(int) *md.Assessment
	GetUserApplicationAssessments(string, uint, int, int) []*md.Assessment
	CreateAssessment(md.Assessment) (*md.Assessment, tr.TParam, error)
	UpdateAssessment(*md.Assessment) (*md.Assessment, tr.TParam, error)
}

type service struct {
	aRepo AssessmentRepository
}

// NewService creates assessment service with the necessary dependencies
func NewService(
	aRepo AssessmentRepository,
) AssessmentService {
	return &service{aRepo}
}

/*
* Get single assessment
* @param applicationID => the ID of the application whose assessment is needed
* @param assessmentID => the ID of the assessment requested.
 */
func (s *service) GetAssessment(applicationID uint, assessmentID string) *md.Assessment {
	return s.aRepo.Get(applicationID, assessmentID)
}

/*
* GetLastAssessment
 */
func (s *service) GetLastAssessment() *md.Assessment {
	return s.aRepo.GetLastAssessment()
}

/*
* Get all assessments
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetAssessments(page, limit int) []*md.Assessment {
	return s.aRepo.GetAll(page, limit)
}

func (s *service) GetAssessmentsByCourse(courseID int) []*md.Assessment {
	return s.aRepo.GetByCourse(courseID)
}

func (s *service) GetSingleAssessmentByCourse(courseID int) *md.Assessment {
	return s.aRepo.GetSingleByCourse(courseID)
}

/*
* Get all assessments of a user
* @param applicationID => the ID of the application whose assessment is needed
* @param userID => the ID of the user whose application assessment is needed
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetUserApplicationAssessments(userID string, applicationID uint, page, limit int) []*md.Assessment {
	return s.aRepo.GetByApplication(userID, applicationID, page, limit)
}

// Create New assessment
func (s *service) CreateAssessment(ase md.Assessment) (*md.Assessment, tr.TParam, error) {
	// Generate ID
	aseID := ut.RandomBase64String(8, "elasm")
	ase.ID = aseID

	assessment, err := s.aRepo.Store(ase)

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
	assessment, err := s.aRepo.Update(p)

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
