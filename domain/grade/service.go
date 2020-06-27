package grade

import (
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
)

// GradeService provides grade operations
type GradeService interface {
	GetGradeReports(int, int) []*md.GradeReport
	GetGrade(uint) *md.Grade
	GetGrades(int, int) []*md.Grade
	GetGradesByApplication(int) *md.Grade
	GetUserGrades(string) []*md.Grade
	GetCourseGrades(string, int, int, int) []*md.Grade
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
func (s *service) GetGrade(gradeID uint) *md.Grade {
	return s.pRepo.Get(gradeID)
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

/*
* Get Grade Reports for each months
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetGradeReports(page, limit int) []*md.GradeReport {
	return s.pRepo.GetReports(page, limit)
}

/*
* Get all grades of a user
* @param userID => the ID of the user whose grade is needed
 */
func (s *service) GetUserGrades(userID string) []*md.Grade {
	return s.pRepo.GetByUser(userID)
}

/*
* Get all grades of a course
* @param userID => the ID of the user who own the course
* @param courseID => the ID of the course whose grade is needed
* @param page => the page number to return
* @param limit => limit per page to return
 */
func (s *service) GetCourseGrades(userID string, courseID, page, limit int) []*md.Grade {
	return s.pRepo.GetByCourse(userID, courseID, page, limit)
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
