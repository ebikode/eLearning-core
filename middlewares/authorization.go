package middleware

import (
	"net/http"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"
)

// Check if Admin IP has changee
func CheckAdminIPAddress() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
			ip := tokenData.IP

			// Set Error Message
			errorMsg := ut.Translate(tr.TParam{Key: "error.ip_changed", TemplateData: nil, PluralCount: nil}, r)

			// Get device info
			deviceInfo := ut.DetectDevice(r)

			if string(deviceInfo.IP) != ip {
				response := ut.Message(false, errorMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}

		return http.HandlerFunc(fn)
	}
}

// IsSuperAdmin - Checks if Admin is the super_admin of the account
func IsSuperAdmin() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
			role := string(tokenData.Role)

			// Set Error Message
			errorMsg := ut.Translate(tr.TParam{Key: "error.unauthorized", TemplateData: nil, PluralCount: nil}, r)

			// Validate
			if role != "super_admin" {
				response := ut.Message(false, errorMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}

		return http.HandlerFunc(fn)
	}
}

// IsManagerAdmin - Check if Admin is Sales Admin
func IsManagerAdmin() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
			role := string(tokenData.Role)

			// Set Error Message
			errorMsg := ut.Translate(tr.TParam{Key: "error.unauthorized", TemplateData: nil, PluralCount: nil}, r)

			// Validate
			if role != "super_admin" && role != "manager" {
				response := ut.Message(false, errorMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}

		return http.HandlerFunc(fn)
	}
}

// IsEditorAdmin - Checks if Admin is editor Admin
func IsEditorAdmin() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
			role := string(tokenData.Role)

			// Set Error Message
			errorMsg := ut.Translate(tr.TParam{Key: "error.unauthorized", TemplateData: nil, PluralCount: nil}, r)

			// Validate
			if role != "super_admin" && role != "editor" && role != "manager" {
				response := ut.Message(false, errorMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}

		return http.HandlerFunc(fn)
	}
}

// IsSalesAdmin - Check if Admin is sales or editor Admin
func IsSalesAdmin() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenData := r.Context().Value("tokenData").(*md.AdminTokenData)
			role := string(tokenData.Role)

			// Set Error Message
			errorMsg := ut.Translate(tr.TParam{Key: "error.unauthorized", TemplateData: nil, PluralCount: nil}, r)

			// Validate
			if role != "super_admin" && role != "sales" && role != "manager" {
				response := ut.Message(false, errorMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}

		return http.HandlerFunc(fn)
	}
}
