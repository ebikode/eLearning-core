package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	que "github.com/ebikode/eLearning-core/domain/question"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetQuestionEndpoint fetch a single question
// func GetQuestionEndpoint(qs que.QuestionService) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Get User Token Data
// 		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
// 		userID := string(tokenData.UserID)

// 		questionID := chi.URLParam(r, "questionID")
// 		applicationID := chi.URLParam(r, "applicationID")

// 		var question *md.Question
// 		question = qs.GetQuestion(applicationID, questionID)

// 		resp := ut.Message(true, "")
// 		resp["question"] = question
// 		ut.Respond(w, r, resp)
// 	}
// }

// GetAdminQuestionsEndpoint fetch a single question
func GetAdminQuestionsEndpoint(qs que.QuestionService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		questions := qs.GetQuestions(page, limit)

		var nextPage int
		if len(questions) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["questions"] = questions
		ut.Respond(w, r, resp)
	}

}

// GetCourseQuestionsEndpoint all questions of a tutor user
func GetCourseQuestionsEndpoint(qus que.QuestionService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		questions := qus.GetQuestionsByCourse(uint(courseID))

		resp := ut.Message(true, "")
		resp["questions"] = questions
		ut.Respond(w, r, resp)
	}

}

// CreateQuestionEndpoint ...
func CreateQuestionEndpoint(qs que.QuestionService, uss usr.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		payloads := []que.Payload{}
		err := json.NewDecoder(r.Body).Decode(&payloads)
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

		//  if the length of the payloads uploaded is less than 1
		if len(payloads) < 1 {
			tParam = tr.TParam{
				Key:          "validation.lesser",
				TemplateData: map[string]interface{}{"Min": 1},
				PluralCount:  nil,
			}
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		if len(payloads) > 50 {
			tParam = tr.TParam{
				Key:          "validation.greater",
				TemplateData: map[string]interface{}{"Max": 50},
				PluralCount:  nil,
			}
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// questions created that will be returned
		var createdQuestions []*md.Question

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

		// loop through the payloads and create them
		for _, v := range payloads {

			// Validate question input
			err = que.Validate(v, r)
			if err != nil {
				validationFields := que.ValidationFields{}
				fmt.Println("third Error check", validationFields)
				b, _ := json.Marshal(err)
				// Respond with an errortranslated
				resp := ut.Message(false, ut.Translate(tParam, r))
				json.Unmarshal(b, &validationFields)
				resp["error"] = validationFields
				ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
				return

			}

			countQuestions := qs.CountQuestionsByCourse(v.CourseID)

			v.OptionA = strings.ToLower(v.OptionA)
			v.OptionB = strings.ToLower(v.OptionB)
			v.OptionC = strings.ToLower(v.OptionC)
			v.Answer = strings.ToLower(v.Answer)

			if countQuestions < 50 {
				question := md.Question{
					CourseID: v.CourseID,
					Question: v.Question,
					OptionA:  v.OptionA,
					OptionB:  v.OptionB,
					OptionC:  v.OptionC,
					Answer:   v.Answer,
					Solution: v.Solution,
				}

				// Create a question
				newQuestion, errParam, err := qs.CreateQuestion(question)
				if err != nil {
					// Check if the error is dupliquestionion error
					cErr := ut.CheckUniqueError(r, err)
					if cErr != nil {
						ut.ErrorRespond(http.StatusBadRequest, w, r, ut.Message(false, cErr.Error()))
						return
					}
					// Respond with an errortranslated
					ut.ErrorRespond(http.StatusBadRequest, w, r, ut.Message(false, ut.Translate(errParam, r)))
					return
				}
				// add the created category to the slice
				createdQuestions = append(createdQuestions, newQuestion)
			}

		}

		tParam = tr.TParam{
			Key:          "general.resource_created",
			TemplateData: nil,
			PluralCount:  nil,
		}

		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["questions"] = createdQuestions
		ut.Respond(w, r, resp)

	}

}

// UpdateQuestionEndpoint
func UpdateQuestionEndpoint(qs que.QuestionService) http.HandlerFunc {

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
		// Parse the question id param
		questionID := chi.URLParam(r, "questionID")

		if len(questionID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		questionPayload := que.Payload{}

		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&questionPayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate question input
		err = que.ValidateUpdates(questionPayload, r)
		if err != nil {
			validationFields := que.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Get the question
		checkQuestion := qs.GetQuestion(questionID)

		if checkQuestion.Course.UserID != userID {
			tParam = tr.TParam{
				Key:          "error.course_not_found",
				TemplateData: nil,
				PluralCount:  nil,
			}
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		questionPayload.OptionA = strings.ToLower(questionPayload.OptionA)
		questionPayload.OptionB = strings.ToLower(questionPayload.OptionB)
		questionPayload.OptionC = strings.ToLower(questionPayload.OptionC)
		questionPayload.Answer = strings.ToLower(questionPayload.Answer)

		// Assign new values
		checkQuestion.Question = questionPayload.Question
		checkQuestion.OptionA = questionPayload.OptionA
		checkQuestion.OptionB = questionPayload.OptionB
		checkQuestion.OptionC = questionPayload.OptionC
		checkQuestion.Answer = questionPayload.Answer
		checkQuestion.Solution = questionPayload.Solution
		checkQuestion.Status = questionPayload.Status

		// Update a question
		updatedQuestion, errParam, err := qs.UpdateQuestion(checkQuestion)
		if err != nil {
			// Check if the error is dupliquestionion error
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
		resp["question"] = updatedQuestion
		ut.Respond(w, r, resp)

	}

}
