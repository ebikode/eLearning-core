package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	act "github.com/ebikode/eLearning-core/domain/activity_log"
	jon "github.com/ebikode/eLearning-core/domain/journal"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetJournalEndpoint fetch a single journal
func GetJournalEndpoint(jos jon.JournalService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		journalID, _ := strconv.ParseUint(chi.URLParam(r, "journalID"), 10, 64)

		var journal *md.Journal
		journal = jos.GetJournal(uint(journalID))
		resp := ut.Message(true, "")
		resp["journal"] = journal
		ut.Respond(w, r, resp)
	}
}

// GetAdminJournalsEndpoint fetch a single journal
func GetAdminJournalsEndpoint(jos jon.JournalService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		salaries := jos.GetJournals(page, limit)

		var nextPage int
		if len(salaries) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["salaries"] = salaries
		ut.Respond(w, r, resp)
	}

}

// CreateJournalEndpoint ...
func CreateJournalEndpoint(jos jon.JournalService, uss usr.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		payload := jon.Payload{}
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

		// Validate journal input
		err = jon.Validate(payload, r)
		if err != nil {
			validationFields := jon.ValidationFields{}
			fmt.Println("third Error check", validationFields)
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		journal := md.Journal{
			UserID:   userID,
			CourseID: payload.CourseID,
			Title:    payload.Title,
			Body:     payload.Body,
		}

		// Create a journal
		newJournal, errParam, err := jos.CreateJournal(journal)
		if err != nil {
			// Check if the error is duplijournalion error
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
		resp["journal"] = newJournal
		ut.Respond(w, r, resp)

	}

}

// UpdateJournalEndpoint
func UpdateJournalEndpoint(jos jon.JournalService, acs act.ActivityLogService) http.HandlerFunc {

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
		// Parse the journal id param
		journalID, pErr := strconv.ParseUint(chi.URLParam(r, "journalID"), 10, 64)

		if pErr != nil || uint(journalID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		journalPayload := jon.Payload{}

		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&journalPayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate journal input
		err = jon.ValidateUpdates(journalPayload, r)
		if err != nil {
			validationFields := jon.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Get the journal
		checkJournal := jos.GetJournal(uint(journalID))

		if checkJournal.UserID != userID {
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
		checkJournal.Title = journalPayload.Title
		checkJournal.Body = journalPayload.Body
		checkJournal.Status = journalPayload.Status

		// Update a journal
		updatedJournal, errParam, err := jos.UpdateJournal(checkJournal)
		if err != nil {
			// Check if the error is duplijournalion error
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
		resp["journal"] = updatedJournal
		ut.Respond(w, r, resp)

	}

}
