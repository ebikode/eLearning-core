package appsetting

import (
	md "github.com/ebikode/eLearning-core/model"
)

// ValidationFields struct to return for validation
type ValidationFields struct {
	Name string `json:"name,omitempty"`
	// SKey   string `json:"s_key,omitempty"`
	Value   string `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`
	Status  string `json:"status,omitempty"`
}

// AppSettingRepository Repository provides access to the md.AppSetting storage.
type AppSettingRepository interface {
	// Get returns the appsetting with given ID.
	Get(uint) *md.AppSetting
	// Get app settings by key
	GetByKey(string, string) *md.AppSetting
	// Get returns all appsettings.
	GetAll(string) []*md.AppSetting
	// Store a given customer appsetting to the repository.
	Store(md.AppSetting) (*md.AppSetting, error)
	// Update a given appsetting in the repository.
	Update(*md.AppSetting) (*md.AppSetting, error)
	// Delete a given appsetting in the repository.
	Delete(md.AppSetting, bool) (bool, error)
}
