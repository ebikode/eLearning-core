package endpoints

import (
	"net/http"

	grd "github.com/ebikode/eLearning-core/domain/grade"
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
	// "fmt"
)

// GetGradeEndpoint fetch single grade
func GetGradeEndpoint(grs grd.GradeService, userType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		gradeID := chi.URLParam(r, "gradeID")
		var userID string
		// if userType == admin then get the userId from the request parameter
		if userType == "admin" {
			userID = chi.URLParam(r, "userID")
		} else {
			// Get User Token Data
			tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
			userID = string(tokenData.UserID)
		}

		var grade *md.Grade
		grade = grs.GetGrade(userID, gradeID)
		resp := ut.Message(true, "")
		resp["grade"] = grade
		ut.Respond(w, r, resp)
	}
}

// GetGradeReportsEndpoint ...
func GetGradeReportsEndpoint(grs grd.GradeService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		reports := grs.GetGradeReports()
		resp := ut.Message(true, "")
		resp["grade_reports"] = reports
		ut.Respond(w, r, resp)
	}
}

// GetGradesEndpoint Admin Enpoint for fetching all grades
func GetGradesEndpoint(grs grd.GradeService, userType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		page, limit := ut.PaginationParams(r)

		var grades []*md.Grade
		grades = grs.GetGrades(page, limit)

		var nextPage int
		if len(grades) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["grades"] = grades
		ut.Respond(w, r, resp)
	}
}

// GetUserGradesEndpoint Fetch All user Grades Endpoint
func GetUserGradesEndpoint(grs grd.GradeService, userType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		page, limit := ut.PaginationParams(r)

		var userID string
		// if userType == admin then get the userId from the request parameter
		if userType == "admin" {
			userID = chi.URLParam(r, "userID")
		} else {
			// Get User Token Data
			tokenData := r.Context().Value("tokenData").(*md.UserTokenData)
			userID = string(tokenData.UserID)
		}

		var grades []*md.Grade
		grades = grs.GetUserGrades(userID, page, limit)

		var nextPage int
		if len(grades) == limit {
			nextPage = page + 1
		}

		resp := ut.Message(true, "")
		resp["current_page"] = page
		resp["next_page"] = nextPage
		resp["limit"] = limit
		resp["grades"] = grades
		ut.Respond(w, r, resp)
	}
}
