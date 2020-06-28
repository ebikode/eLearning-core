package application

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ApplicationService  provides application operations
type ApplicationService interface {
	GetApplication(uint) *md.Application
	GetUserApplications(string) []*md.Application
	GetApplicationsByCourse(int) []*md.Application
	GetApplications(int, int) []*md.Application
	CreateApplication(md.Application) (*md.Application, tr.TParam, error)
	UpdateApplication(*md.Application) (*md.Application, tr.TParam, error)
}

type service struct {
	aRepo ApplicationRepository
}

// NewService creates a application service with the necessary dependencies
func NewService(
	aRepo ApplicationRepository,
) ApplicationService {
	return &service{aRepo}
}

// Get a application
func (s *service) GetApplication(id uint) *md.Application {
	return s.aRepo.Get(id)
}

// GetApplications Get all applications from DB
//
// @userType == admin | customer
func (s *service) GetApplications(page, limit int) []*md.Application {
	return s.aRepo.GetAll(page, limit)
}

func (s *service) GetUserApplications(userID string) []*md.Application {
	return s.aRepo.GetByUser(userID)
}

func (s *service) GetApplicationsByCourse(courseID int) []*md.Application {
	return s.aRepo.GetByCourse(courseID)
}

// CreateApplication Creates New application
func (s *service) CreateApplication(c md.Application) (*md.Application, tr.TParam, error) {

	// Generate ref
	ref := ut.RandomBase64String(8, "elref")
	c.ReferenceNo = ref

	application, err := s.aRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return application, tParam, err
	}

	return application, tr.TParam{}, nil

}

// UpdateApplication update existing application
func (s *service) UpdateApplication(c *md.Application) (*md.Application, tr.TParam, error) {
	application, err := s.aRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return application, tParam, err
	}

	return application, tr.TParam{}, nil

}

// Validate Function for validating application input
func Validate(application Payload, r *http.Request) error {
	return validation.ValidateStruct(&application,
		validation.Field(&application.CourseID, ut.IDRule(r)...),
	)
}

// ValidateUpdates Function for validating application update input
func ValidateUpdates(application Payload, r *http.Request) error {
	return validation.ValidateStruct(&application)
}
