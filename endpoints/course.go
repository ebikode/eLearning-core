package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	cor "github.com/ebikode/eLearning-core/domain/course"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetCourseEndpoint fetch a single course
func GetCourseEndpoint(cos cor.CourseService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		var course *md.Course
		course = cos.GetCourse(uint(courseID))
		resp := ut.Message(true, "")
		resp["course"] = course
		ut.Respond(w, r, resp)
	}
}

// GetUserCourseEndpoint fetch a single course
func GetUserCourseEndpoint(cos cor.CourseService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		var course *md.Course
		course = cos.GetSingleCourseByUserID(userID, uint(courseID))
		resp := ut.Message(true, "")
		resp["course"] = course
		ut.Respond(w, r, resp)
	}
}

// GetCoursesEndpoint fetch a single course
func GetCoursesEndpoint(cos cor.CourseService, userType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		courses := cos.GetCourses(page, limit, userType)

		var nextPage int
		if len(courses) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["courses"] = courses
		ut.Respond(w, r, resp)
	}

}

// GetUserCoursesEndpoint all courses of a tutor user
func GetUserCoursesEndpoint(cos cor.CourseService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)

		page, limit := ut.PaginationParams(r)

		courses := cos.GetCoursesByUser(userID, page, limit)

		var nextPage int
		if len(courses) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["courses"] = courses
		ut.Respond(w, r, resp)
	}

}

// CreateCourseEndpoint ...
func CreateCourseEndpoint(cos cor.CourseService, uss usr.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		payload := cor.Payload{}
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

		// Validate course input
		err = cor.Validate(payload, r)
		if err != nil {
			validationFields := cor.ValidationFields{}
			fmt.Println("third Error check", validationFields)
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		course := md.Course{
			UserID:              userID,
			Title:               payload.Title,
			Description:         payload.Description,
			Image:               payload.Image,
			DurationPerQuestion: payload.DurationPerQuestion,
			Mode:                payload.Mode,
		}

		// Generate ref
		ref := ut.RandomBase64String(8, "elref")

		course.ReferenceNo = ref

		// Create a course
		newCourse, errParam, err := cos.CreateCourse(course)
		if err != nil {
			// Check if the error is duplicourseion error
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
		resp["course"] = newCourse
		ut.Respond(w, r, resp)

	}

}

// UpdateCourseEndpoint ...
func UpdateCourseEndpoint(cos cor.CourseService) http.HandlerFunc {

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
		// Parse the course id param
		courseID, pErr := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)
		if pErr != nil || uint(courseID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		coursePayload := cor.Payload{}
		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&coursePayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate course input
		err = cor.ValidateUpdates(coursePayload, r)
		if err != nil {
			validationFields := cor.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}
		fmt.Printf("courseID:: %v \n", uint(courseID))
		// Get the course
		checkCourse := cos.GetCourse(uint(courseID))

		if checkCourse.UserID != userID {
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
		checkCourse.Title = coursePayload.Title
		checkCourse.Description = coursePayload.Description
		checkCourse.Image = coursePayload.Image
		checkCourse.Mode = coursePayload.Mode

		// Update a course
		updatedCourse, errParam, err := cos.UpdateCourse(checkCourse)
		if err != nil {
			// Check if the error is duplicourseion error
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
		resp["course"] = updatedCourse
		ut.Respond(w, r, resp)

	}

}

// UpdateCourseStatusEndpoint ...
func UpdateCourseStatusEndpoint(cos cor.CourseService) http.HandlerFunc {

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
		// Parse the course id param
		courseID, pErr := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)
		if pErr != nil || uint(courseID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		coursePayload := cor.Payload{}
		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&coursePayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate course input
		err = cor.ValidateUpdates(coursePayload, r)
		if err != nil {
			validationFields := cor.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}
		fmt.Printf("courseID:: %v \n", uint(courseID))
		// Get the course
		checkCourse := cos.GetCourse(uint(courseID))

		if checkCourse.UserID != userID {
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
		checkCourse.Status = coursePayload.Status

		// Update a course
		updatedCourse, errParam, err := cos.UpdateCourse(checkCourse)
		if err != nil {
			// Check if the error is duplicourseion error
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
		resp["course"] = updatedCourse
		ut.Respond(w, r, resp)

	}

}
