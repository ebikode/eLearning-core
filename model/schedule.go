package model

// Schedule - a struct to rep plan database model
type Schedule struct {
	BaseIntModel
	CourseID       uint    `json:"course_id" gorm:"not null;type:int(15)"`
	WeekDay        string  `json:"week_day" gorm:"type:enum('1', '2', '3', '4', '5', '6', '7');default:'1'"`
	TimeFromHour   int     `json:"time_from_hour" gorm:"type:int(3)"`
	TimeFromMinute int     `json:"time_from_minute" gorm:"type:int(3)"`
	TimeToHour     int     `json:"time_to_hour" gorm:"type:int(3)"`
	TimeToMunite   int     `json:"time_to_minute" gorm:"type:int(3)"`
	Status         string  `json:"status" gorm:"type:enum('disabled','active');default:'active'"`
	Course         *Course `json:"course"`
}
