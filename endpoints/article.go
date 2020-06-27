package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	art "github.com/ebikode/eLearning-core/domain/article"
	usr "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

// GetArticleEndpoint fetch a single article
func GetArticleEndpoint(ars art.ArticleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		articleID, _ := strconv.ParseUint(chi.URLParam(r, "articleID"), 10, 64)

		var article *md.Article
		article = ars.GetArticle(uint(articleID))
		resp := ut.Message(true, "")
		resp["article"] = article
		ut.Respond(w, r, resp)
	}
}

// GetAdminArticlesEndpoint fetch a single article
func GetAdminArticlesEndpoint(ars art.ArticleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		page, limit := ut.PaginationParams(r)

		articles := ars.GetArticles(page, limit)

		var nextPage int
		if len(articles) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["articles"] = articles
		ut.Respond(w, r, resp)
	}

}

// GetUserArticlesEndpoint all articles of a tutor user
func GetUserArticlesEndpoint(ars art.ArticleService, userType string) http.HandlerFunc {

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

		articles := ars.GetArticlesByUser(userID, page, limit)

		var nextPage int
		if len(articles) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["articles"] = articles
		ut.Respond(w, r, resp)
	}

}

// GetCourseArticlesEndpoint all articles of a tutor user
func GetCourseArticlesEndpoint(ars art.ArticleService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		courseID, _ := strconv.ParseUint(chi.URLParam(r, "courseID"), 10, 64)

		articles := ars.GetArticlesByCourse(int(courseID))

		resp := ut.Message(true, "")
		resp["articles"] = articles
		ut.Respond(w, r, resp)
	}

}

// CreateArticleEndpoint ...
func CreateArticleEndpoint(ars art.ArticleService, uss usr.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get User Token Data
		tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
		userID := string(tokenData.UserID)
		payload := art.Payload{}
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

		// Validate article input
		err = art.Validate(payload, r)
		if err != nil {
			validationFields := art.ValidationFields{}
			fmt.Println("third Error check", validationFields)
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		article := md.Article{
			UserID:   userID,
			CourseID: payload.CourseID,
			Title:    payload.Title,
			Body:     payload.Body,
		}

		// Create a article
		newArticle, errParam, err := ars.CreateArticle(article)
		if err != nil {
			// Check if the error is dupliarticleion error
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
		resp["article"] = newArticle
		ut.Respond(w, r, resp)

	}

}

// UpdateArticleEndpoint
func UpdateArticleEndpoint(ars art.ArticleService) http.HandlerFunc {

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
		// Parse the article id param
		articleID, pErr := strconv.ParseUint(chi.URLParam(r, "articleID"), 10, 64)

		if pErr != nil || uint(articleID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		articlePayload := art.Payload{}

		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&articlePayload)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate article input
		err = art.ValidateUpdates(articlePayload, r)
		if err != nil {
			validationFields := art.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Get the article
		checkArticle := ars.GetArticle(uint(articleID))

		if checkArticle.UserID != userID {
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
		checkArticle.Title = articlePayload.Title
		checkArticle.Body = articlePayload.Body
		checkArticle.Status = articlePayload.Status

		// Update a article
		updatedArticle, errParam, err := ars.UpdateArticle(checkArticle)
		if err != nil {
			// Check if the error is dupliarticleion error
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
		resp["article"] = updatedArticle
		ut.Respond(w, r, resp)

	}

}
