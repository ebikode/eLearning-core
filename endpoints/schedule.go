package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	shd "github.com/ebikode/eLearning-core/domain/schedule"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetScheduleEndpoint fetch a single schedule
func GetScheduleEndpoint(shs shd.ScheduleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		scheduleID, _ := strconv.ParseUint(chi.URLParam(r, "scheduleID"), 10, 64)

		var schedule *md.Schedule
		schedule = shs.GetSchedule(uint(scheduleID))
		resp := ut.Message(true, "")
		resp["schedule"] = schedule
		ut.Respond(w, r, resp)
	}
}

// GetAdminSchedulesEndpoint fetch a single schedule
func GetAdminSchedulesEndpoint(shs shd.ScheduleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		schedules := shs.GetSchedules(page, limit)

		var nextPage int
		if len(schedules) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["schedules"] = schedules
		ut.Respond(w, r, resp)
	}

}

// GetCourseSchedulesEndpoint all schedules of a tutor user
func GetCourseSchedulesEndpoint(shs shd.ScheduleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		schedules := shs.GetSchedulesByCourse(uint(courseID))

		resp := ut.Message(true, "")
		resp["schedules"] = schedules
		ut.Respond(w, r, resp)
	}

}

// GetTutorCourseSchedulesEndpoint all schedules of a tutor user
func GetTutorCourseSchedulesEndpoint(shs shd.ScheduleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)

		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		schedules := shs.GetSchedulesByCourseOwner(userID, uint(courseID))

		resp := ut.Message(true, "")
		resp["schedules"] = schedules
		ut.Respond(w, r, resp)
	}

}

// CreateScheduleEndpoint ...
func CreateScheduleEndpoint(shs shd.ScheduleService, uss usr.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		payload := shd.Payload{}
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

		// Validate schedule input
		err = shd.Validate(payload, r)
		if err != nil {
			validationFields := shd.ValidationFields{}
			fmt.Println("third Error check", validationFields)
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		schedule := md.Schedule{
			CourseID:       payload.CourseID,
			WeekDay:        payload.WeekDay,
			TimeFromHour:   payload.TimeFromHour,
			TimeFromMinute: payload.TimeFromMinute,
			TimeToHour:     payload.TimeToHour,
			TimeToMunite:   payload.TimeToMunite,
		}

		// Create a schedule
		newSchedule, errParam, err := shs.CreateSchedule(schedule)
		if err != nil {
			// Check if the error is duplischeduleion error
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
		resp["schedule"] = newSchedule
		ut.Respond(w, r, resp)

	}

}

// UpdateScheduleEndpoint
func UpdateScheduleEndpoint(shs shd.ScheduleService) http.HandlerFunc {

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
		// Parse the schedule id param
		scheduleID, pErr := strconv.ParseUint(chi.URLParam(r, "scheduleID"), 10, 64)

		if pErr != nil || uint(scheduleID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		schedulePayload := shd.Payload{}

		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&schedulePayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate schedule input
		err = shd.ValidateUpdates(schedulePayload, r)
		if err != nil {
			validationFields := shd.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Get the schedule
		checkSchedule := shs.GetSchedule(uint(scheduleID))

		if checkSchedule.Course.UserID != userID {
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
		checkSchedule.WeekDay = schedulePayload.WeekDay
		checkSchedule.TimeFromHour = schedulePayload.TimeFromHour
		checkSchedule.TimeFromMinute = schedulePayload.TimeFromMinute
		checkSchedule.TimeToHour = schedulePayload.TimeToHour
		checkSchedule.TimeToMunite = schedulePayload.TimeToMunite
		checkSchedule.Status = schedulePayload.Status

		// Update a schedule
		updatedSchedule, errParam, err := shs.UpdateSchedule(checkSchedule)
		if err != nil {
			// Check if the error is duplischeduleion error
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
		resp["schedule"] = updatedSchedule
		ut.Respond(w, r, resp)

	}

}
