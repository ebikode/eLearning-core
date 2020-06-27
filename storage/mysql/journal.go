package storage

import (
	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

// DBJournalStorage encapsulates DB Connection Model
type DBJournalStorage struct {
	*MDatabase
}

// NewDBJournalStorage Initialize Journal Storage
func NewDBJournalStorage(db *MDatabase) *DBJournalStorage {
	return &DBJournalStorage{db}
}

// Get Fetch Single Journal fron DB
func (jdb *DBJournalStorage) Get(id uint) *md.Journal {
	journal := md.Journal{}
	// Select resource from database
	err := jdb.db.
		Preload("Course").
		Preload("User").
		Where("journals.id=?", id).First(&journal).Error

	if journal.ID < 1 || err != nil {
		return nil
	}

	return &journal
}

// GetByUserID Fetch Single Journal fron DB
func (jdb *DBJournalStorage) GetByUserID(id string) *md.Journal {
	journal := md.Journal{}
	// Select resource from database
	err := jdb.db.
		Preload("Course").
		Preload("User").
		Where("user_id=?", id).First(&journal).Error

	if journal.ID < 1 || err != nil {
		return nil
	}

	return &journal
}

// GetAll Fetch all journals from DB
func (jdb *DBJournalStorage) GetAll(page, limit int) []*md.Journal {
	var journals []*md.Journal

	pagination.Paging(&pagination.Param{
		DB: jdb.db.
			Preload("Course").
			Preload("User").
			Order("created_at desc").
			Find(&journals),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &journals)

	return journals

}

// GetByUser Fetch all user' journals from DB
func (jdb *DBJournalStorage) GetByUser(userID string, page, limit int) []*md.Journal {
	var journals []*md.Journal

	pagination.Paging(&pagination.Param{
		DB: jdb.db.
			Preload("Course").
			Preload("User").
			Where("user_id=?", userID).
			Find(&journals),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &journals)
	return journals
}

// GetByCourse ...
func (jdb *DBJournalStorage) GetByCourse(courseID int) []*md.Journal {
	var journals []*md.Journal
	// Select resource from database
	jdb.db.
		Preload("Course").
		Preload("User").
		Where("course_id=?", courseID).Order("created_at desc").Find(&journals)

	return journals
}

// Store Add a new journal
func (jdb *DBJournalStorage) Store(p md.Journal) (*md.Journal, error) {

	journal := p

	err := jdb.db.Create(&journal).Error

	if err != nil {
		return nil, err
	}
	return jdb.Get(journal.ID), nil
}

// Update a journal
func (jdb *DBJournalStorage) Update(journal *md.Journal) (*md.Journal, error) {

	err := jdb.db.Save(&journal).Error

	if err != nil {
		return nil, err
	}

	return journal, nil
}

// Delete a journal
func (jdb *DBJournalStorage) Delete(c md.Journal, isPermarnant bool) (bool, error) {

	var err error
	if isPermarnant {
		err = jdb.db.Unscoped().Delete(c).Error
	}
	if !isPermarnant {
		err = jdb.db.Delete(c).Error
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
