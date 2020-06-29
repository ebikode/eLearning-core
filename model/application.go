package model

// Application - a struct to rep plan database model
type Application struct {
	BaseIntModel
	UserID                     string   `json:"user_id" gorm:"not null;type:varchar(20)"`
	CourseID                   uint     `json:"course_id" gorm:"not null;type:int(15)"`
	CertificateIssuerID        string   `json:"certificate_issuer_id" gorm:"not null;type:varchar(20)"`
	ReferenceNo                string   `json:"reference_no" gorm:"type:varchar(50)"`
	AssessmentStartTimestamp   int64    `json:"assessment_start_timestamp" gorm:"type:int(20)"`
	AssessmentEndTimestamp     int64    `json:"assessment_end_timestamp" gorm:"type:int(20)"`
	IsAssessmentCompleted      bool     `json:"is_assessment_completed" gorm:"type:tinyint(1);default:0"`
	IsCertificateIssued        bool     `json:"is_certificate_issued" gorm:"type:tinyint(1);default:0"`
	CertificateIssuedTimestamp int64    `json:"certificate_issued_timestamp" gorm:"type:int(20)"`
	Status                     string   `json:"status" gorm:"type:enum('pending','active','suspended','completed');default:'active'"`
	CertificateIssuer          *Admin   `json:"issuer"`
	Grade                      *Grade   `json:"grade" gorm:"foreignKey:ApplicationID"`
	Course                     *Course  `json:"course"`
	User                       *PubUser `json:"user"`
}
