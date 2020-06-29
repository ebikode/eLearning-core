package model

// ActivityLog a struct to encapsulate activities performmed by admins
type ActivityLog struct {
	BaseModel
	AdminID     string `json:"admin_id" gorm:"not null;type:varchar(20)"`
	AppLocation string `json:"app_location" gorm:"not null;type:varchar(150)"`
	Action      string `json:"action" gorm:"not null;type:text(5000)"`
	Admin       *Admin `json:"admin"`
}
