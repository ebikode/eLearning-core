package main

import (
	"fmt"

	"github.com/ebikode/eLearning-core/config"
	alog "github.com/ebikode/eLearning-core/domain/activity_log"
	adm "github.com/ebikode/eLearning-core/domain/admin"
	ast "github.com/ebikode/eLearning-core/domain/app_setting"
	app "github.com/ebikode/eLearning-core/domain/application"
	art "github.com/ebikode/eLearning-core/domain/article"
	asm "github.com/ebikode/eLearning-core/domain/assessment"
	aud "github.com/ebikode/eLearning-core/domain/authd_device"
	cou "github.com/ebikode/eLearning-core/domain/course"
	jon "github.com/ebikode/eLearning-core/domain/journal"
	que "github.com/ebikode/eLearning-core/domain/question"
	shd "github.com/ebikode/eLearning-core/domain/schedule"
	us "github.com/ebikode/eLearning-core/domain/user"

	grd "github.com/ebikode/eLearning-core/domain/grade"
	endP "github.com/ebikode/eLearning-core/endpoints"
	mw "github.com/ebikode/eLearning-core/middlewares"
	storage "github.com/ebikode/eLearning-core/storage/mysql"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// InitRoutes Initialize all routes
func InitRoutes(cfg config.Constants, mdb *storage.MDatabase) *chi.Mux {
	baseURL := cfg.Server.URL
	clientURL := cfg.Client.URL
	sendGridKey := cfg.SendGrid.ApiKey

	fmt.Println(baseURL)
	// fmt.Println(sendGridKey)
	// fmt.Println(cfg.Server.AppKey)
	// fmt.Println(cfg.Pexportal.BaseURL)
	// fmt.Println(cfg.Auth.AccountUserTokenSecret)

	var userStorage us.UserRepository
	var adminStorage adm.AdminRepository
	var gradeStorage grd.GradeRepository
	var articleStorage art.ArticleRepository
	var journalStorage jon.JournalRepository
	var applicationStorage app.ApplicationRepository
	var courseStorage cou.CourseRepository
	var questionStorage que.QuestionRepository
	var assessmentStorage asm.AssessmentRepository
	var scheduleStorage shd.ScheduleRepository
	var authdDeviceStorage aud.AuthdDeviceRepository
	var activityLogStorage alog.ActivityLogRepository
	var appSettingStorage ast.AppSettingRepository

	// initalising all domain storage for db manipulation
	userStorage = storage.NewDBUserStorage(mdb)
	adminStorage = storage.NewDBAdminStorage(mdb)
	gradeStorage = storage.NewDBGradeStorage(mdb)
	articleStorage = storage.NewDBArticleStorage(mdb)
	journalStorage = storage.NewDBJournalStorage(mdb)
	assessmentStorage = storage.NewDBAssessmentStorage(mdb)
	courseStorage = storage.NewDBCourseStorage(mdb)
	questionStorage = storage.NewDBQuestionStorage(mdb)
	applicationStorage = storage.NewDBApplicationStorage(mdb)
	authdDeviceStorage = storage.NewDBAuthdDeviceStorage(mdb)
	scheduleStorage = storage.NewDBScheduleStorage(mdb)
	activityLogStorage = storage.NewDBActivityLogStorage(mdb)
	appSettingStorage = storage.NewDBAppSettingStorage(mdb)

	// Initailizing application domain services
	usService := us.NewService(userStorage)
	admService := adm.NewService(adminStorage)
	audService := aud.NewService(authdDeviceStorage)
	grdService := grd.NewService(gradeStorage)
	jonService := jon.NewService(journalStorage)
	couService := cou.NewService(courseStorage)
	queService := que.NewService(questionStorage)
	appService := app.NewService(applicationStorage)
	shdService := shd.NewService(scheduleStorage)
	asmService := asm.NewService(assessmentStorage)
	alogService := alog.NewService(activityLogStorage)
	astService := ast.NewService(appSettingStorage)
	articleService := art.NewService(articleStorage)

	// ustService := ust.NewService(userSettingStorage)
	// Initialize router
	router := chi.NewRouter()

	// Add middlewares to router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger, // Log API request calls
		//middleware.Compress,        // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/api/v1", func(r chi.Router) {

		// USER ROUTES

		r.Get("/user/verify/email/{userID}/{emailToken}", endP.VerifyUserEmailEndpoint(usService))

		r.Route("/user", func(r chi.Router) {
			r.Use(
				mw.JwtUserAuthentication(cfg.Auth.UserTokenSecret, cfg.Server.AppKey), // Authentication middleware
			)

			r.Post("/create", endP.CreateUserEndpoint(usService, clientURL, sendGridKey))
			r.Post("/authenticate", endP.AuthenticateUserEndpoint(cfg.Auth.UserTokenSecret,
				usService, appService, audService))
			r.Put("/update/{userID}", endP.UpdateUserEndpoint(usService, alogService))
			r.Get("/me", endP.GetUserEndpoint(usService))

			// Grade Endpoints
			r.Route("/grades", func(r chi.Router) {
				r.Get("/", endP.GetUserGradesEndpoint(grdService, ut.UserRole))
			})

			// Applications Endpoints
			r.Route("/applications", func(r chi.Router) {
				r.Get("/", endP.GetUserApplicationsEndpoint(appService, ut.UserRole))
				r.Post("/{applicationID}", endP.CreateApplicationEndpoint(appService, usService, couService))
			})

			// Questions Endpoints
			r.Route("/questions", func(r chi.Router) {
				r.Get("/course/{courseID}", endP.GetCourseQuestionsEndpoint(queService))
			})

			// Courses Endpoint
			r.Route("/courses", func(r chi.Router) {
				r.Get("/", endP.GetCoursesEndpoint(couService, ut.UserRole))
			})

			// Schedules Endpoint
			r.Route("/schedules", func(r chi.Router) {
				r.Get("/course/{courseID}", endP.GetCourseSchedulesEndpoint(shdService))
			})

			// Article Endpoint
			r.Route("/article", func(r chi.Router) {
				r.Get("/{articleID}", endP.GetArticleEndpoint(articleService))
				r.Get("/course/{courseID}", endP.GetCourseArticlesEndpoint(articleService))
			})

			// Journal Endpoint
			r.Route("/journals", func(r chi.Router) {
				r.Get("/{journalID}", endP.GetJournalEndpoint(jonService))
				r.Get("/course/{courseID}", endP.GetCourseJournalsEndpoint(jonService))
			})

			// Assessment Endpoint
			r.Route("/assessments", func(r chi.Router) {
				r.Get("/{applicationID}", endP.GetUserApplicationAssessmentsEndpoint(asmService, ut.UserRole))
				r.Post("/submit/{applicationID}", endP.CreateAssessmentEndpoint(asmService, appService, grdService, usService, queService))
			})

			// TUTOR Endpoints - Only Super admin access
			r.Route("/tutor", func(r chi.Router) {
				r.Use(
					mw.IsTutorUser(), // Super middleware
				)

				// Applications Endpoints
				r.Route("/applications", func(r chi.Router) {
					r.Get("/", endP.GetCourseOwnerApplicationsEndpoint(appService, ut.TutorRole))
					r.Get("/course/{courseID}", endP.GetCourseApplicationsEndpoint(appService))
				})

				// Article Endpoint
				r.Route("/article", func(r chi.Router) {
					r.Get("/", endP.GetUserArticlesEndpoint(articleService, ut.UserRole))
					r.Post("/", endP.CreateArticleEndpoint(articleService, usService))
					r.Put("/{articleID}", endP.UpdateArticleEndpoint(articleService))
				})

				// Journal Endpoint
				r.Route("/journals", func(r chi.Router) {
					r.Get("/", endP.GetUserJournalsEndpoint(jonService, ut.UserRole))
					r.Post("/", endP.CreateJournalEndpoint(jonService, usService))
					r.Put("/{journalID}", endP.UpdateJournalEndpoint(jonService))
				})

				// Questions Endpoints
				r.Route("/questions", func(r chi.Router) {
					r.Get("/course/{courseID}", endP.GetTutorCourseQuestionsEndpoint(queService))
					r.Post("/", endP.CreateQuestionEndpoint(queService, usService))
					r.Put("/{questionID}", endP.UpdateQuestionEndpoint(queService))
				})

				// Courses Endpoint
				r.Route("/courses", func(r chi.Router) {
					r.Get("/", endP.GetUserCoursesEndpoint(couService))
					r.Post("/", endP.CreateCourseEndpoint(couService, usService))
					r.Put("/{courseID}", endP.UpdateCourseEndpoint(couService))
				})

				// Schedules Endpoint
				r.Route("/schedules", func(r chi.Router) {
					r.Get("/", endP.GetAdminSchedulesEndpoint(shdService))
					r.Get("/course/{courseID}", endP.GetTutorCourseSchedulesEndpoint(shdService))
					r.Post("/", endP.CreateScheduleEndpoint(shdService, usService))
					r.Put("/{scheduleID}", endP.UpdateScheduleEndpoint(shdService))
				})

				// Assessment Endpoint
				r.Route("/assessments", func(r chi.Router) {
					r.Get("/{userID}/{applicationID}", endP.GetUserApplicationAssessmentsEndpoint(asmService, ut.AdminRole))
				})

				// Grade Endpoints
				r.Route("/grades", func(r chi.Router) {
					r.Get("/{courseID}", endP.GetCourseGradesEndpoint(grdService, ut.TutorRole))
				})
			})
		})

		// ADMIN ROUTES - All Admin have access
		r.Route("/admin", func(r chi.Router) {
			r.Use(
				mw.JwtAdminAuthentication(cfg.Auth.AdminTokenSecret), // Authentication middleware
			)

			r.Post("/authenticate", endP.AuthenticateAdminEndpoint(cfg.Auth.AdminTokenSecret, admService, appService))

			// Get an admin
			r.Get("/me", endP.GetAdminEndpoint(admService))

			// Endpoints for Users Access
			r.Route("/users", func(r chi.Router) {
				r.Get("/", endP.GetUsersEndpoint(usService))
			})

			// Applications Endpoints
			r.Route("/applications", func(r chi.Router) {
				r.Get("/", endP.GetAdminApplicationsEndpoint(appService))
				r.Get("/course/{courseID}", endP.GetCourseApplicationsEndpoint(appService))
			})

			// Courses Endpoint
			r.Route("/courses", func(r chi.Router) {
				r.Get("/", endP.GetCoursesEndpoint(couService, ut.AdminRole))
			})

			// Article Endpoint
			r.Route("/articles", func(r chi.Router) {
				r.Get("/", endP.GetAdminArticlesEndpoint(articleService))
				r.Get("/{articleID}", endP.GetArticleEndpoint(articleService))
				r.Get("/course/{courseID}", endP.GetCourseArticlesEndpoint(articleService))
			})

			// Journal Endpoint
			r.Route("/journals", func(r chi.Router) {
				r.Get("/", endP.GetAdminJournalsEndpoint(jonService))
				r.Get("/{articleID}", endP.GetJournalEndpoint(jonService))
				r.Get("/course/{courseID}", endP.GetCourseJournalsEndpoint(jonService))
			})

			// Grades Endpoints
			r.Route("/grades", func(r chi.Router) {
				r.Get("/", endP.GetGradesEndpoint(grdService, ut.AdminRole))
				r.Get("/reports", endP.GetGradeReportsEndpoint(grdService))
				r.Get("/user/{userID}", endP.GetUserGradesEndpoint(grdService, ut.AdminRole))
				r.Get("/course/{userID}/{courseID}", endP.GetCourseGradesEndpoint(grdService, ut.AdminRole))
			})

			// General Admin App Settings Endpoints
			r.Route("/app_settings", func(r chi.Router) {
				r.Get("/", endP.GetAppSettingsEndpoint(astService, "admin"))
				r.Get("/key/{sKEY}", endP.GetAppSettingByKeyEndpoint(astService, "admin"))
				r.Get("/{appsettingID}", endP.GetAppSettingEndpoint(astService))
			})

			// Super Admin Endpoints - Only Super admin access
			r.Route("/super_admin", func(r chi.Router) {
				r.Use(
					mw.IsSuperAdmin(), // Super middleware
				)
				// Admin account endpoints
				r.Route("/account", func(r chi.Router) {
					r.Post("/create", endP.CreateAdminEndpoint(admService))
				})

				// Activity logs endpoints
				r.Route("/activity_log", func(r chi.Router) {
					r.Get("/", endP.GetActivityLogsEndpoint(alogService))
				})

				// App Settings endpoints
				r.Route("/app_settings", func(r chi.Router) {
					r.Post("/", endP.CreateAppSettingEndpoint(astService, alogService))
					r.Put("/{appsettingID}", endP.UpdateAppSettingEndpoint(astService, alogService))
				})
			})

			// Manager Admin Endpoints - Only sales admin and super admin accesss
			r.Route("/manager", func(r chi.Router) {
				r.Use(
					mw.IsManagerAdmin(), //Sales Admin middleware
				)

				// Courses Endpoints
				r.Route("/courses", func(r chi.Router) {
					r.Put("/status/{courseID}", endP.UpdateCourseStatusEndpoint(couService))
				})

				// Applications Endpoints
				r.Route("/applications", func(r chi.Router) {
					r.Put("/issue_certificate/{applicationID}", endP.IssueCertificateEndpoint(appService))
				})
			})

		})

	})

	return router
}
