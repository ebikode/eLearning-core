package endpoints

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	adm "github.com/ebikode/eLearning-core/domain/admin"
	app "github.com/ebikode/eLearning-core/domain/application"
	aud "github.com/ebikode/eLearning-core/domain/authd_device"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
)

// AuthenticateUserEndpoint Authenticate a user
func AuthenticateUserEndpoint(appSecret string, us usr.UserService, aps app.ApplicationService, au aud.AuthdDeviceService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		payload := usr.Payload{}

		err := json.NewDecoder(r.Body).Decode(&payload)

		tParam := tr.TParam{
			Key:          "error.request_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		if err != nil {
			// Respond with an errortra nslated

			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		user := &md.PubUser{}

		user, tParam, err = us.AuthenticateUser(payload.Email, payload.Password)

		if err != nil {
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusUnauthorized, w, r, resp)
			return
		}

		// Detect and save the device the user used in loggin in
		deviceInfo := ut.DetectDevice(r)
		ip := deviceInfo.IP.To4().String()

		aDevice := md.AuthdDevice{
			UserID:         user.ID,
			IP:             ip,
			Browser:        deviceInfo.Browser,
			BrowserVersion: deviceInfo.BrowserVersion,
			Platform:       deviceInfo.Platform,
			DeviceOS:       deviceInfo.DeviceOS,
			OSVersion:      deviceInfo.OSVersion,
			DeviceType:     deviceInfo.Type,
			AccessType:     "web_app",
			Status:         "active",
		}

		audID := ut.RandomBase64String(8, "MDdv")

		aDevice.ID = audID

		audd, _, _ := au.CreateAuthdDevice(aDevice)

		resp := ut.Message(true, ut.Translate(tParam, r))

		applications := []*md.Application{}
		userDashoardData := &md.UserDashbordData{}
		tutorDashBoardData := &md.TutorDashbordData{}

		if user.Role == ut.TutorRole {
			tutorDashBoardData = us.GetTutorDashbordData(user.ID)
			resp["dashboard_data"] = tutorDashBoardData
			applications = aps.GetApplicationsByCourseOwner(user.ID)
		}
		if user.Role == ut.UserRole {
			userDashoardData = us.GetUserDashboardData(user.ID)
			resp["dashboard_data"] = userDashoardData
			applications = aps.GetUserApplications(user.ID)
		}

		// Create JWT token for application
		tk := &md.UserTokenData{
			UserID:   user.ID,
			DeviceID: audd.ID,
			Username: user.Username,
			Role:     user.Role,
			ExpireOn: time.Now().Add(time.Duration(31536000)).UTC(),
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(appSecret))

		tParam = tr.TParam{
			Key:          "general.login_success",
			TemplateData: nil,
			PluralCount:  nil,
		}
		resp["token"] = tokenString
		resp["user"] = user
		resp["recent_applications"] = applications
		ut.Respond(w, r, resp)
	}
}

// AuthenticateAdminEndpoint - Authenticate admin
func AuthenticateAdminEndpoint(appSecret string, as adm.AdminService, aps app.ApplicationService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		admin := &md.Admin{}

		err := json.NewDecoder(r.Body).Decode(admin)

		tParam := tr.TParam{
			Key:          "error.request_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		if err != nil {
			// Respond with an errortra nslated

			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		admin, tParam, err = as.AuthenticateAdmin(admin.Email, admin.Password)

		if err != nil {
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusUnauthorized, w, r, resp)
			return
		}
		applications := []*md.Application{}

		dashoardData := as.GetAdminDashboardData()

		applications = aps.GetApplications(1, 20)

		//Create JWT token
		tk := &md.AdminTokenData{
			AdminID:  admin.ID,
			Role:     admin.Role,
			ExpireOn: time.Now().Add(time.Duration(86400)).UTC(),
		}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(appSecret))

		tParam = tr.TParam{
			Key:          "general.admin_login_success",
			TemplateData: nil,
			PluralCount:  nil,
		}
		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["token"] = tokenString
		resp["admin"] = admin
		resp["dashboard_data"] = dashoardData
		resp["recent_applications"] = applications
		ut.Respond(w, r, resp)
	}

}
