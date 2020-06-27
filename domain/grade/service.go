package grade

import (
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
)

// GradeService provides grade operations
type GradeService interface {
	GetGradeReports() []*md.GradeReport
	GetGrade(string, string) *md.Grade
	GetGrades(int, int) []*md.Grade
	GetGradesByApplication(int) *md.Grade
	GetUserGrades(string, int, int) []*md.Grade
	CreateGrade(md.Grade) (*md.Grade, tr.TParam, error)
	UpdateGrade(*md.Grade) (*md.Grade, tr.TParam, error)
}

type service struct {
	pRepo GradeRepository
}

// NewService creates p grade service with the necessary dependencies
func NewService(
	pRepo GradeRepository,
) GradeService {
	return &service{pRepo}
}

/*
* Get single grade of a user
* @param userID => the ID of the user whose grade is needed
* @param gradeID => the ID of the grade requested.
 */
func (s *service) GetGrade(userID, gradeID string) *md.Grade {
	return s.pRepo.Get(userID, gradeID)
}

/*
* Get all grades
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetGrades(page, limit int) []*md.Grade {
	return s.pRepo.GetAll(page, limit)
}

func (s *service) GetGradesByApplication(applicationID int) *md.Grade {
	return s.pRepo.GetByApplication(applicationID)
}

func (s *service) GetGradeReports() []*md.GradeReport {
	return s.pRepo.GetReports()
}

/*
* Get all grades of a user
* @param userID => the ID of the user whose grade is needed
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetUserGrades(userID string, page, limit int) []*md.Grade {
	return s.pRepo.GetByUser(userID, page, limit)
}

// Create New grade
func (s *service) CreateGrade(p md.Grade) (*md.Grade, tr.TParam, error) {

	grade, err := s.pRepo.Store(p)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return grade, tParam, err
	}

	return grade, tr.TParam{}, nil
}

// update existing grade
func (s *service) UpdateGrade(p *md.Grade) (*md.Grade, tr.TParam, error) {
	grade, err := s.pRepo.Update(p)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return grade, tParam, err
	}
	return grade, tr.TParam{}, nil
}
