package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	act "github.com/ebikode/eLearning-core/domain/activity_log"
	aps "github.com/ebikode/eLearning-core/domain/app_setting"
	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
	"github.com/go-chi/chi"
)

func GetAppSettingEndpoint(ast aps.AppSettingService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		stID, _ := strconv.ParseUint(chi.URLParam(r, "appsettingID"), 10, 64)

		var appsetting *md.AppSetting
		appsetting = ast.GetAppSetting(uint(stID))
		resp := ut.Message(true, "")
		resp["appsetting"] = appsetting
		ut.Respond(w, r, resp)
	}
}

func GetAppSettingByKeyEndpoint(ast aps.AppSettingService, customerType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		sKey := chi.URLParam(r, "sKEY")

		var appsetting *md.AppSetting
		appsetting = ast.GetAppSettingByKey(sKey, customerType)
		resp := ut.Message(true, "")
		resp["appsetting"] = appsetting
		ut.Respond(w, r, resp)
	}
}

// Get AllAppSettingsEndpoint
func GetAppSettingsEndpoint(ast aps.AppSettingService, customerType string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var stegories []*md.AppSetting
		stegories = ast.GetAppSettings(customerType)
		resp := ut.Message(true, "")
		resp["app_settings"] = stegories
		ut.Respond(w, r, resp)
	}
}

func CreateAppSettingEndpoint(ast aps.AppSettingService, acs act.ActivityLogService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get Admin Token Data
		tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
		adminID := string(tokenData.AdminID)
		appsetting := md.AppSetting{}
		fmt.Println(appsetting)
		err := json.NewDecoder(r.Body).Decode(&appsetting)

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
		// Validate appsetting input
		err = aps.Validate(appsetting, r)
		if err != nil {
			validationFields := aps.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}

		// Create a appsetting
		st, errParam, err := ast.CreateAppSetting(appsetting)
		if err != nil {
			// Check if the error is duplistion error
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

		// Decode to json so it can be used in the activity log
		decoded, _ := json.Marshal(st)

		// Log activity
		aLog := md.ActivityLog{
			AdminID:     adminID,
			AppLocation: "App Setting Creation Function",
			Action:      "Created appsetting *" + string(decoded) + "*",
		}
		defer acs.CreateActivityLog(aLog)

		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["app_setting"] = st
		ut.Respond(w, r, resp)

	}

}

func UpdateAppSettingEndpoint(ast aps.AppSettingService, acs act.ActivityLogService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get Admin Token Data
		tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
		adminID := string(tokenData.AdminID)
		// Translation Param
		tParam := tr.TParam{
			Key:          "error.request_error",
			TemplateData: nil,
			PluralCount:  nil,
		}
		// Parse the appsetting id param
		stID, pErr := strconv.ParseUint(chi.URLParam(r, "appsettingID"), 10, 64)
		if pErr != nil || uint(stID) < 1 {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}
		appsetting := md.AppSetting{}
		// dECODE THE REQUEST BODY
		err := json.NewDecoder(r.Body).Decode(&appsetting)

		if err != nil {
			// Respond with an error translated
			resp := ut.Message(false, ut.Translate(tParam, r))
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return
		}

		// Validate appsetting input
		err = aps.ValidateUpdates(appsetting, r)
		if err != nil {
			validationFields := aps.ValidationFields{}
			b, _ := json.Marshal(err)
			// Respond with an errortranslated
			resp := ut.Message(false, ut.Translate(tParam, r))
			json.Unmarshal(b, &validationFields)
			resp["error"] = validationFields
			ut.ErrorRespond(http.StatusBadRequest, w, r, resp)
			return

		}
		fmt.Printf("stID:: %v \n", uint(stID))
		// Get the appsetting
		st := ast.GetAppSetting(uint(stID))
		// Set former appsetting to be used on activity log creation
		formerAppSetting := st
		// Assign new values
		st.Name = appsetting.Name
		// st.SKey = appsetting.SKey
		st.Value = appsetting.Value
		st.Comment = appsetting.Comment
		st.Status = appsetting.Status

		// Update a appsetting
		updatedAppSetting, errParam, err := ast.UpdateAppSetting(st)
		if err != nil {
			// Check if the error is duplistion error
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
		decoded, _ := json.Marshal(st)
		decodedFormer, _ := json.Marshal(formerAppSetting)

		// Log activity
		aLog := md.ActivityLog{
			AdminID:     adminID,
			AppLocation: "App Settings Update Function",
			Action:      "Updated app settings from *" + string(decodedFormer) + "* to *" + string(decoded) + "*",
		}
		defer acs.CreateActivityLog(aLog)

		resp := ut.Message(true, ut.Translate(tParam, r))
		resp["app_setting"] = updatedAppSetting
		ut.Respond(w, r, resp)

	}

}
