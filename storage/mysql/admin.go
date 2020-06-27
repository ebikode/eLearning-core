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
	adb.db.Model(&user).Where("status=?", "active").Count(&result)

	data.ActiveUsersCount = result
	result = 0

	// count All Pending Users
	adb.db.Model(&user).Where("status=?", "pending").Count(&result)

	data.PendingUsersCount = result
	result = 0

	// Sum GrossSalaryPaid
	err := adb.db.Table("payrolls").
		Select("sum(gross_salary) as gross_salary_paid").
		Where("payment_status=?", ut.Success).
		Scan(&data).Error

	// Sum NetSalaryPaid
	err = adb.db.Table("payrolls").
		Select("sum(net_salary) as net_salary_paid").
		Where("payment_status=?", ut.Success).
		Scan(&data).Error

	// Sum PensionPaid
	err = adb.db.Table("payrolls").
		Select("sum(pension) as pension_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payrolls.payment_status=?", ut.Success).
		Scan(&data).Error

	// Sum PayePaid
	err = adb.db.Table("payrolls").
		Select("sum(paye) as paye_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payrolls.payment_status=?", ut.Success).
		Scan(&data).Error

	// Sum NsitfPaid
	err = adb.db.Table("payrolls").
		Select("sum(nsitf) as nsitf_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payrolls.payment_status=?", ut.Success).
		Scan(&data).Error

	// Sum NhfPaid
	err = adb.db.Table("payrolls").
		Select("sum(nhf) as nhf_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payrolls.payment_status=?", ut.Success).
		Scan(&data).Error

	// Sum ItfPaid
	err = adb.db.Table("payrolls").
		Select("sum(itf) as itf_paid").
		Joins("JOIN taxes as tax ON tax.payroll_id = payrolls.id").
		Where("payrolls.payment_status=?", ut.Success).
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
