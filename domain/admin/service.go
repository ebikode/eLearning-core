package admin

import (
	"errors"
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// Service provides admin operations
type AdminService interface {
	// CheckAdminCreated - Checks if a default admin has already been created
	// used when the server is ran for the very first time so as to create
	// a default admin if it returns false
	CheckAdminCreated() bool
	GetAdmin(string) *md.Admin
	GetAdminDashboardData() *md.DashbordData
	AuthenticateAdmin(string, string) (*md.Admin, tr.TParam, error)
	CreateAdmin(md.Admin) (md.Admin, tr.TParam, error)
	UpdateAdmin(*md.Admin) (*md.Admin, tr.TParam, error)
}

type service struct {
	aRepo AdminRepository
}

// NewService creates a admin service with the necessary dependencies
func NewService(
	aRepo AdminRepository,
) AdminService {
	return &service{aRepo}
}

// CheckAdminCreated - Checks if a default admin has already been created
// used when the server is ran for the very first time so as to create
// a default admin if it returns false
func (s *service) CheckAdminCreated() bool {
	return s.aRepo.CheckAdminCreated()
}

// Authenticate a admin
func (s *service) AuthenticateAdmin(staffID string, password string) (*md.Admin, tr.TParam, error) {
	tParam := tr.TParam{
		Key:          "error.login_error",
		TemplateData: nil,
		PluralCount:  nil,
	}

	if len(staffID) < 5 {

		return nil, tParam, errors.New("Error")
	}

	admin, err := s.aRepo.Authenticate(staffID)

	if err != nil || admin.Status != "active" {
		return nil, tParam, errors.New("Error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, tParam, errors.New("Error")
	}

	if admin.Status == "offline" {
		admin.Status = "online"

		s.UpdateAdmin(admin)
	}
	admin.Password = ""

	return admin, tParam, nil

}

// GetAdminDashboardData ...
func (s *service) GetAdminDashboardData() *md.DashbordData {
	return s.aRepo.GetDashbordData()
}

// Get a admin
func (s *service) GetAdmin(id string) *md.Admin {
	return s.aRepo.Get(id)
}

// create new admin
func (s *service) CreateAdmin(u md.Admin) (md.Admin, tr.TParam, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.Password = string(hashedPassword)

	// Generate admin ID
	uID := ut.RandomBase64String(8, "MDs")

	u.ID = uID
	u.Status = "active"

	admin, err := s.aRepo.Store(u)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return admin, tParam, err
	}

	return admin, tr.TParam{}, nil
}

// update existing admin
func (s *service) UpdateAdmin(u *md.Admin) (*md.Admin, tr.TParam, error) {
	admin, err := s.aRepo.Update(u)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return admin, tParam, err
	}

	return admin, tr.TParam{}, nil

}

// Validate - Function for validating customer input
func Validate(admin md.Admin, r *http.Request) error {
	return validation.ValidateStruct(&admin,
		// Phone cannot be empty, and th length must between 7 and 20
		validation.Field(&admin.Phone, ut.PhoneRule(r)...),
		validation.Field(&admin.Email, ut.EmailRule(r)...),
		validation.Field(&admin.FirstName, ut.NameRule(r)...),
		validation.Field(&admin.LastName, ut.NameRule(r)...),
		validation.Field(&admin.Password, ut.PasswordRule(r)...),
	)
}

// ValidateUpdate Function for validating customer input
func ValidateUpdate(admin md.Admin, r *http.Request) error {
	return validation.ValidateStruct(&admin,
		// Phone cannot be empty, and th length must between 7 and 20
		validation.Field(&admin.ID, ut.IDRule(r)...),
		validation.Field(&admin.Phone, ut.PhoneRule(r)...),
		validation.Field(&admin.Email, ut.EmailRule(r)...),
		validation.Field(&admin.FirstName, ut.NameRule(r)...),
		validation.Field(&admin.LastName, ut.NameRule(r)...),
		validation.Field(&admin.Password, ut.PasswordRule(r)...),
	)
}
