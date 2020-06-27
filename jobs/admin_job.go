package jobs

import (
	"fmt"

	adm "github.com/ebikode/eLearning-core/domain/admin"
	apset "github.com/ebikode/eLearning-core/domain/app_setting"
	slr "github.com/ebikode/eLearning-core/domain/salary"
	emp "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

// RunCreateDefaultSuperAdmin - Create Default admin  if it doesn't exist
// the first time the server is launch
func RunCreateDefaultSuperAdmin(adms adm.AdminService) {

	isDefaultAdminCreated := adms.CheckAdminCreated()

	if !isDefaultAdminCreated {
		// password := ut.RandomBase64String(10, "")
		password := "PR@DM1N2020"
		admin := md.Admin{
			Phone:     "2347067413685",
			FirstName: "Super",
			LastName:  "Admin",
			Email:     "superadmin@payroll-demo.com",
			Password:  password,
			Role:      "super_admin",
		}

		fmt.Println("Admin Password:: ", password)
		fmt.Println("Admin Email:: ", admin.Email)

		adms.CreateAdmin(admin)
	}
}

// RunCreateDefaultSettings - Create Default App Settings  if it doesn't exist
// the first time the server is launch
func RunCreateDefaultSettings(aps apset.AppSettingService) {

	settings := []md.AppSetting{
		{
			Name:    "Payroll Generation Day",
			SKey:    ut.PayrollGenerationDayKey,
			Value:   "23",
			Comment: "Used for Automation. Payroll is auto-generated on this day.",
		},
		{
			Name:    "Pay Day",
			SKey:    ut.PayDayKey,
			Value:   "28",
			Comment: "Used for Automation. Payment is processed on this day if approval is automated.",
		},
		{
			Name:    "Auto Approve?",
			SKey:    ut.IsAutoApproveKey,
			Value:   "false",
			Comment: "Used for for Approval and Payment Automation.",
		},
		{
			Name:    "Auto Approval Note",
			SKey:    ut.AutoApprovalNoteKey,
			Value:   "This payroll was automatically approved by the system.",
			Comment: "Used for for Approval and Payment Automation Note.",
		},
	}

	for _, v := range settings {
		setting := aps.GetAppSettingByKey(v.SKey, "admin")

		if setting == nil {
			aps.CreateAppSetting(v)
		}
	}

}

// RunCreateDefaultUsers - Create Default App Settings  if it doesn't exist
// the first time the server is launch
func RunCreateDefaultUsers(emps emp.UserService, sls slr.SalaryService) {
	password := "EMPASSWORD2020"
	users := []md.User{
		{
			FirstName:       "Ada",
			LastName:        "Musa",
			Username:        "adamusa",
			Position:        "Marketing Officer",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "adamusa@payroll-demo.com",
			Password:        password,
			Phone:           "08012345678",
			AccountName:     "Ada Musa",
			AccountNumber:   "01234534789",
			BankName:        "Doe Bank",
			IsEmailVerified: true,
			Status:          "active",
		},
		{
			FirstName:       "Adesuwa",
			LastName:        "Habib",
			Username:        "adesuwa",
			Position:        "Software Developer",
			Address:         "21 Edo Nigeria",
			About:           "Brilliant and hard working",
			Email:           "adesuwa@payroll-demo.com",
			Password:        password,
			Phone:           "08012345677",
			AccountName:     "Adesuwa Habib",
			AccountNumber:   "0123456788",
			BankName:        "Doe Bank",
			IsEmailVerified: true,
			Status:          "active",
		},
		{
			FirstName:       "Aisha",
			LastName:        "Emeka",
			Username:        "aishaemeka",
			Position:        "Human Resource Personnel",
			Address:         "21 Kano Nigeria",
			About:           "Brilliant and hard working",
			Email:           "aishaemeka@payroll-demo.com",
			Password:        password,
			Phone:           "08012345676",
			AccountName:     "Aisha Emeka",
			AccountNumber:   "0123456789",
			BankName:        "Doe Bank",
			IsEmailVerified: true,
			Status:          "active",
		},
		{
			FirstName:       "Bello",
			LastName:        "Gigabit",
			Username:        "bellogig",
			Position:        "Software Developer",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "bellogig@payroll-demo.com",
			Password:        password,
			Phone:           "08012345667",
			AccountName:     "Bello Gigabit",
			AccountNumber:   "0123456785",
			BankName:        "Doe Bank",
			IsEmailVerified: true,
			Status:          "active",
		},
		{
			FirstName:       "Osaro",
			LastName:        "Megabit",
			Username:        "osaromegabit",
			Position:        "Accountant",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "osaromegabit@payroll-demo.com",
			Password:        password,
			Phone:           "08012345675",
			AccountName:     "Osaro Megabit",
			AccountNumber:   "0123456784",
			BankName:        "Doe Bank",
			IsEmailVerified: true,
			Status:          "active",
		},
		{
			FirstName:       "Joromi",
			LastName:        "Doe",
			Username:        "joromidoe",
			Position:        "Content Creator",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "joromidoe@payroll-demo.com",
			Password:        password,
			Phone:           "08012345674",
			AccountName:     "Joromi Doe",
			AccountNumber:   "0123456783",
			BankName:        "Doe Bank",
			IsEmailVerified: true,
			Status:          "active",
		},
	}

	for _, v := range users {
		salary := md.Salary{
			Salary:  250000.00,
			Pension: 5.0,
			Paye:    7.0,
			Nsitf:   2.0,
			Nhf:     2.0,
			Itf:     2.0,
		}
		user := emps.GetUserByEmail(v.Email)

		if user == nil {
			user, _, _ := emps.CreateUser(v)
			salary.UserID = user.ID
			if user.Position == "Software Developer" {
				salary.Salary = 550000.00
			}
			sls.CreateSalary(salary)
		}
	}

}
