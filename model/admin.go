package model

//Admin - a struct to rep admin account
type Admin struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"not null;type:varchar(20)"`
	LastName  string `json:"last_name" gorm:"not null;type:varchar(20)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password,omitempty" gorm:"not null;type:varchar(250)"`
	Phone     string `json:"phone" gorm:"not null;type: varchar(20);unique_index"`
	Avatar    string `json:"avatar" gorm:"type:varchar(500)"`
	Thumb     string `json:"thumb" gorm:"type:varchar(500)"`
	Role      string `json:"role" gorm:"type:enum('manager','super_admin');default:'manager'"`
	Status    string `json:"status" gorm:"type:enum('pending','active','suspended','deleted');default:'pending'"`
	Token     string `json:"token,omitempty" gorm:"-"`
}

// DashbordData - struct encapsulating admin dashboard data
type DashbordData struct {
	UsersCount           int64 `json:"users_count"`
	ActiveUsersCount     int64 `json:"active_users_count"`
	PendingUsersCount    int64 `json:"pending_users_count"`
	TutorsCount          int64 `json:"tutors_count"`
	ActiveTutorsCount    int64 `json:"active_tutors_count"`
	PendingTutorsCount   int64 `json:"pending_tutors_count"`
	Certificates         int64 `json:"certificates"`
	Applications         int64 `json:"applications"`
	Courses              int64 `json:"courses"`
	CompletedAssessments int64 `json:"completed_assessments"`
}
