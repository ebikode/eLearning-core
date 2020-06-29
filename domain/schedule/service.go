package schedule

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ScheduleService  provides schedule operations
type ScheduleService interface {
	GetSchedule(uint) *md.Schedule
	GetSchedulesByCourse(uint) []*md.Schedule
	GetSchedulesByCourseOwner(string, uint) []*md.Schedule
	GetSchedules(int, int) []*md.Schedule
	CreateSchedule(md.Schedule) (*md.Schedule, tr.TParam, error)
	UpdateSchedule(*md.Schedule) (*md.Schedule, tr.TParam, error)
}

type service struct {
	sRepo ScheduleRepository
}

// NewService creates a schedule service with the necessary dependencies
func NewService(
	sRepo ScheduleRepository,
) ScheduleService {
	return &service{sRepo}
}

// Get a schedule
func (s *service) GetSchedule(id uint) *md.Schedule {
	return s.sRepo.Get(id)
}

// GetSchedules Get all schedules from DB
//
// @userType == admin | customer
func (s *service) GetSchedules(page, limit int) []*md.Schedule {
	return s.sRepo.GetAll(page, limit)
}

func (s *service) GetSchedulesByCourse(courseID uint) []*md.Schedule {
	return s.sRepo.GetByCourse(courseID)
}

func (s *service) GetSchedulesByCourseOwner(userID string, courseID uint) []*md.Schedule {
	return s.sRepo.GetByCourseOwner(userID, courseID)
}

// CreateSchedule Creates New schedule
func (s *service) CreateSchedule(c md.Schedule) (*md.Schedule, tr.TParam, error) {

	schedule, err := s.sRepo.Store(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_creation_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return schedule, tParam, err
	}

	return schedule, tr.TParam{}, nil

}

// UpdateSchedule update existing schedule
func (s *service) UpdateSchedule(c *md.Schedule) (*md.Schedule, tr.TParam, error) {
	schedule, err := s.sRepo.Update(c)

	if err != nil {
		tParam := tr.TParam{
			Key:          "error.resource_update_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		return schedule, tParam, err
	}

	return schedule, tr.TParam{}, nil

}

// Validate Function for validating schedule input
func Validate(schedule Payload, r *http.Request) error {
	return validation.ValidateStruct(&schedule,
		validation.Field(&schedule.CourseID, ut.IDRule(r)...),
		validation.Field(&schedule.WeekDay, ut.RequiredRule(r, "general.week_day")...),
		validation.Field(&schedule.TimeFromHour, ut.RequiredRule(r, "general.hour")...),
		validation.Field(&schedule.TimeFromMinute, ut.RequiredRule(r, "general.minute")...),
		validation.Field(&schedule.TimeToHour, ut.RequiredRule(r, "general.hour")...),
		validation.Field(&schedule.TimeToHour, ut.RequiredRule(r, "general.minute")...),
	)
}

// ValidateUpdates Function for validating schedule update input
func ValidateUpdates(schedule Payload, r *http.Request) error {
	return validation.ValidateStruct(&schedule,
		validation.Field(&schedule.WeekDay, ut.RequiredRule(r, "general.week_day")...),
		validation.Field(&schedule.TimeFromHour, ut.RequiredRule(r, "general.hour")...),
		validation.Field(&schedule.TimeFromMinute, ut.RequiredRule(r, "general.minute")...),
		validation.Field(&schedule.TimeToHour, ut.RequiredRule(r, "general.hour")...),
		validation.Field(&schedule.TimeToHour, ut.RequiredRule(r, "general.minute")...),
	)
}
