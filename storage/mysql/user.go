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

	// Sum GrossSalaryPaid
	err := edb.db.Table("payrolls").
		Select("sum(gross_salary) as gross_salary_earned").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	// Sum NetSalaryPaid
	err = edb.db.Table("payrolls").
		Select("sum(net_salary) as net_salary_earned").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	// Sum PensionPaid
	err = edb.db.Table("payrolls").
		Select("sum(pension) as pension_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	// Sum PayePaid
	err = edb.db.Table("payrolls").
		Select("sum(paye) as paye_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	// Sum NsitfPaid
	err = edb.db.Table("payrolls").
		Select("sum(nsitf) as nsitf_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	// Sum NhfPaid
	err = edb.db.Table("payrolls").
		Select("sum(nhf) as nhf_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	// Sum ItfPaid
	err = edb.db.Table("payrolls").
		Select("sum(itf) as itf_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payment_status=? AND user_id=?", ut.Success, userID).
		Scan(&data).Error

	if err != nil {
		fmt.Println(err)
	}

	return &data
}

// Authenticate a user
func (edb *DBUserStorage) Authenticate(email string) (*md.User, error) {
	user := &md.User{}
	q := edb.db.
		Preload("Salary")

	err := q.Where("email = ? AND is_email_verified = ?", email, true).First(&user).Error

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
		Preload("Salary").
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
		Preload("Salary").
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
		Preload("Salary").
		Where("status=?", ut.Active).
		Find(&user)

	return user
}

//GetUserByEmail Gets user by Email
func (edb *DBUserStorage) GetUserByEmail(email string) *md.User {
	user := md.User{}

	// Select User
	err := edb.db.
		Preload("Salary").
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

	q := edb.db.
		Preload("Salary")

	pagination.Paging(&pagination.Param{
		DB:      q.Find(&users),
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
