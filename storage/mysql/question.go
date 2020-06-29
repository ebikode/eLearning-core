package storage

import (
	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

// DBQuestionStorage encapsulates DB Connection Model
type DBQuestionStorage struct {
	*MDatabase
}

// NewDBQuestionStorage Initialize Question Storage
func NewDBQuestionStorage(db *MDatabase) *DBQuestionStorage {
	return &DBQuestionStorage{db}
}

// CountByCourse ...
func (adb *DBQuestionStorage) CountByCourse(courseID uint) int {
	var count int

	adb.db.Table("questions").
		Where("course_id=?", courseID).
		Count(&count)
	return count
}

// Get Fetch Single Question fron DB
func (adb *DBQuestionStorage) Get(id string) *md.Question {
	question := md.Question{}
	// Select resource from database
	err := adb.db.
		Preload("Course").
		Where("id=?", id).First(&question).Error

	if len(question.ID) < 1 || err != nil {
		return nil
	}

	return &question
}

// GetSingleByCourseID Fetch Single Question fron DB
func (adb *DBQuestionStorage) GetSingleByCourseID(id string) *md.Question {
	question := md.Question{}
	// Select resource from database
	err := adb.db.
		Preload("Course").
		Where("course_id=?", id).First(&question).Error

	if len(question.ID) < 1 || err != nil {
		return nil
	}

	return &question
}

// GetAll Fetch all questions from DB
func (adb *DBQuestionStorage) GetAll(page, limit int) []*md.Question {
	var questions []*md.Question

	pagination.Paging(&pagination.Param{
		DB: adb.db.
			Preload("Course").
			Order("created_at desc").
			Find(&questions),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &questions)

	return questions

}

// GetByCourse Fetch all course' questions from DB
func (adb *DBQuestionStorage) GetByCourse(courseID uint) []*md.PubQuestion {
	var questions []*md.PubQuestion

	adb.db.
		Preload("Course").
		Where("course_id=? AND status=?", courseID, ut.Active).
		Find(&questions)
	return questions
}

// GetByCourseOwner Fetch all course' questions from DB
func (adb *DBQuestionStorage) GetByCourseOwner(userID string, courseID uint) []*md.Question {
	var questions []*md.Question

	adb.db.
		Preload("Course").
		Where("course_id=?", courseID).
		Find(&questions)
	return questions
}

// Store Add a new question
func (adb *DBQuestionStorage) Store(p md.Question) (*md.Question, error) {

	question := p

	err := adb.db.Create(&question).Error

	if err != nil {
		return nil, err
	}
	return adb.Get(question.ID), nil
}

// Update a question
func (adb *DBQuestionStorage) Update(question *md.Question) (*md.Question, error) {

	err := adb.db.Save(&question).Error

	if err != nil {
		return nil, err
	}

	return question, nil
}

// Delete a question
func (adb *DBQuestionStorage) Delete(c md.Question, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = adb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = adb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
