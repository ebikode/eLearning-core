package model

// AuthdDevice a struct to rep Authrnticated Device
type AuthdDevice struct {
	BaseModel
	UserID         string `json:"user_id" gorm:"not null;type:varchar(20)"`
	IP             string `json:"ip" gorm:"type:varchar(50)"`
	Browser        string `json:"browser" gorm:"type:varchar(100)"`
	BrowserVersion string `json:"browser_version" gorm:"type:varchar(50)"`
	Platform       string `json:"platform" gorm:"type:varchar(50)"`
	DeviceOS       string `json:"device_os" gorm:"type:varchar(50)"`
	OSVersion      string `json:"os_version" gorm:"type:varchar(50)"`
	DeviceType     string `json:"device_type" gorm:"type:varchar(50)"`
	AccessType     string `json:"access_type" gorm:"type:enum('','web_app','mobile_app');default:''"`
	Status         string `json:"status" gorm:"type:enum('active','logout');default:'active'"`
	User           *User  `json:"user"`
}
