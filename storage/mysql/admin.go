package storage

import (
	"fmt"

	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

// DBAdminStorage ...
type DBAdminStorage struct {
	*MDatabase
}

// NewDBAdminStorage Initialize Admin Storage
func NewDBAdminStorage(db *MDatabase) *DBAdminStorage {
	return &DBAdminStorage{db}
}

// Authenticate an admin
func (adb *DBAdminStorage) Authenticate(email string) (*md.Admin, error) {
	admin := &md.Admin{}

	err := adb.db.Where("email = ?", email).First(&admin).Error

	if admin.ID == "" || err != nil {
		return nil, err
	}
	return admin, nil
}

// GetDashbordData ...
func (adb *DBAdminStorage) GetDashbordData() *md.DashbordData {
	data := md.DashbordData{}
	user := md.User{}

	var result int64

	// count All Users
	adb.db.Model(&user).Count(&result)

	data.UsersCount = result
	result = 0

	// count All Active Users
	adb.db.Model(&user).Where("status=? AND role=?", ut.Active, ut.UserRole).Count(&result)

	data.ActiveUsersCount = result
	result = 0

	// count All Pending Users
	adb.db.Model(&user).Where("status=? AND role=?", ut.Pending, ut.UserRole).Count(&result)

	data.PendingUsersCount = result
	result = 0

	// count All Active Tutors
	adb.db.Model(&user).Where("status=? AND role=?", ut.Active, ut.TutorRole).Count(&result)

	data.ActiveTutorsCount = result
	result = 0

	// count All Pending Tutors
	adb.db.Model(&user).Where("status=? AND role=?", ut.Pending, ut.TutorRole).Count(&result)

	data.PendingTutorsCount = result
	result = 0

	// Certificates
	err := adb.db.Table("applications").
		Select("count(distinct(id)) as certificates").
		Where("is_certificate_issued=?", true).
		Scan(&data).Error

	// Applications
	err = adb.db.Table("applications").
		Select("count(distinct(id)) as applications").
		Scan(&data).Error

	// Completed Assessments
	err = adb.db.Table("applications").
		Select("count(distinct(id)) as completed_assessments").
		Where("is_assessment_completed=?", true).
		Scan(&data).Error

	// Number Courses Applied For
	err = adb.db.Table("courses").
		Select("count(distinct(id)) as courses").
		Scan(&data).Error

	if err != nil {
		fmt.Println(err)
	}

	return &data
}

// Get ...
func (adb *DBAdminStorage) Get(id string) *md.Admin {
	admin := md.Admin{}
	// Select Admin
	err := adb.db.Where("id=?", id).First(&admin).Error

	if admin.ID == "" || err != nil {
		return nil
	}
	admin.Password = ""
	return &admin
}

// CheckAdminCreated - Checks if a default admin has already been created
// used when the server is ran for the very first time so as to create
// a default admin if it returns false
func (adb *DBAdminStorage) CheckAdminCreated() bool {
	admin := md.Admin{}
	// Select Admin
	err := adb.db.First(&admin).Error

	if admin.ID == "" || err != nil {
		return false
	}

	return true
}

// Store Add a new admin
func (adb *DBAdminStorage) Store(u md.Admin) (md.Admin, error) {

	usr := &u

	err := adb.db.Create(&usr).Error

	if err != nil {
		return u, err
	}
	return u, nil
}

// Update a admin
func (adb *DBAdminStorage) Update(u *md.Admin) (*md.Admin, error) {

	err := adb.db.Save(&u).Error

	if err != nil {
		return u, err
	}

	return u, nil
}

// Delete a admin
func (adb *DBAdminStorage) Delete(u md.Admin, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = adb.db.Unscoped().Delete(u).Error
	}
	if !isPermarnant {
		err = adb.db.Delete(u).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
