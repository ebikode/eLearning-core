package jobs

import (
	"fmt"

	adm "github.com/ebikode/eLearning-core/domain/admin"
	apset "github.com/ebikode/eLearning-core/domain/app_setting"
	cou "github.com/ebikode/eLearning-core/domain/course"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

// RunCreateDefaultSuperAdmin - Create Default admin  if it doesn't exist
// the first time the server is launch
func RunCreateDefaultSuperAdmin(adms adm.AdminService) {

	isDefaultAdminCreated := adms.CheckAdminCreated()

	if !isDefaultAdminCreated {
		// password := ut.RandomBase64String(10, "")
		password := "EL@DM1N2020"
		admin := md.Admin{
			Phone:     "2347067413685",
			FirstName: "Super",
			LastName:  "Admin",
			Email:     "superadmin@elearning-demo.com",
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
			Name:    "Maximum Number of Questions",
			SKey:    ut.MaxNumOfQuestions,
			Value:   "50",
			Comment: "Check the maximum number of assessment questions a tutor can add to a course",
		},
		{
			Name:    "Maximum Number of Course",
			SKey:    ut.MaxNumOfCourses,
			Value:   "100",
			Comment: "Check maximum number of courses a tutor can create",
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
func RunCreateDefaultUsers(us usr.UserService, cos cou.CourseService) {
	password := "ELPASSWORD2020"
	users := []md.User{
		{
			FirstName:       "Ada",
			LastName:        "Musa",
			Username:        "adamusa",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "adamusa@elearning-demo.com",
			Password:        password,
			Phone:           "08012345678",
			IsEmailVerified: true,
			Status:          ut.Active,
			Role:            ut.UserRole,
		},
		{
			FirstName:       "Adesuwa",
			LastName:        "Habib",
			Username:        "adesuwa",
			Address:         "21 Edo Nigeria",
			About:           "Brilliant and hard working",
			Email:           "adesuwa@elearning-demo.com",
			Password:        password,
			Phone:           "08012345677",
			IsEmailVerified: true,
			Status:          ut.Active,
			Role:            ut.TutorRole,
		},
		{
			FirstName:       "Aisha",
			LastName:        "Emeka",
			Username:        "aishaemeka",
			Address:         "21 Kano Nigeria",
			About:           "Brilliant and hard working",
			Email:           "aishaemeka@elearning-demo.com",
			Password:        password,
			Phone:           "08012345676",
			IsEmailVerified: true,
			Status:          ut.Active,
			Role:            ut.TutorRole,
		},
		{
			FirstName:       "Bello",
			LastName:        "Gigabit",
			Username:        "bellogig",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "bellogig@elearning-demo.com",
			Password:        password,
			Phone:           "08012345667",
			IsEmailVerified: true,
			Status:          ut.Active,
			Role:            ut.UserRole,
		},
		{
			FirstName:       "Osaro",
			LastName:        "Megabit",
			Username:        "osaromegabit",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "osaromegabit@elearning-demo.com",
			Password:        password,
			Phone:           "08012345675",
			IsEmailVerified: true,
			Status:          ut.Active,
			Role:            ut.UserRole,
		},
		{
			FirstName:       "Joromi",
			LastName:        "Doe",
			Username:        "joromidoe",
			Address:         "21 Lagos Nigeria",
			About:           "Brilliant and hard working",
			Email:           "joromidoe@elearning-demo.com",
			Password:        password,
			Phone:           "08012345674",
			IsEmailVerified: true,
			Status:          ut.Active,
			Role:            ut.UserRole,
		},
	}

	courses := []md.Course{
		{
			Title:               "Web Development",
			Description:         "Focused on Teaching you web development using HTML/CSS and JavaScript",
			DurationPerQuestion: 30,
			Status:              ut.Approved,
		},
		{
			Title:               "Building Micro-Services with Golang",
			Description:         "Focused on Teaching you micro services programming and architecture",
			DurationPerQuestion: 30,
			Status:              ut.Approved,
		},
	}

	for _, v := range users {

		user := us.GetUserByEmail(v.Email)

		if user == nil {
			user, _, _ := us.CreateUser(v)

			if user.Role == ut.TutorRole {
				for _, cv := range courses {
					cv.UserID = user.ID
					cos.CreateCourse(cv)
				}
			}
		}
	}

}
