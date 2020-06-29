package model

// Grade - a struct to rep plan database model
type Grade struct {
	BaseIntModel
	ApplicationID    uint         `json:"application_id" gorm:"not null;type:int(15)"`
	Scores           int          `json:"scores" gorm:"type:int(10)"`
	TotalScores      int          `json:"total_scores" gorm:"type:int(10)"`
	PercentageScores float64      `json:"percentage_scores" gorm:"type:float(5,2)"`
	Grade            string       `json:"grade" gorm:"type:enum('A', 'B', 'C', 'D', 'E', 'F', 'null');default:'null'"`
	Application      *Application `json:"application"`
}

// GradeReport a struct to rep application monthly reports
type GradeReport struct {
	Year      string `json:"year"`
	Month     string `json:"month"`
	MonthName string `json:"month_name"`
	ACount    int    `json:"a_count"`
	BCount    int    `json:"b_count"`
	CCount    int    `json:"c_count"`
	DCount    int    `json:"d_count"`
	ECount    int    `json:"e_count"`
	FCount    int    `json:"f_count"`
}
