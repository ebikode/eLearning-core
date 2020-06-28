package model

// Schedule - a struct to rep plan database model
type Schedule struct {
	BaseIntModel
	CourseID       uint    `json:"course_id" gorm:"not null;type:int(15)"`
	WeekDay        string  `json:"week_day" gorm:"type:enum('1', '2', '3', '4', '5', '6', '7');default:'1'"`
	TimeFromHour   string  `json:"time_from_hour" gorm:"type:varchar(3)"`
	TimeFromMinute string  `json:"time_from_minute" gorm:"type:varchar(3)"`
	TimeToHour     string  `json:"time_to_hour" gorm:"type:varchar(3)"`
	TimeToMunite   string  `json:"time_to_minute" gorm:"type:varchar(3)"`
	Status         string  `json:"status" gorm:"type:enum('disabled','active');default:'active'"`
	Course         *Course `json:"course"`
}
