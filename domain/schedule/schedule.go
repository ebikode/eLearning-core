package schedule

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Payload ...
type Payload struct {
	CourseID       uint   `json:"course_id"`
	WeekDay        string `json:"week_day"`
	TimeFromHour   string `json:"time_from_hour"`
	TimeFromMinute string `json:"time_from_minute"`
	TimeToHour     string `json:"time_to_hour"`
	TimeToMunite   string `json:"time_to_minute"`
	Status         string `json:"status"`
}

// ValidationFields struct to return for validation
type ValidationFields struct {
	CourseID       string `json:"course_id"`
	WeekDay        string `json:"week_day"`
	TimeFromHour   string `json:"time_from_hour"`
	TimeFromMinute string `json:"time_from_minute"`
	TimeToHour     string `json:"time_to_hour"`
	TimeToMunite   string `json:"time_to_minute"`
	Status         string `json:"status"`
}

// ScheduleRepository  provides access to the md.Schedule storage.
type ScheduleRepository interface {
	// Get returns the schedule with given ID.
	Get(uint) *md.Schedule
	GetByCourse(uint) []*md.Schedule
	GetByCourseOwner(string, uint) []*md.Schedule
	// Get returns all schedules.
	GetAll(int, int) []*md.Schedule
	// Store a given schedule to the repository.
	Store(md.Schedule) (*md.Schedule, error)
	// Update a given schedule in the repository.
	Update(*md.Schedule) (*md.Schedule, error)
	// Delete a given schedule in the repository.
	Delete(md.Schedule, bool) (bool, error)
}
