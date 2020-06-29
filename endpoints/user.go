package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	act "github.com/ebikode/eLearning-core/domain/activity_log"
	us "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetUserEndpoint fetch Authenticated user account
func GetUserEndpoint(uss us.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)

		user := uss.GetUser(userID)
		resp := ut.Message(true, "")
		resp["user"] = user
		ut.Respond(w, r, resp)
	}
}

// GetUsersEndpoint Admin Endpoint for getting users
func GetUsersEndpoint(uss us.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		users := uss.GetAllUsers(page, limit)

		var nextPage int
		if len(users) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["users"] = users
		ut.Respond(w, r, resp)
	}
}

// CreateUserEndpoint An endpoint for creating users new account
func CreateUserEndpoint(uss us.UserService, clientURL, sendGridKey string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		payload := us.Payload{}

		err := json.NewDecoder(r.Body).Decode(&payload)

		tParam := tr.TParam{
			Key:          "error.request_error",
			TemplateData: nil,
			PluralCount:  nil,
		}

		if err != nil {
			// Respond with an error translated

			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		checkUser := uss.GetUserByEmail(payload.Email)

		// if user already exist send pincode to the user for verification/Authentication
		if checkUser != nil {

			tParam = tr.TParam{
				Key:          "error.email_already_exist",
				TemplateData: nil,
				PluralCount:  nil,
			}

			resp := ut.Message(true, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		user := md.User{}

		// username := strings(user.FirstName).LowerCase() + "_" + strings(user.LastName).LowerCase()

		user.FirstName = payload.FirstName
		user.LastName = payload.LastName
		user.Phone = payload.Phone
		user.Username = payload.Username
		user.Email = payload.Email
		user.Password = payload.Password
		user.Avatar = payload.Avatar
		user.Thumb = payload.Thumb
		user.Role = payload.Role
		user.IsEmailVerified = true
		user.Status = ut.Active
		// user.Username = username

		// Validate user input
		err = us.Validate(user, r)
		if err != nil {
			validationFields := us.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Create a user
		newUser, errParam, err := uss.CreateUser(user)
		if err != nil {
			fmt.Println("Error:: err")
			cErr := ut.CheckUniqueError(r, err)
			if cErr != nil {
				ut.ErrorRespond(http.StatusBadRequest, w, r, ut.Message(false, cErr.Error()))
				return
			}
			ut.ErrorRespond(http.StatusBadRequest, w, r, ut.Message(false, ut.Translate(errParam, r)))
			return
		}
		tParam = tr.TParam{
			Key:          "general.registration_success",
			TemplateData: nil,
			PluralCount:  nil,
		}
		// fmt.Println(sendGridKey)
		vURL := fmt.Sprintf("%s/verify-email/%s/%s", clientURL, newUser.ID, newUser.EmailToken)

		// userName := newUser.FirstName //fmt.Sprintf("%s %s", newUser.FirstName, newUser.LastName)
		// Set up Email Data
		// emailText := "Thank you for being part of our Team. Please click the link below to confirm your email address and view your account"
		// emailData := ut.EmailData{
		// 	To: []*mail.Email{
		// 		mail.NewEmail(userName, newUser.Email),
		// 	},
		// 	PageTitle:     "Email Verification",
		// 	Subject:       "Email Verification: Welcome Aboard!",
		// 	Preheader:     "User Account Created! ",
		// 	BodyTitle:     fmt.Sprintf("Welcome, %s", userName),
		// 	FirstBodyText: emailText,
		// }
		// emailData.Button.Text = "Verify Email"
		// emailData.Button.URL = vURL

		// // Send A Welcome/Verification Email to User
		// emailBody := ut.ProcessEmail(emailData)
		// go ut.SendEmail(emailBody, sendGridKey)

		// Decode to json so it can be used in the activity log
		// decoded, _ := json.Marshal(newUser)

		// // Log activity
		// aLog := md.ActivityLog{
		// 	AdminID:     adminID,
		// 	AppLocation: "Salary Update Function",
		// 	Action:      "Created user *" + string(decoded) + "*",
		// }
		// defer acs.CreateActivityLog(aLog)

		// resp := ut.Message(true, ut.Translate(tParam, r))
		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["verify_url"] = vURL
		ut.Respond(w, r, resp)

	}

}

// UpdateUserEndpoint - Update authenticated user account
func UpdateUserEndpoint(uss us.UserService, acs act.ActivityLogService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get Admin Token Data
		tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
		adminID := string(tokenData.AdminID)

		userID := chi.URLParam(r, "userID")
		user := md.User{}

		err := json.NewDecoder(r.Body).Decode(&user)

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

		// Validate user input
		err = us.ValidateUpdates(user, r)
		fmt.Println("ValidateUpdates ", err)

		if err != nil {
			validationFields := us.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		// Get the user
		us := uss.GetUser(userID)

		formerEmployerData := us

		us.FirstName = user.FirstName
		us.LastName = user.LastName
		us.Email = user.Email
		us.About = user.About
		us.Phone = user.Phone
		us.Avatar = user.Avatar
		us.Thumb = user.Thumb

		// Update a user
		updatedEmp, errParam, err := uss.UpdateUser(us)
		if err != nil {
			// Check if the error is duplication error
			cErr := ut.CheckUniqueError(r, err)
			if cErr != nil {
				ut.ErrorRespond(http.StatusBadRequest, w, r, ut.Message(false, cErr.Error()))
				return
			}
			// Respond with an errortranslated
			ut.ErrorRespond(http.StatusBadRequest, w, r, ut.Message(false, ut.Translate(errParam, r)))
			return
		}

		tParam = tr.TParam{
			Key:          "general.update_success",
			TemplateData: nil,
			PluralCount:  nil,
		}

		// Decode to json so it can be used in the activity log
		decoded, _ := json.Marshal(updatedEmp)
		decodedFormer, _ := json.Marshal(formerEmployerData)

		// Log activity
		aLog := md.ActivityLog{
			AdminID:     adminID,
			AppLocation: "Salary Update Function",
			Action:      "Updated user from *" + string(decodedFormer) + "* to *" + string(decoded) + "*",
		}
		defer acs.CreateActivityLog(aLog)

		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["user"] = updatedEmp
		ut.Respond(w, r, resp)

	}

}

// VerifyUserEmailEndpoint - Verify user Email
func VerifyUserEmailEndpoint(uss us.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		emailToken := chi.URLParam(r, "emailToken")

		checkUser := uss.GetUser(userID)

		// if user already exist send pincode to the user for verification/Authentication
		if checkUser == nil {

			tParam := tr.TParam{
				Key:          "error.invalid_token",
				TemplateData: nil,
				PluralCount:  nil,
			}

			resp := ut.Message(true, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		isTokenValid := ut.ValidatePassword(checkUser.EmailToken, emailToken)
		if !isTokenValid {

			tParam := tr.TParam{
				Key:          "error.invalid_token",
				TemplateData: nil,
				PluralCount:  nil,
			}

			resp := ut.Message(true, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		checkUser.IsEmailVerified = true
		checkUser.Status = "active"

		uss.UpdateUser(checkUser)

		tParam := tr.TParam{
			Key:          "general.verification_success",
			TemplateData: nil,
			PluralCount:  nil,
		}

		resp := ut.Message(true, ut.Translate(tParam, r))
		ut.Respond(w, r, resp)
	}
}
