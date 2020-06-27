package main

import (
	adm "github.com/ebikode/eLearning-core/domain/admin"
	ast "github.com/ebikode/eLearning-core/domain/app_setting"
	pyr "github.com/ebikode/eLearning-core/domain/payroll"
	slr "github.com/ebikode/eLearning-core/domain/salary"
	tax "github.com/ebikode/eLearning-core/domain/tax"
	emp "github.com/ebikode/eLearning-core/domain/user"
	jb "github.com/ebikode/eLearning-core/jobs"
	storage "github.com/ebikode/eLearning-core/storage/mysql"
	"github.com/whiteshtef/clockwork"
)

// InitJobs Initialize all scheduled jobs
func InitJobs(mdb *storage.MDatabase) {

	var adminStorage adm.AdminRepository
	var userStorage emp.UserRepository
	var payrollStorage pyr.PayrollRepository
	var taxStorage tax.TaxRepository
	var salaryStorage slr.SalaryRepository
	var appSettingStorage ast.AppSettingRepository

	// initalising all domain storage for db manipulation
	adminStorage = storage.NewDBAdminStorage(mdb)
	payrollStorage = storage.NewDBPayrollStorage(mdb)
	userStorage = storage.NewDBUserStorage(mdb)
	salaryStorage = storage.NewDBSalaryStorage(mdb)
	taxStorage = storage.NewDBTaxStorage(mdb)
	appSettingStorage = storage.NewDBAppSettingStorage(mdb)

	// Initailizing application domain services
	admService := adm.NewService(adminStorage)
	empService := emp.NewService(userStorage)
	pyrService := pyr.NewService(payrollStorage)
	taxService := tax.NewService(taxStorage)
	salaryService := slr.NewService(salaryStorage)
	astService := ast.NewService(appSettingStorage)

	// Initialize clockwork schedules
	sched := clockwork.NewScheduler()

	// go runJobs(leaguesURL, fixtureURL, ls, cs, ss, fs, ts)
	var runJobs = func() {

		// RunCreateDefaultSuperAdmin - Create Default admin  if it doesn't exist
		// the first time the server is launch
		jb.RunCreateDefaultSuperAdmin(admService)

		// Create default app settings
		jb.RunCreateDefaultSettings(astService)

		// Create default users
		jb.RunCreateDefaultUsers(empService, salaryService)

		// Create Payroll Demo Data
		jb.RunDefaultPayrollGenerationJob(pyrService, empService, taxService)

		var runGeneratePayrollJob = func() {
			jb.RunPayrollGenerationJob(pyrService, astService, empService, taxService)
		}

		var runPayrollPaymentJob = func() {
			jb.RunPayrollPaymentJob(pyrService, astService)
		}

		runGeneratePayrollJob()
		runPayrollPaymentJob()

		// This runs every 12 Hours
		go sched.Schedule().Every(12).Hours().Do(runGeneratePayrollJob)

		// This runs every 60 Minutes
		go sched.Schedule().Every(60).Minutes().Do(runPayrollPaymentJob)

		sched.Run()
	}

	go runJobs()

}
