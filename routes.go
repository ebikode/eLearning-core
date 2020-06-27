package main

import (
	"fmt"

	"github.com/ebikode/eLearning-core/config"
	alog "github.com/ebikode/eLearning-core/domain/activity_log"
	adm "github.com/ebikode/eLearning-core/domain/admin"
	ast "github.com/ebikode/eLearning-core/domain/app_setting"
	aud "github.com/ebikode/eLearning-core/domain/authd_device"
	slr "github.com/ebikode/eLearning-core/domain/salary"
	emp "github.com/ebikode/eLearning-core/domain/user"

	pyr "github.com/ebikode/eLearning-core/domain/payroll"
	tax "github.com/ebikode/eLearning-core/domain/tax"
	endP "github.com/ebikode/eLearning-core/endpoints"
	mw "github.com/ebikode/eLearning-core/middlewares"
	storage "github.com/ebikode/eLearning-core/storage/mysql"
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

	var userStorage emp.UserRepository
	var adminStorage adm.AdminRepository
	var payrollStorage pyr.PayrollRepository
	var salaryStorage slr.SalaryRepository
	var authdDeviceStorage aud.AuthdDeviceRepository
	var activityLogStorage alog.ActivityLogRepository
	var appSettingStorage ast.AppSettingRepository
	var taxStorage tax.TaxRepository

	// initalising all domain storage for db manipulation
	userStorage = storage.NewDBUserStorage(mdb)
	adminStorage = storage.NewDBAdminStorage(mdb)
	payrollStorage = storage.NewDBPayrollStorage(mdb)
	salaryStorage = storage.NewDBSalaryStorage(mdb)
	authdDeviceStorage = storage.NewDBAuthdDeviceStorage(mdb)
	taxStorage = storage.NewDBTaxStorage(mdb)
	activityLogStorage = storage.NewDBActivityLogStorage(mdb)
	appSettingStorage = storage.NewDBAppSettingStorage(mdb)

	// Initailizinf application domain services
	empService := emp.NewService(userStorage)
	admService := adm.NewService(adminStorage)
	audService := aud.NewService(authdDeviceStorage)
	pyrService := pyr.NewService(payrollStorage)
	taxService := tax.NewService(taxStorage)
	alogService := alog.NewService(activityLogStorage)
	astService := ast.NewService(appSettingStorage)
	salaryService := slr.NewService(salaryStorage)
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

		r.Get("/user/verify/email/{userID}/{emailToken}", endP.VerifyUserEmailEndpoint(empService))

		r.Route("/user", func(r chi.Router) {
			r.Use(
				mw.JwtUserAuthentication(cfg.Auth.UserTokenSecret, cfg.Server.AppKey), // Authentication middleware
			)
			r.Post("/authenticate", endP.AuthenticateUserEndpoint(cfg.Auth.UserTokenSecret,
				empService, pyrService, audService))
			r.Get("/me", endP.GetUserEndpoint(empService))

			// r.Route("/settings", func(r chi.Router) {
			// 	r.Put("/update/billing", endP.UpdateBillingSettingsEndpoint(ustService))
			// })

			// Tax
			r.Route("/tax", func(r chi.Router) {
				r.Get("/", endP.GetUserTaxesEndpoint(taxService, "user"))
			})

			// Payroll Endpoints
			r.Route("/payrolls", func(r chi.Router) {
				r.Get("/", endP.GetUserPayrollsEndpoint(pyrService, "user"))
			})
		})

		// ADMIN ROUTES - All Admin have access
		r.Route("/admin", func(r chi.Router) {
			r.Use(
				mw.JwtAdminAuthentication(cfg.Auth.AdminTokenSecret), // Authentication middleware
			)

			r.Post("/authenticate", endP.AuthenticateAdminEndpoint(cfg.Auth.AdminTokenSecret, admService, pyrService))

			// Get an admin
			r.Get("/me", endP.GetAdminEndpoint(admService))

			// Endpoints for Users Access
			r.Route("/users", func(r chi.Router) {
				r.Get("/", endP.GetUsersEndpoint(empService))
			})

			// Salary Endpoint
			r.Route("/salary", func(r chi.Router) {
				r.Get("/", endP.GetAdminSalariesEndpoint(salaryService))
				r.Get("/{salaryID}", endP.GetSalaryEndpoint(salaryService))
			})

			// Tax Enpoints
			r.Route("/taxes", func(r chi.Router) {
				r.Get("/", endP.GetAdminTaxesEndpoint(taxService))
				r.Get("/{taxID}", endP.GetTaxEndpoint(taxService))
				r.Get("/user/{userID}", endP.GetUserTaxesEndpoint(taxService, "admin"))
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

				r.Route("/user", func(r chi.Router) {
					r.Post("/create", endP.CreateUserEndpoint(empService, alogService, clientURL, sendGridKey))
					r.Put("/update/{userID}", endP.UpdateUserEndpoint(empService, alogService))
				})

				// Payrolls Endpoints
				r.Route("/payrolls", func(r chi.Router) {
					r.Get("/", endP.GetPayrollsEndpoint(pyrService, "admin"))
					r.Get("/reports", endP.GetPayrollReportsEndpoint(pyrService))
					r.Get("/filters", endP.GetPayrollAllMonthAndYearEndpoint(pyrService))
					r.Get("/by_month/{month}/{year}", endP.GetPayrollsByMonthYearEndpoint(pyrService, "admin"))
					r.Get("/single/{userID}/{payrollID}", endP.GetPayrollEndpoint(pyrService, "admin"))
					r.Get("/user/{userID}", endP.GetUserPayrollsEndpoint(pyrService, "admin"))
					r.Put("/update/status", endP.UpdatePayrollStatusEndpoint(pyrService, alogService))
					r.Put("/update/payment_status", endP.UpdatePayrollPaymentStatusEndpoint(pyrService, alogService))
				})

				// Salary Endpoint
				r.Route("/salary", func(r chi.Router) {
					r.Post("/", endP.CreateSalaryEndpoint(salaryService, empService, alogService))
					r.Put("/{salaryID}", endP.UpdateSalaryEndpoint(salaryService, alogService))
				})
			})

		})

	})

	return router
}
