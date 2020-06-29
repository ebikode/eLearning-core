package storage

import (
	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

// DBApplicationStorage encapsulates DB Connection Model
type DBApplicationStorage struct {
	*MDatabase
}

// NewDBApplicationStorage Initialize Application Storage
func NewDBApplicationStorage(db *MDatabase) *DBApplicationStorage {
	return &DBApplicationStorage{db}
}

// Get Fetch Single Application fron DB
func (adb *DBApplicationStorage) Get(id uint) *md.Application {
	application := md.Application{}
	// Select resource from database
	err := adb.db.
		Preload("Course").
		Preload("Grade").
		Preload("User").
		Where("applications.id=?", id).First(&application).Error

	if application.ID < 1 || err != nil {
		return nil
	}

	return &application
}

// GetAll Fetch all applications from DB
func (adb *DBApplicationStorage) GetAll(page, limit int) []*md.Application {
	var applications []*md.Application

	pagination.Paging(&pagination.Param{
		DB: adb.db.
			Preload("Course").
			Preload("Grade").
			Preload("User").
			Order("created_at desc").
			Find(&applications),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &applications)

	return applications

}

// GetByUser Fetch all user' applications from DB
func (adb *DBApplicationStorage) GetByUser(userID string) []*md.Application {
	var applications []*md.Application

	adb.db.
		Preload("Course").
		Preload("Grade").
		Preload("User").
		Where("user_id=?", userID).
		Find(&applications)
	return applications
}

// GetByCourse ...
func (adb *DBApplicationStorage) GetByCourse(courseID int) []*md.Application {
	var applications []*md.Application
	// Select resource from database
	adb.db.
		Preload("Course").
		Preload("Grade").
		Preload("User").
		Where("course_id=?", courseID).Order("created_at desc").Find(&applications)

	return applications
}

// GetByCourse ...
func (adb *DBApplicationStorage) GetByCourseOwner(userID string) []*md.Application {
	var applications []*md.Application
	// Select resource from database
	adb.db.
		Preload("Course").
		Preload("Grade").
		Preload("User").
		Joins("JOIN courses as co ON co.id = applications.course_id").
		Where("co.user_id=?", userID).
		Order("created_at desc").Find(&applications)

	return applications
}

// Store Add a new application
func (adb *DBApplicationStorage) Store(p md.Application) (*md.Application, error) {

	application := p

	err := adb.db.Create(&application).Error

	if err != nil {
		return nil, err
	}
	return adb.Get(application.ID), nil
}

// Update a application
func (adb *DBApplicationStorage) Update(application *md.Application) (*md.Application, error) {

	err := adb.db.Save(&application).Error

	if err != nil {
		return nil, err
	}

	return application, nil
}

// Delete a application
func (adb *DBApplicationStorage) Delete(c md.Application, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = adb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = adb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
