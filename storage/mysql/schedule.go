package storage

import (
	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

// DBScheduleStorage encapsulates DB Connection Model
type DBScheduleStorage struct {
	*MDatabase
}

// NewDBScheduleStorage Initialize Schedule Storage
func NewDBScheduleStorage(db *MDatabase) *DBScheduleStorage {
	return &DBScheduleStorage{db}
}

// Get Fetch Single Schedule fron DB
func (jdb *DBScheduleStorage) Get(id uint) *md.Schedule {
	schedule := md.Schedule{}
	// Select resource from database
	err := jdb.db.
		Preload("Course").
		Where("schedules.id=?", id).First(&schedule).Error

	if schedule.ID < 1 || err != nil {
		return nil
	}

	return &schedule
}

// GetByCourseID Fetch Single Schedule fron DB
func (jdb *DBScheduleStorage) GetByCourseID(id string) *md.Schedule {
	schedule := md.Schedule{}
	// Select resource from database
	err := jdb.db.
		Preload("Course").
		Where("course_id=?", id).First(&schedule).Error

	if schedule.ID < 1 || err != nil {
		return nil
	}

	return &schedule
}

// GetAll Fetch all schedules from DB
func (jdb *DBScheduleStorage) GetAll(page, limit int) []*md.Schedule {
	var schedules []*md.Schedule

	pagination.Paging(&pagination.Param{
		DB: jdb.db.
			Preload("Course").
			Order("created_at desc").
			Find(&schedules),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &schedules)

	return schedules

}

// GetByCourse Fetch all course schedules from DB
func (jdb *DBScheduleStorage) GetByCourse(courseID string) []*md.Schedule {
	var schedules []*md.Schedule

	jdb.db.
		Preload("Course").
		Where("course_id=?", courseID).
		Find(&schedules)
	return schedules
}

// Store Add a new schedule
func (jdb *DBScheduleStorage) Store(p md.Schedule) (*md.Schedule, error) {

	schedule := p

	err := jdb.db.Create(&schedule).Error

	if err != nil {
		return nil, err
	}
	return jdb.Get(schedule.ID), nil
}

// Update a schedule
func (jdb *DBScheduleStorage) Update(schedule *md.Schedule) (*md.Schedule, error) {

	err := jdb.db.Save(&schedule).Error

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// Delete a schedule
func (jdb *DBScheduleStorage) Delete(c md.Schedule, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = jdb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = jdb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
