package model

// AppSetting a struct to rep app settings account
type AppSetting struct {
	BaseIntModel
	Name    string `json:"name" gorm:"not null;type:varchar(50);"`
	SKey    string `json:"s_key" gorm:"not null;type:varchar(100);unique_index"`
	Value   string `json:"value" gorm:"not null;type:varchar(200)"`
	Comment string `json:"comment" gorm:"type:varchar(200)"`
	Status  string `json:"status" gorm:"type:enum('active','disabled');default:'active'"`
}
