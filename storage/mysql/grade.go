package storage

import (
	"fmt"

	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

// DBGradeStorage encapsulates DB Connection Model
type DBGradeStorage struct {
	*MDatabase
}

// NewDBGradeStorage Initialize Grade Storage
func NewDBGradeStorage(db *MDatabase) *DBGradeStorage {
	return &DBGradeStorage{db}
}

// GetReports - Fetches Grade Monthly count data
func (gdb *DBGradeStorage) GetReports(page, limit int) []*md.GradeReport {
	var reports []*md.GradeReport
	offset := (limit * page) - limit
	// Count Grades Monthly Data
	err := gdb.db.Table("grades").
		Select("SUM(if(grade = 'A', 1,0)) as a_count, SUM(if(grade = 'B', 1,0)) as b_count, SUM(if(grade = 'C', 1,0)) as c_count, SUM(if(grade = 'D', 1,0)) as d_count, SUM(if(grade = 'e', 1,0)) as e_count, SUM(if(grade = 'F', 1,0)) as f_count, MONTH(created_at) as month, MONTHNAME(created_at) as month_name, YEAR(created_at) as year").
		Group("YEAR(created_at), MONTH(created_at) DESC").
		Offset(offset).
		Limit(limit).
		Scan(&reports).Error

	if err != nil {
		fmt.Println(err)
	}

	return reports
}

// Get Fetch Single Grade fron DB
func (gdb *DBGradeStorage) Get(id uint) *md.Grade {
	grade := md.Grade{}
	// Select resource from database
	err := gdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Joins("JOIN courses as pr ON pr.id = grades.course_id").
		Where("grades.id=?", id).First(&grade).Error

	if grade.ID < 1 || err != nil {
		return nil
	}

	return &grade
}

// GetAll Fetch all grades from DB
func (gdb *DBGradeStorage) GetAll(page, limit int) []*md.Grade {
	var grades []*md.Grade

	pagination.Paging(&pagination.Param{
		DB: gdb.db.
			Preload("Application").
			Preload("Application.Course").
			Preload("Application.User").
			Joins("JOIN courses as pr ON pr.id = grades.course_id").
			Order("created_at desc").
			Find(&grades),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &grades)

	return grades

}

// GetUserGrades Fetch all user' grades from DB
func (gdb *DBGradeStorage) GetUserGrades(userID string) []*md.Grade {
	var grades []*md.Grade

	gdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Joins("JOIN courses as pr ON pr.id = grades.course_id").
		Where("user_id=?", userID).
		Find(&grades)
	return grades
}

// GetByApplication ...
func (gdb *DBGradeStorage) GetByApplication(applicationID int) *md.Grade {
	var grade *md.Grade
	// Select resource from database
	gdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Where("application_id=?", applicationID).Order("created_at desc").First(&grade)

	return grade
}

// Store Add a new grade
func (gdb *DBGradeStorage) Store(p md.Grade) (*md.Grade, error) {

	grade := p

	err := gdb.db.Create(&grade).Error

	if err != nil {
		return nil, err
	}
	return gdb.Get(grade.ID), nil
}

// Update a grade
func (gdb *DBGradeStorage) Update(grade *md.Grade) (*md.Grade, error) {

	err := gdb.db.Save(&grade).Error

	if err != nil {
		return nil, err
	}

	return grade, nil
}

// Delete a grade
func (gdb *DBGradeStorage) Delete(c md.Grade, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = gdb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = gdb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
