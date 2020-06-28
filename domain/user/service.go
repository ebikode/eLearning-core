package user

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// UserService provides user operations
type UserService interface {
	GetUserDashboardData(string) *md.UserDashbordData
	GetTutorDashbordData(string) *md.TutorDashbordData
	GetUser(string) *md.User
	GetPubUser(string) *md.PubUser
	GetAllActivePubUser() []*md.PubUser
	GetUserByEmail(string) *md.User
	GetAllUsers(int, int) []*md.PubUser
	AuthenticateUser(string, string) (*md.PubUser, tr.TParam, error)
	CreateUser(md.User) (md.User, tr.TParam, error)
	UpdateUser(*md.User) (*md.User, tr.TParam, error)
}

type service struct {
	userRepo UserRepository
}

// NewService creates a user service with the necessary dependencies
func NewService(
	userRepo UserRepository,
) UserService {
	return &service{userRepo}
}

// Authenticate a user
func (s *service) AuthenticateUser(email string, password string) (*md.PubUser, tr.TParam, error) {
	tParam := tr.TParam{
		Key:          "error.login_error",
		TemplateData: nil,
		PluralCount:  nil,
	}

	user, err := s.userRepo.Authenticate(email)
	if err != nil {
		return nil, tParam, errors.New("Error")
	}

	isPasswordValid := ut.ValidatePassword(user.Password, password)
	if !isPasswordValid {
		return nil, tParam, errors.New("Error")
	}

	emp := s.GetPubUser(user.ID)

	return emp, tParam, nil

}

func (s *service) GetUserDashboardData(id string) *md.UserDashbordData {
	return s.userRepo.GetDashbordData(id)
}

func (s *service) GetTutorDashbordData(id string) *md.TutorDashbordData {
	return s.userRepo.GetTutorDashbordData(id)
}

// Get a user
func (s *service) GetUser(id string) *md.User {
	return s.userRepo.Get(id)
}

// Get a public user
func (s *service) GetPubUser(id string) *md.PubUser {
	return s.userRepo.GetPubUser(id)
}

// Get a public user
func (s *service) GetAllActivePubUser() []*md.PubUser {
	return s.userRepo.GetActivePubUsers()
}

// GetUserByEmail Get a user using their phone number
func (s *service) GetUserByEmail(email string) *md.User {
	return s.userRepo.GetUserByEmail(email)
}

// Get md.User
func (s *service) GetAllUsers(page, limit int) []*md.PubUser {
	return s.userRepo.GetUsers(page, limit)
}

// create new user
func (s *service) CreateUser(u md.User) (md.User, tr.TParam, error) {

	token := ut.RandomBase64String(30, "")
	tokenCopy := token

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	hashedToken, _ := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)

	u.Password = string(hashedPassword)
	u.EmailToken = string(hashedToken)
	u.PincodeSentAt = time.Now().UTC()

	// Generate user ID
	uID := ut.RandomBase64String(8, "elus")

	u.ID = uID

	user, err := s.userRepo.Store(u)

	fmt.Printf("Error:: %s \n", err)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return user, tParam, err
	}

	user.EmailToken = tokenCopy

	return user, tr.TParam{}, nil
}

// update existing user
func (s *service) UpdateUser(u *md.User) (*md.User, tr.TParam, error) {

	user, err := s.userRepo.Update(u)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return user, tParam, err
	}

	return user, tr.TParam{}, nil

}

// Validate - Function for validating user input during creation
func Validate(user md.User, r *http.Request) error {
	return validation.ValidateStruct(&user,
		// Phone cannot be empty, and th length must between 7 and 20
		validation.Field(&user.Phone, ut.PhoneRule(r)...),
		validation.Field(&user.Email, ut.EmailRule(r)...),
		validation.Field(&user.FirstName, ut.NameRule(r)...),
		validation.Field(&user.LastName, ut.NameRule(r)...),
		validation.Field(&user.Username, ut.NameRule(r)...),
		validation.Field(&user.Password, ut.PasswordRule(r)...),
	)
}

// ValidateUpdates - Function for validating user input during update
func ValidateUpdates(user md.User, r *http.Request) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, ut.EmailRule(r)...),
		validation.Field(&user.FirstName, ut.NameRule(r)...),
		validation.Field(&user.LastName, ut.NameRule(r)...),
		validation.Field(&user.Username, ut.NameRule(r)...),
		validation.Field(&user.Avatar, ut.AvatarRule(r)...),
		validation.Field(&user.Thumb, ut.AvatarRule(r)...),
	)
}
