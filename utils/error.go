package utils

import (
	"net/http"
	"strings"

	tr "github.com/ebikode/eLearning-core/translation"
	"github.com/go-sql-driver/mysql"
)

const (
	// ER_DUP_ENTRY
	ER_DUP_ENTRY      = 1062
	AdminStaffID      = "uix_admins_staff_id"
	AdminPhone        = "uix_admins_phone"
	AdminEmail        = "uix_admins_email"
	AccountPhone      = "uix_accounts_phone"
	AccountEmail      = "uix_accounts_email"
	CustomerUsername  = "uix_customers_username"
	CustomerPhone     = "uix_customers_phone"
	CustomerEmail     = "uix_customers_email"
	SearchPlanName    = "uix_categories_name"
	SearchPlanLangKey = "uix_categories_lang_key"
	AppSettingName    = "uix_app_settings_name"
	AppSettingKey     = "uix_app_settings_s_key"
)

// UniqueDuplicateError - encapsulates database duplication error
type UniqueDuplicateError struct {
	Key     string // Translation key
	Request *http.Request
}

func (e *UniqueDuplicateError) Error() string {
	field := Translate(tr.TParam{Key: e.Key, TemplateData: nil, PluralCount: nil}, e.Request)
	errorMsg := Translate(
		tr.TParam{
			Key:          "validation.unique",
			TemplateData: map[string]interface{}{"Field": field},
			PluralCount:  nil,
		}, e.Request)

	return errorMsg
}

// isUniqueConstraintError Check unique constraint
func isUniqueConstraintError(err error, constraintName string) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		containsString := strings.Contains(err.Error(), constraintName)
		return mysqlErr.Number == ER_DUP_ENTRY && containsString
	}
	return false
}

// CheckUniqueError Check DB conraint error
func CheckUniqueError(r *http.Request, err error) error {

	if isUniqueConstraintError(err, AdminEmail) || isUniqueConstraintError(err, CustomerEmail) ||
		isUniqueConstraintError(err, AccountEmail) {
		return &UniqueDuplicateError{Key: "general.email", Request: r}
	}
	if isUniqueConstraintError(err, AdminPhone) || isUniqueConstraintError(err, CustomerPhone) ||
		isUniqueConstraintError(err, AccountPhone) {
		return &UniqueDuplicateError{Key: "general.phone", Request: r}
	}
	if isUniqueConstraintError(err, AdminStaffID) {
		return &UniqueDuplicateError{Key: "general.staff_id", Request: r}
	}
	if isUniqueConstraintError(err, CustomerUsername) {
		return &UniqueDuplicateError{Key: "general.username", Request: r}
	}
	if isUniqueConstraintError(err, SearchPlanName) ||
		isUniqueConstraintError(err, AppSettingName) {
		return &UniqueDuplicateError{Key: "general.name", Request: r}
	}
	if isUniqueConstraintError(err, AppSettingKey) {
		return &UniqueDuplicateError{Key: "general.key", Request: r}
	}
	return nil
}
