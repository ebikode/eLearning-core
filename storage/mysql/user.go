package storage

import (
	"fmt"

	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

// DBUserStorage ...
type DBUserStorage struct {
	*MDatabase
}

// NewDBUserStorage Initialize User Storage
func NewDBUserStorage(db *MDatabase) *DBUserStorage {
	return &DBUserStorage{db}
}

// GetDashbordData ...
func (edb *DBUserStorage) GetDashbordData(userID string) *md.UserDashbordData {
	data := md.UserDashbordData{}

	// Certificates
	err := edb.db.Table("applications").
		Select("count(distinct(id)) as certificates").
		Where("is_certificate_issued=? AND user_id=?", true, userID).
		Scan(&data).Error

	// Applications
	err = edb.db.Table("applications").
		Select("count(distinct(id)) as applications").
		Where("user_id=?", userID).
		Scan(&data).Error

	// Completed Assessments
	err = edb.db.Table("applications").
		Select("count(distinct(id)) as completed_assessments").
		Where("is_assessment_completed=? AND user_id=?", true, userID).
		Scan(&data).Error

	// Number Courses Applied For
	err = edb.db.Table("applications").
		Select("count(distinct(co.id)) as courses").
		Joins("JOIN courses as co ON co.id = applications.course_id").
		Where("applications.user_id=?", userID).
		Scan(&data).Error

	if err != nil {
		fmt.Println(err)
	}

	return &data
}

// GetTutorDashbordData ...
func (edb *DBUserStorage) GetTutorDashbordData(userID string) *md.TutorDashbordData {
	data := md.TutorDashbordData{}

	// Certificates
	err := edb.db.Table("applications").
		Select("count(distinct(applications.id)) as certificates").
		Joins("JOIN courses as co ON co.id = applications.course_id").
		Where("applications.is_certificate_issued=? AND co.user_id=?", true, userID).
		Scan(&data).Error

	// Applications
	err = edb.db.Table("applications").
		Select("count(distinct(applications.id)) as applications").
		Joins("JOIN courses as co ON co.id = applications.course_id").
		Where("co.user_id=?", userID).
		Scan(&data).Error

	// Completed Assessments
	err = edb.db.Table("applications").
		Select("count(distinct(applications.id)) as completed_assessments").
		Joins("JOIN courses as co ON co.id = applications.course_id").
		Where("applications.is_assessment_completed=? AND co.user_id=?", true, userID).
		Scan(&data).Error

	// Number Courses Applied For
	err = edb.db.Table("courses").
		Select("count(distinct(id)) as courses").
		Where("user_id=?", userID).
		Scan(&data).Error

	if err != nil {
		fmt.Println(err)
	}

	return &data
}

// Authenticate a user
func (edb *DBUserStorage) Authenticate(email string) (*md.User, error) {
	user := &md.User{}

	err := edb.db.Where("email = ? AND is_email_verified = ?", email, true).First(&user).Error

	if user.ID == "" || err != nil {
		return nil, err
	}
	return user, nil
}

// Get user by ID
func (edb *DBUserStorage) Get(id string) *md.User {
	user := md.User{}
	// Select User
	err := edb.db.
		Where("id=?", id).
		First(&user).Error

	if user.ID == "" || err != nil {
		return nil
	}

	return &user
}

//GetPubUser Gets public user by ID
func (edb *DBUserStorage) GetPubUser(id string) *md.PubUser {
	user := md.PubUser{}
	// Select User
	err := edb.db.
		Where("id=?", id).
		First(&user).Error

	if user.ID == "" || err != nil {
		return nil
	}

	return &user
}

//GetActivePubUsers Gets public user by ID
func (edb *DBUserStorage) GetActivePubUsers() []*md.PubUser {
	user := []*md.PubUser{}
	// Select User
	edb.db.
		Where("status=?", ut.Active).
		Find(&user)

	return user
}

//GetUserByEmail Gets user by Email
func (edb *DBUserStorage) GetUserByEmail(email string) *md.User {
	user := md.User{}

	// Select User
	err := edb.db.
		Where("email=?", email).
		First(&user).Error

	if user.ID == "" || err != nil {
		return nil
	}

	return &user
}

// GetUsers gets all users with paging using @param page and limit
func (edb *DBUserStorage) GetUsers(page, limit int) []*md.PubUser {
	var users []*md.PubUser

	pagination.Paging(&pagination.Param{
		DB:      edb.db.Find(&users),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &users)

	return users
}

//Store a new user
func (edb *DBUserStorage) Store(u md.User) (md.User, error) {

	usr := &u

	err := edb.db.Create(&usr).Error

	if err != nil {
		return u, err
	}
	return u, nil
}

// Update a user
func (edb *DBUserStorage) Update(u *md.User) (*md.User, error) {

	err := edb.db.Save(&u).Error

	if err != nil {
		return u, err
	}

	return u, nil
}

// Delete a user
func (edb *DBUserStorage) Delete(u md.User, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = edb.db.Unscoped().Delete(u).Error
	}
	if !isPermarnant {
		err = edb.db.Delete(u).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
