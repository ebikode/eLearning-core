package storage

import (
	md "github.com/ebikode/eLearning-core/model"
)

type DBAppSettingStorage struct {
	*MDatabase
}

// Initialize AppSetting Storage
func NewDBAppSettingStorage(db *MDatabase) *DBAppSettingStorage {
	return &DBAppSettingStorage{db}
}

func (cdb *DBAppSettingStorage) Get(id uint) *md.AppSetting {
	appSetting := md.AppSetting{}
	// Select resource from database
	err := cdb.db.Where("id=?", id).First(&appSetting).Error

	if appSetting.ID < 1 || err != nil {
		return nil
	}

	return &appSetting
}

// Get appSetting By language key
func (cdb *DBAppSettingStorage) GetByKey(key, userType string) *md.AppSetting {
	appSetting := md.AppSetting{}

	var err error
	// Select resource from database
	if userType == "admin" {
		err = cdb.db.Where("s_key = ?", key).First(&appSetting).Error
	} else {
		err = cdb.db.Where("s_key = ? AND status=?", key, "active").First(&appSetting).Error

	}
	if appSetting.ID < 1 || err != nil {
		return nil
	}

	return &appSetting
}

// Get all appSettings
func (cdb *DBAppSettingStorage) GetAll(userType string) []*md.AppSetting {
	var appSettings []*md.AppSetting
	// Select resource from database
	if userType == "admin" {
		cdb.db.Order("name asc").Find(&appSettings)
	} else {
		cdb.db.Where("status=?", "active").Order("name asc").Find(&appSettings)
	}
	return appSettings
}

// Add a new appSetting
func (cdb *DBAppSettingStorage) Store(s md.AppSetting) (*md.AppSetting, error) {

	sett := s

	err := cdb.db.Create(&sett).Error

	if err != nil {
		return nil, err
	}
	return cdb.GetByKey(sett.SKey, "admin"), nil
}

// Update a appSetting
func (cdb *DBAppSettingStorage) Update(appSetting *md.AppSetting) (*md.AppSetting, error) {

	err := cdb.db.Save(&appSetting).Error

	if err != nil {
		return nil, err
	}

	return appSetting, nil
}

// Delete a appSetting
func (cdb *DBAppSettingStorage) Delete(c md.AppSetting, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = cdb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = cdb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
