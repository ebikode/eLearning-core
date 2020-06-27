package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	app "github.com/ebikode/eLearning-core/domain/application"
	cou "github.com/ebikode/eLearning-core/domain/course"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetApplicationEndpoint fetch a single application
func GetApplicationEndpoint(aps app.ApplicationService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		applicationID, _ := strconv.ParseUint(chi.URLParam(r, "applicationID"), 10, 64)

		var application *md.Application
		application = aps.GetApplication(uint(applicationID))
		resp := ut.Message(true, "")
		resp["application"] = application
		ut.Respond(w, r, resp)
	}
}

// GetAdminApplicationsEndpoint fetch a single application
func GetAdminApplicationsEndpoint(aps app.ApplicationService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		applications := aps.GetApplications(page, limit)

		var nextPage int
		if len(applications) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["applications"] = applications
		ut.Respond(w, r, resp)
	}

}

// GetUserApplicationsEndpoint all applications of a tutor user
func GetUserApplicationsEndpoint(aps app.ApplicationService, userType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var userID string
		// if userType == admin then get the userId from the request parameter
		if userType == "admin" {
			userID = chi.URLParam(r, "userID")
		} else {
			// Get User Token Data
			tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
			userID = string(tokenData.UserID)
		}

		page, limit := ut.PaginationParams(r)

		applications := aps.GetUserApplications(userID)

		var nextPage int
		if len(applications) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["applications"] = applications
		ut.Respond(w, r, resp)
	}

}

// GetCourseApplicationsEndpoint all applications of a tutor user
func GetCourseApplicationsEndpoint(aps app.ApplicationService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		applications := aps.GetApplicationsByCourse(int(courseID))

		resp := ut.Message(true, "")
		resp["applications"] = applications
		ut.Respond(w, r, resp)
	}

}

// CreateApplicationEndpoint ...
func CreateApplicationEndpoint(aps app.ApplicationService, uss usr.UserService, cos cou.CourseService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		payload := app.Payload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		fmt.Println("second Error check", err)

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

		checkUser := uss.GetUser(userID)

		if checkUser == nil {
			tParam = tr.TParam{
				Key:          "error.user_not_found",
				TemplateData: nil,
				PluralCount:  nil,
			}
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		checkCourse := cos.GetCourse(payload.CourseID)

		if checkCourse == nil {
			tParam = tr.TParam{
				Key:          "error.course_not_found",
				TemplateData: nil,
				PluralCount:  nil,
			}
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate application input
		err = app.Validate(payload, r)
		if err != nil {
			validationFields := app.ValidationFields{}
			fmt.Println("third Error check", validationFields)
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		application := md.Application{
			UserID:   userID,
			CourseID: payload.CourseID,
		}

		// Create a application
		newApplication, errParam, err := aps.CreateApplication(application)
		if err != nil {
			// Check if the error is dupliapplicationion error
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
			Key:          "general.resource_created",
			TemplateData: nil,
			PluralCount:  nil,
		}

		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["application"] = newApplication
		ut.Respond(w, r, resp)

	}

}

// UpdateApplicationEndpoint ...
func UpdateApplicationEndpoint(aps app.ApplicationService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		// Translation Param
		tParam := tr.TParam{
			Key:          "error.request_error",
			TemplateData: nil,
			PluralCount:  nil,
		}
		// Parse the application id param
		applicationID, pErr := strconv.ParseUint(chi.URLParam(r, "applicationID"), 10, 64)

		if pErr != nil || uint(applicationID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		applicationPayload := app.Payload{}

		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&applicationPayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate application input
		err = app.ValidateUpdates(applicationPayload, r)
		if err != nil {
			validationFields := app.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Get the application
		checkApplication := aps.GetApplication(uint(applicationID))

		if checkApplication.UserID != userID {
			tParam = tr.TParam{
				Key:          "error.course_not_found",
				TemplateData: nil,
				PluralCount:  nil,
			}
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Assign new values
		checkApplication.Status = applicationPayload.Status

		// Update a application
		updatedApplication, errParam, err := aps.UpdateApplication(checkApplication)
		if err != nil {
			// Check if the error is dupliapplicationion error
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

		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["application"] = updatedApplication
		ut.Respond(w, r, resp)

	}

}
