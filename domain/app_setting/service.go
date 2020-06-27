package appsetting

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Service provides appsetting operations
type AppSettingService interface {
	GetAppSetting(uint) *md.AppSetting
	GetAppSettingByKey(string, string) *md.AppSetting
	GetAppSettings(string) []*md.AppSetting
	CreateAppSetting(md.AppSetting) (*md.AppSetting, tr.TParam, error)
	UpdateAppSetting(*md.AppSetting) (*md.AppSetting, tr.TParam, error)
}

type service struct {
	asRepo AppSettingRepository
}

// NewService creates a appsetting service with the necessary dependencies
func NewService(
	asRepo AppSettingRepository,
) AppSettingService {
	return &service{asRepo}
}

// Get a appsetting
// userType = admin or customer
func (s *service) GetAppSetting(id uint) *md.AppSetting {
	return s.asRepo.Get(id)
}

// et appsetting by key
// userType = admin or customer
func (s *service) GetAppSettingByKey(key, userType string) *md.AppSetting {
	return s.asRepo.GetByKey(key, userType)
}

// Get a appsetting
// userType = admin or customer
func (s *service) GetAppSettings(userType string) []*md.AppSetting {
	return s.asRepo.GetAll(userType)
}

// Create New appsetting
func (s *service) CreateAppSetting(c md.AppSetting) (*md.AppSetting, tr.TParam, error) {

	c.Status = "active"
	appsetting, err := s.asRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return appsetting, tParam, err
	}

	return appsetting, tr.TParam{}, nil

}

// update existing appsetting
func (s *service) UpdateAppSetting(c *md.AppSetting) (*md.AppSetting, tr.TParam, error) {
	appsetting, err := s.asRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return appsetting, tParam, err
	}

	return appsetting, tr.TParam{}, nil

}

// Function for validating st input
func Validate(st md.AppSetting, r *http.Request) error {
	return validation.ValidateStruct(&st,
		validation.Field(&st.Name, ut.NameRule(r)...),
		validation.Field(&st.SKey, ut.KeyRule(r)...),
		validation.Field(&st.Comment, ut.RequiredCommentRule(r)...),
		validation.Field(&st.Value, ut.RequiredRule(r, "value")...),
	)
}

// Function for validating st input
func ValidateUpdates(st md.AppSetting, r *http.Request) error {
	return validation.ValidateStruct(&st,
		validation.Field(&st.Name, ut.NameRule(r)...),
		// validation.Field(&st.SKey, ut.KeyRule(r)...),
		validation.Field(&st.Value, ut.RequiredRule(r, "value")...),
		validation.Field(&st.Comment, ut.RequiredCommentRule(r)...),
		validation.Field(&st.Status, ut.EnumRule(r, "general.status", "active", "disabled")...),
	)
}
