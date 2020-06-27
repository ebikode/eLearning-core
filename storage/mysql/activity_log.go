package storage

import (
	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

type DBActivityLogStorage struct {
	*MDatabase
}

// Initialize ActivityLog Storage
func NewDBActivityLogStorage(db *MDatabase) *DBActivityLogStorage {
	return &DBActivityLogStorage{db}
}

// Get all activityLogs
func (aldb *DBActivityLogStorage) GetAll(page, limit int) []*md.ActivityLog {
	var activityLogs []*md.ActivityLog

	q := aldb.db.Preload("Admin").Order("created_at desc")
	pagination.Paging(&pagination.Param{
		DB:      q.Find(&activityLogs),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &activityLogs)

	return activityLogs
}

// Add a new actegory
func (aldb *DBActivityLogStorage) Store(c md.ActivityLog) error {

	act := c

	err := aldb.db.Create(&act).Error

	if err != nil {
		return err
	}
	return nil
}

// Delete a actegory
func (aldb *DBActivityLogStorage) Delete(c md.ActivityLog, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = aldb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = aldb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
