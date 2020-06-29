package authdDevice

import (
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
)

// Service provides authdDevice operations
type AuthdDeviceService interface {
	GetAuthdDevice(string) *md.AuthdDevice
	CreateAuthdDevice(md.AuthdDevice) (md.AuthdDevice, tr.TParam, error)
	UpdateAuthdDevice(*md.AuthdDevice) (*md.AuthdDevice, tr.TParam, error)
}

type service struct {
	audRepo AuthdDeviceRepository
}

// NewService creates a authdDevice service with the necessary dependencies
func NewService(
	audRepo AuthdDeviceRepository,
) AuthdDeviceService {
	return &service{audRepo}
}

// Get a authdDevice
func (s *service) GetAuthdDevice(id string) *md.AuthdDevice {
	return s.audRepo.Get(id)
}

// Create New authdDevice
func (s *service) CreateAuthdDevice(c md.AuthdDevice) (md.AuthdDevice, tr.TParam, error) {
	authdDevice, err := s.audRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return authdDevice, tParam, err
	}

	return authdDevice, tr.TParam{}, nil

}

// update existing authdDevice
func (s *service) UpdateAuthdDevice(c *md.AuthdDevice) (*md.AuthdDevice, tr.TParam, error) {
	authdDevice, err := s.audRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return authdDevice, tParam, err
	}

	return authdDevice, tr.TParam{}, nil

}
