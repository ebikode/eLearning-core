package model

// Journal - a struct to rep plan database model
type Journal struct {
	BaseIntModel
	UserID   string   `json:"user_id" gorm:"not null;type:varchar(20)"`
	CourseID uint     `json:"course_id" gorm:"not null;type:int(15)"`
	Title    string   `json:"title" gorm:"type:varchar(200)"`
	Body     string   `json:"body" gorm:"type:varchar(5000)"`
	Status   string   `json:"status" gorm:"type:enum('disabled','active');default:'active'"`
	Course   *Course  `json:"course"`
	User     *PubUser `json:"user"`
}
