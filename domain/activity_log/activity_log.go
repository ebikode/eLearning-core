package activitylog

import (
	md "github.com/ebikode/eLearning-core/model"
)

// Repository provides access to the ActivityLog storage.
type ActivityLogRepository interface {
	// Get returns the activity log with given ID.
	GetAll(int, int) []*md.ActivityLog
	// Store a given customer activity log to the repository.
	Store(md.ActivityLog) error
}
