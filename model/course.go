package model

// Course - a struct to rep plan database model
type Course struct {
	BaseIntModel
	UserID              string   `json:"user_id" gorm:"not null;type:varchar(20)"`
	Title               string   `json:"title" gorm:"type:varchar(100)"`
	ReferenceNo         string   `json:"reference_no" gorm:"type:varchar(50)"`
	Description         string   `json:"description" gorm:"type:varchar(1000)"`
	Image               string   `json:"image" gorm:"type:varchar(500)"`
	DurationPerQuestion int      `json:"duration_per_question" gorm:"type:int(10)"`
	Mode                string   `json:"mode" gorm:"type:enum('instructor_led', 'blended');default:'instructor_led'"`
	Status              string   `json:"status" gorm:"type:enum('pending', 'approved', 'suspended', 'disabled');default:'pending'"`
	User                *PubUser `json:"user"`
}

// CourseReport a struct to rep course monthly reports
type CourseReport struct {
	Year      string `json:"year"`
	Month     string `json:"month"`
	MonthName string `json:"month_name"`
	Count     int    `json:"count"`
}
