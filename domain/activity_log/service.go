package activitylog

import (
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

// Service provides activityLog operations
type ActivityLogService interface {
	GetActivityLogs(int, int) []*md.ActivityLog
	CreateActivityLog(md.ActivityLog) error
}

type service struct {
	alRepo ActivityLogRepository
}

// NewService creates a activityLog service with the necessary dependencies
func NewService(
	alRepo ActivityLogRepository,
) ActivityLogService {
	return &service{alRepo}
}

// Get a activityLog
func (s *service) GetActivityLogs(page, limit int) []*md.ActivityLog {
	return s.alRepo.GetAll(page, limit)
}

// Create New activityLog
func (s *service) CreateActivityLog(c md.ActivityLog) error {
	aID := ut.RandomBase64String(8, "MDlg")
	c.ID = aID
	err := s.alRepo.Store(c)

	if err != nil {
		return err
	}

	return nil

}
