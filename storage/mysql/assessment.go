package storage

import (
	"fmt"

	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

// DBAssessmentStorage ...
type DBAssessmentStorage struct {
	*MDatabase
}

// NewDBAssessmentStorage Initialize Assessment Storage
func NewDBAssessmentStorage(db *MDatabase) *DBAssessmentStorage {
	return &DBAssessmentStorage{db}
}

// Get assessment using application_id and assessment_id
func (asdb *DBAssessmentStorage) Get(applicationID uint, assessmentID string) *md.Assessment {
	assessment := md.Assessment{}
	// Select resource from database
	err := asdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Preload("Question").
		Where("application_id=? AND id=?", applicationID, assessmentID).First(&assessment).Error

	if len(assessment.ID) < 1 || err != nil {
		return nil
	}

	return &assessment
}

// GetLastAssessment ...
func (asdb *DBAssessmentStorage) GetLastAssessment() *md.Assessment {
	assessment := md.Assessment{}
	// Select resource from database
	err := asdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Preload("Question").
		Order("created_at").
		Limit(1).First(&assessment).Error

	if len(assessment.ID) < 1 || err != nil {
		return nil
	}

	return &assessment
}

// GetAll Get all assessments
func (asdb *DBAssessmentStorage) GetAll(page, limit int) []*md.Assessment {
	var assessments []*md.Assessment
	// Select resource from database
	q := asdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Preload("Question")

	pagination.Paging(&pagination.Param{
		DB:      q.Order("created_at desc").Find(&assessments),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &assessments)

	return assessments
}

// GetByApplication Get all assessments of a application  form DB
func (asdb *DBAssessmentStorage) GetByApplication(userID string, applicationID uint, page, limit int) []*md.Assessment {
	var assessments []*md.Assessment
	// Select resource from database
	q := asdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Preload("Question")

	pagination.Paging(&pagination.Param{
		DB: q.
			Joins("JOIN applications as app ON app.id = assessments.application_id").
			Where("app.user_id=? AND assessments.application_id=?", userID, applicationID).
			Order("created_at desc").
			Find(&assessments),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &assessments)
	return assessments
}

// GetByCourse ...
func (asdb *DBAssessmentStorage) GetByCourse(courseID int) []*md.Assessment {
	var assessments []*md.Assessment
	// Select resource from database
	asdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Preload("Question").
		Where("course_id=?", courseID).Order("created_at desc").Find(&assessments)

	return assessments
}

// GetSingleByCourse ...
func (asdb *DBAssessmentStorage) GetSingleByCourse(courseID int) *md.Assessment {
	assessment := md.Assessment{}
	// Select resource from database
	err := asdb.db.
		Preload("Application").
		Preload("Application.Course").
		Preload("Application.User").
		Preload("Question").
		Where("course_id=?", courseID).Order("created_at desc").
		First(&assessment).Error

	if len(assessment.ID) < 1 || err != nil {
		return nil
	}

	return &assessment
}

// Store Add a new assessment
func (asdb *DBAssessmentStorage) Store(a md.Assessment) (*md.Assessment, error) {
	fmt.Println("STORE HITS")

	ase := a
	fmt.Println(ase)

	err := asdb.db.Create(&ase).Error

	if err != nil {
		return nil, err
	}

	fmt.Println("STORE HITS")
	return asdb.Get(ase.ApplicationID, ase.ID), nil
}

// Update a assessment
func (asdb *DBAssessmentStorage) Update(assessment *md.Assessment) (*md.Assessment, error) {

	err := asdb.db.Save(&assessment).Error

	if err != nil {
		return nil, err
	}

	return assessment, nil
}

// Delete a assessment
func (asdb *DBAssessmentStorage) Delete(p *md.Assessment, isPermarnant bool) (bool, error) {

	// var err error
	// if isPermarnant {
	// 	err = asdb.db.Unscoped().Delete(p).Error
	// }
	// if !isPermarnant {
	// 	err = asdb.db.Delete(p).Error
	// }

	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}
