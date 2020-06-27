package model

// Question - a struct to rep question database model
type Question struct {
	BaseModel
	CourseID uint    `json:"course_id" gorm:"not null;type:int(15)"`
	Question string  `json:"question" gorm:"type:varchar(200)"`
	OptionA  string  `json:"option_a" gorm:"type:varchar(200)"`
	OptionB  string  `json:"option_b" gorm:"type:varchar(200)"`
	OptionC  string  `json:"option_c" gorm:"type:varchar(200)"`
	Answer   string  `json:"answer" gorm:"type:varchar(200)"`
	Solution string  `json:"solution" gorm:"type:varchar(500)"`
	Status   string  `json:"status" gorm:"type:enum('disabled','active');default:'active'"`
	Course   *Course `json:"course"`
}

// PubQuestion - a struct to rep question database model shown without the answer
type PubQuestion struct {
	BaseModel
	CourseID uint    `json:"course_id" gorm:"not null;type:int(15)"`
	Question string  `json:"question" gorm:"type:varchar(200)"`
	OptionA  string  `json:"option_a" gorm:"type:varchar(200)"`
	OptionB  string  `json:"option_b" gorm:"type:varchar(200)"`
	OptionC  string  `json:"option_c" gorm:"type:varchar(200)"`
	Status   string  `json:"status" gorm:"type:enum('disabled','active');default:'active'"`
	Course   *Course `json:"course"`
}

// TableName Set PubQuestion's table name to be `Users`
func (PubQuestion) TableName() string {
	return "questions"
}
