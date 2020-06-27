package admin

import (
	md "github.com/ebikode/eLearning-core/model"
)

// ValidationFields struct to return for validation
type ValidationFields struct {
	Phone     string `json:"phone,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

// AdminRepository provides access to the admin storage.
type AdminRepository interface {
	// Check if a default admin has been created.
	CheckAdminCreated() bool
	// Get returns the admin with given ID.
	Get(string) *md.Admin
	// GetDashbordData  returns the admin Dashboar data.
	GetDashbordData() *md.DashbordData
	// Authenticate a admin
	Authenticate(string) (*md.Admin, error)
	// Saves a given admin to the repository.
	Store(md.Admin) (md.Admin, error)
	// Update a given admin in the repository.
	Update(*md.Admin) (*md.Admin, error)
	// Delete a given admin in the repository.
	Delete(md.Admin, bool) (bool, error)
}
