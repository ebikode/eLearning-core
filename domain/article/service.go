package article

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ArticleService  provides article operations
type ArticleService interface {
	GetArticle(uint) *md.Article
	GetArticlesByCourse(int) []*md.Article
	GetArticlesByUser(string, int, int) []*md.Article
	GetArticles(int, int) []*md.Article
	CreateArticle(md.Article) (*md.Article, tr.TParam, error)
	UpdateArticle(*md.Article) (*md.Article, tr.TParam, error)
}

type service struct {
	qRepo ArticleRepository
}

// NewService creates a article service with the necessary dependencies
func NewService(
	qRepo ArticleRepository,
) ArticleService {
	return &service{qRepo}
}

// Get a article
func (s *service) GetArticle(id uint) *md.Article {
	return s.qRepo.Get(id)
}

// GetArticles Get all articles from DB
//
// @userType == admin | customer
func (s *service) GetArticles(page, limit int) []*md.Article {
	return s.qRepo.GetAll(page, limit)
}

func (s *service) GetArticlesByCourse(courseID int) []*md.Article {
	return s.qRepo.GetByCourse(courseID)
}

func (s *service) GetArticlesByUser(userID string, page, limit int) []*md.Article {
	return s.qRepo.GetByUser(userID, page, limit)
}

// CreateArticle Creates New article
func (s *service) CreateArticle(c md.Article) (*md.Article, tr.TParam, error) {

	article, err := s.qRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return article, tParam, err
	}

	return article, tr.TParam{}, nil

}

// UpdateArticle update existing article
func (s *service) UpdateArticle(c *md.Article) (*md.Article, tr.TParam, error) {
	article, err := s.qRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return article, tParam, err
	}

	return article, tr.TParam{}, nil

}

// Validate Function for validating article input
func Validate(article Payload, r *http.Request) error {
	return validation.ValidateStruct(&article,
		validation.Field(&article.CourseID, ut.IDRule(r)...),
		validation.Field(&article.Title, ut.RequiredRule(r, "general.title")...),
		validation.Field(&article.Body, ut.RequiredRule(r, "general.body")...),
	)
}

// ValidateUpdates Function for validating article update input
func ValidateUpdates(article Payload, r *http.Request) error {
	return validation.ValidateStruct(&article,
		validation.Field(&article.Title, ut.RequiredRule(r, "general.title")...),
		validation.Field(&article.Body, ut.RequiredRule(r, "general.body")...),
	)
}
