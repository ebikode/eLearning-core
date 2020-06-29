package endpoints

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	adm "github.com/ebikode/eLearning-core/domain/admin"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
)

// CreateAdminEndpoint ...
func CreateAdminEndpoint(us adm.AdminService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		admin := md.Admin{}

		err := json.NewDecoder(r.Body).Decode(&admin)

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

		// Validate admin input
		err = adm.Validate(admin, r)

		fmt.Println(err)
		if err != nil {
			validationFields := adm.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Create an admin
		_, errParam, err := us.CreateAdmin(admin)
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
			Key:          "general.admin_reg_success",
			TemplateData: nil,
			PluralCount:  nil,
		}

		resp := ut.Message(true, ut.Translate(tParam, r))
		ut.Respond(w, r, resp)

	}

}

// GetAdminEndpoint ...
func GetAdminEndpoint(as adm.AdminService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var adminID string

		// Get Customer Token Data
		tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
		adminID = string(tokenData.AdminID)

		var admin *md.Admin
		admin = as.GetAdmin(adminID)
		dashoardData := as.GetAdminDashboardData()

		resp := ut.Message(true, "")
		resp["admin"] = admin
		resp["dashboard_data"] = dashoardData
		ut.Respond(w, r, resp)
	}
}
