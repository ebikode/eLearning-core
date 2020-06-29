package main

import (
	adm "github.com/ebikode/eLearning-core/domain/admin"
	ast "github.com/ebikode/eLearning-core/domain/app_setting"
	cou "github.com/ebikode/eLearning-core/domain/course"
	emp "github.com/ebikode/eLearning-core/domain/user"
	jb "github.com/ebikode/eLearning-core/jobs"
	storage "github.com/ebikode/eLearning-core/storage/mysql"
	"github.com/whiteshtef/clockwork"
)

// InitJobs Initialize all scheduled jobs
func InitJobs(mdb *storage.MDatabase) {

	var adminStorage adm.AdminRepository
	var userStorage emp.UserRepository
	var courseStorage cou.CourseRepository
	var appSettingStorage ast.AppSettingRepository

	// initalising all domain storage for db manipulation
	adminStorage = storage.NewDBAdminStorage(mdb)
	courseStorage = storage.NewDBCourseStorage(mdb)
	userStorage = storage.NewDBUserStorage(mdb)
	appSettingStorage = storage.NewDBAppSettingStorage(mdb)

	// Initailizing application domain services
	admService := adm.NewService(adminStorage)
	usService := emp.NewService(userStorage)
	couService := cou.NewService(courseStorage)
	astService := ast.NewService(appSettingStorage)

	// Initialize clockwork schedules
	sched := clockwork.NewScheduler()

	var runJobs = func() {

		// RunCreateDefaultSuperAdmin - Create Default admin  if it doesn't exist
		// the first time the server is launch
		jb.RunCreateDefaultSuperAdmin(admService)

		// Create default app settings
		jb.RunCreateDefaultSettings(astService)

		// Create default users
		jb.RunCreateDefaultUsers(usService, couService)

		sched.Run()
	}

	go runJobs()

}
