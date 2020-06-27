package model

// Assessment a struct encapsulating payment
type Assessment struct {
	BaseModel
	ApplicationID  uint      `json:"application_id" gorm:"not null;type:int(15)"`
	QuestionID     string    `json:"question_id" gorm:"not null;type:varchar(20)"`
	SelectedAnswer string    `json:"selected_answer" gorm:"type:varchar(100)"`
	CorrectAnswer  string    `json:"correct_answer" gorm:"type:varchar(100)"`
	IsCorrect      bool      `json:"is_correct" gorm:"type:tinyint(1);default:0"`
	Question       *Question `json:"question"`
}
