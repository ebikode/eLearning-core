package user

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload struct to return for validation
type Payload struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Thumb     string `json:"thumb,omitempty"`
	Role      string `json:"role,omitempty"`
	Status    string `json:"status,omitempty"`
}

// ValidationFields struct to return for validation
type ValidationFields struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Password  string `json:"password,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Thumb     string `json:"thumb,omitempty"`
}

//UserRepository Repository provides access to the user storage.
type UserRepository interface {
	GetDashbordData(string) *md.UserDashbordData
	// Get returns the user with given ID.
	Get(string) *md.User
	// returns the public user with given ID.
	GetPubUser(string) *md.PubUser
	GetActivePubUsers() []*md.PubUser
	// GetUserByPhone returns the user with given phone number.
	GetUserByEmail(string) *md.User
	// GetUsers returns all users paginated.
	GetUsers(int, int) []*md.PubUser
	// Authenticate a user
	Authenticate(string) (*md.User, error)
	// Saves a given user to the repository.
	Store(md.User) (md.User, error)
	// Update a given user in the repository.
	Update(*md.User) (*md.User, error)
	// Delete a given user in the repository.
	Delete(md.User, bool) (bool, error)
}
