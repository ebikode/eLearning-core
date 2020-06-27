package storage

import (
	"github.com/biezhi/gorm-paginator/pagination"
	md "github.com/ebikode/eLearning-core/model"
)

// DBArticleStorage encapsulates DB Connection Model
type DBArticleStorage struct {
	*MDatabase
}

// NewDBArticleStorage Initialize Article Storage
func NewDBArticleStorage(db *MDatabase) *DBArticleStorage {
	return &DBArticleStorage{db}
}

// Get Fetch Single Article fron DB
func (adb *DBArticleStorage) Get(id uint) *md.Article {
	article := md.Article{}
	// Select resource from database
	err := adb.db.
		Preload("User").
		Where("articles.id=?", id).First(&article).Error

	if article.ID < 1 || err != nil {
		return nil
	}

	return &article
}

// GetByUserID Fetch Single Article fron DB
func (adb *DBArticleStorage) GetByUserID(id string) *md.Article {
	article := md.Article{}
	// Select resource from database
	err := adb.db.
		Preload("User").
		Where("user_id=?", id).First(&article).Error

	if article.ID < 1 || err != nil {
		return nil
	}

	return &article
}

// GetAll Fetch all articles from DB
func (adb *DBArticleStorage) GetAll(page, limit int) []*md.Article {
	var articles []*md.Article

	pagination.Paging(&pagination.Param{
		DB: adb.db.
			Preload("User").
			Order("created_at desc").
			Find(&articles),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &articles)

	return articles

}

// GetByUser Fetch all user' articles from DB
func (adb *DBArticleStorage) GetByUser(userID string) []*md.Article {
	var articles []*md.Article

	adb.db.
		Preload("User").
		Where("user_id=?", userID).
		Find(&articles)
	return articles
}

// Store Add a new article
func (adb *DBArticleStorage) Store(p md.Article) (*md.Article, error) {

	article := p

	err := adb.db.Create(&article).Error

	if err != nil {
		return nil, err
	}
	return adb.Get(article.ID), nil
}

// Update a article
func (adb *DBArticleStorage) Update(article *md.Article) (*md.Article, error) {

	err := adb.db.Save(&article).Error

	if err != nil {
		return nil, err
	}

	return article, nil
}

// Delete a article
func (adb *DBArticleStorage) Delete(c md.Article, isPermarnant bool) (bool, error) {

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
