package journal

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// JournalService  provides journal operations
type JournalService interface {
	GetJournal(uint) *md.Journal
	GetJournalsByCourse(int) []*md.Journal
	GetJournalsByUser(string, int, int) []*md.Journal
	GetJournals(int, int) []*md.Journal
	CreateJournal(md.Journal) (*md.Journal, tr.TParam, error)
	UpdateJournal(*md.Journal) (*md.Journal, tr.TParam, error)
}

type service struct {
	qRepo JournalRepository
}

// NewService creates a journal service with the necessary dependencies
func NewService(
	qRepo JournalRepository,
) JournalService {
	return &service{qRepo}
}

// Get a journal
func (s *service) GetJournal(id uint) *md.Journal {
	return s.qRepo.Get(id)
}

// GetJournals Get all journals from DB
//
// @userType == admin | customer
func (s *service) GetJournals(page, limit int) []*md.Journal {
	return s.qRepo.GetAll(page, limit)
}

func (s *service) GetJournalsByCourse(courseID int) []*md.Journal {
	return s.qRepo.GetByCourse(courseID)
}

func (s *service) GetJournalsByUser(userID string, page, limit int) []*md.Journal {
	return s.qRepo.GetByUser(userID, page, limit)
}

// CreateJournal Creates New journal
func (s *service) CreateJournal(c md.Journal) (*md.Journal, tr.TParam, error) {

	journal, err := s.qRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return journal, tParam, err
	}

	return journal, tr.TParam{}, nil

}

// UpdateJournal update existing journal
func (s *service) UpdateJournal(c *md.Journal) (*md.Journal, tr.TParam, error) {
	journal, err := s.qRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return journal, tParam, err
	}

	return journal, tr.TParam{}, nil

}

// Validate Function for validating journal input
func Validate(journal Payload, r *http.Request) error {
	return validation.ValidateStruct(&journal,
		validation.Field(&journal.CourseID, ut.IDRule(r)...),
		validation.Field(&journal.Title, ut.RequiredRule(r, "general.title")...),
		validation.Field(&journal.Body, ut.RequiredRule(r, "general.body")...),
	)
}

// ValidateUpdates Function for validating journal update input
func ValidateUpdates(journal Payload, r *http.Request) error {
	return validation.ValidateStruct(&journal,
		validation.Field(&journal.Title, ut.RequiredRule(r, "general.title")...),
		validation.Field(&journal.Body, ut.RequiredRule(r, "general.body")...),
	)
}
