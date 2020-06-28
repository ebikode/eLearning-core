package model

import "time"

//User - a struct to rep User account
type User struct {
	BaseModel
	Lang            string    `json:"lang" gorm:"not null;type:varchar(5);default:'en'"`
	FirstName       string    `json:"first_name" gorm:"not null;type:varchar(50);"`
	LastName        string    `json:"last_name" gorm:"not null;type:varchar(50);"`
	Username        string    `json:"username" gorm:"not null;type:varchar(20);unique_index"`
	Address         string    `json:"address" gorm:"type:varchar(150)"`
	About           string    `json:"about" gorm:"type:varchar(150)"`
	Email           string    `json:"email" gorm:"not null,type:varchar(100);unique_index"`
	EmailToken      string    `json:"email_token" gorm:"type:varchar(200)"`
	Password        string    `json:"password,omitempty" gorm:"not null;type:varchar(250)"`
	Pincode         string    `json:"pincode,omitempty" gorm:"not null;type:varchar(250)"`
	Phone           string    `json:"phone" gorm:"not null;type: varchar(20);unique_index"`
	IsPincodeUsed   bool      `json:"is_pincode_used" gorm:"type:tinyint(1);default:0"`
	PincodeSentAt   time.Time `json:"pincode_sent_at"`
	IsPhoneVerified bool      `json:"is_phone_verified" gorm:"type:tinyint(1);default:0"`
	IsEmailVerified bool      `json:"is_email_verified" gorm:"type:tinyint(1);default:0"`
	Avatar          string    `json:"avatar" gorm:"type:varchar(500)"`
	Thumb           string    `json:"thumb" gorm:"type:varchar(500)"`
	Role            string    `json:"role" gorm:"type:enum('user','tutor');default:'user'"`
	Status          string    `json:"status" gorm:"type:enum('pending','active','suspended','resigned','fired','deleted');default:'pending'"`
	// Added for request body validation only
	Token string `json:"token,omitempty" gorm:"-"`
}

//PubUser - a struct to rep User account shown to others
// e.g Admin
type PubUser struct {
	BaseModel
	Lang            string `json:"lang"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	Address         string `json:"address"`
	About           string `json:"about"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	IsPhoneVerified bool   `json:"is_phone_verified"`
	IsEmailVerified bool   `json:"is_email_verified"`
	Avatar          string `json:"avatar"`
	Thumb           string `json:"thumb"`
	Role            string `json:"role"`
	Status          string `json:"status"`
}

// TutorDashbordData - struct encapsulating tutor dashboard data
type TutorDashbordData struct {
	Certificates         int64 `json:"certificates"`
	Applications         int64 `json:"applications"`
	CompletedAssessments int64 `json:"completed_assessments"`
	Courses              int64 `json:"courses"`
}

// UserDashbordData - struct encapsulating user dashboard data
type UserDashbordData struct {
	Certificates         int64 `json:"certificate"`
	Applications         int64 `json:"applications"`
	CompletedAssessments int64 `json:"completed_assessments"`
	Courses              int64 `json:"courses"`
}

// TableName Set PubUser's table name to be `Users`
func (PubUser) TableName() string {
	return "users"
}
