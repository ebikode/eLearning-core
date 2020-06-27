package course

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// CourseService  provides course operations
type CourseService interface {
	GetCourse(uint) *md.Course
	GetCoursesByUser(string) *md.Course
	GetCourses(int, int) []*md.Course
	CreateCourse(md.Course) (*md.Course, tr.TParam, error)
	UpdateCourse(*md.Course) (*md.Course, tr.TParam, error)
}

type service struct {
	qRepo CourseRepository
}

// NewService creates a course service with the necessary dependencies
func NewService(
	qRepo CourseRepository,
) CourseService {
	return &service{qRepo}
}

// Get a course
func (s *service) GetCourse(id uint) *md.Course {
	return s.qRepo.Get(id)
}

// GetCourses Get all courses from DB
//
// @userType == admin | customer
func (s *service) GetCourses(page, limit int) []*md.Course {
	return s.qRepo.GetAll(page, limit)
}

func (s *service) GetCoursesByUser(userID string) *md.Course {
	return s.qRepo.GetByUser(userID)
}

// CreateCourse Creates New course
func (s *service) CreateCourse(c md.Course) (*md.Course, tr.TParam, error) {

	course, err := s.qRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return course, tParam, err
	}

	return course, tr.TParam{}, nil

}

// UpdateCourse update existing course
func (s *service) UpdateCourse(c *md.Course) (*md.Course, tr.TParam, error) {
	course, err := s.qRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return course, tParam, err
	}

	return course, tr.TParam{}, nil

}

// Validate Function for validating course input
func Validate(course Payload, r *http.Request) error {
	return validation.ValidateStruct(&course,
		validation.Field(&course.Title, ut.RequiredRule(r, "general.title")...),
		validation.Field(&course.Description, ut.RequiredRule(r, "general.description")...),
		validation.Field(&course.DurationPerQuestion, ut.RequiredRule(r, "general.duration")...),
	)
}

// ValidateUpdates Function for validating course update input
func ValidateUpdates(course Payload, r *http.Request) error {
	return validation.ValidateStruct(&course,
		validation.Field(&course.Title, ut.RequiredRule(r, "general.title")...),
		validation.Field(&course.Mode, ut.RequiredRule(r, "general.mode")...),
		validation.Field(&course.Description, ut.RequiredRule(r, "general.description")...),
		validation.Field(&course.DurationPerQuestion, ut.RequiredRule(r, "general.duration")...),
	)
}
