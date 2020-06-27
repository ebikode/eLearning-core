package storage

import (
	md "github.com/ebikode/eLearning-core/model"
)

type DBAuthdDeviceStorage struct {
	*MDatabase
}

// Initialize AuthdDevice Storage
func NewDBAuthdDeviceStorage(db *MDatabase) *DBAuthdDeviceStorage {
	return &DBAuthdDeviceStorage{db}
}

func (audb *DBAuthdDeviceStorage) Get(id string) *md.AuthdDevice {
	aud := md.AuthdDevice{}
	// Select AuthdDevice
	err := audb.db.
		Preload("User").
		Where("id=?", id).
		First(&aud).Error

	if aud.ID == "" || err != nil {
		return nil
	}
	return &aud
}

// Store Add a new aud
func (audb *DBAuthdDeviceStorage) Store(a md.AuthdDevice) (md.AuthdDevice, error) {

	aud := a
	// Create new auth device
	err := audb.db.Create(&aud).Error

	if err != nil {
		return a, err
	}
	return a, nil
}

// Update a aud
func (audb *DBAuthdDeviceStorage) Update(a *md.AuthdDevice) (*md.AuthdDevice, error) {

	aud := a

	err := audb.db.Save(&aud).Error

	if err != nil {
		return a, err
	}
	audb.db.Model(&aud).Related(&aud.User, "User")

	return a, nil
}

// Delete a aud
func (audb *DBAuthdDeviceStorage) Delete(a *md.AuthdDevice, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = audb.db.Unscoped().Delete(a).Error
	}
	if !isPermarnant {
		err = audb.db.Delete(a).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
