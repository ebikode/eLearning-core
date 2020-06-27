package utils

import (
	"net/http"

	tr "github.com/ebikode/eLearning-core/translation"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// REquired rule
func NotNilRule(r *http.Request, key string) []validation.Rule {
	field := Translate(tr.TParam{Key: key, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.NotNil.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": field},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// RequiredRule  rule
func RequiredRule(r *http.Request, key string) []validation.Rule {
	field := Translate(tr.TParam{Key: key, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": field},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// REquired integer rule
func RequiredIntRule(r *http.Request, key string) []validation.Rule {
	field := Translate(tr.TParam{Key: key, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": field},
					PluralCount:  nil,
				},
				r,
			),
		),
		// is.Int.Error(
		// 	Translate(
		// 		tr.TParam{
		// 			Key:          "validation.number",
		// 			TemplateData: map[string]interface{}{"Field": field},
		// 			PluralCount:  nil,
		// 		},
		// 		r,
		// 	),
		// ),
	}
}

// Validation Rule for IDs
// ID cannot be empty
func IDRule(r *http.Request) []validation.Rule {
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": "id"},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for position
// position cannot be empty and must be an integer
func PositionRule(r *http.Request) []validation.Rule {
	position := Translate(tr.TParam{Key: "general.position", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": position},
					PluralCount:  nil,
				},
				r,
			),
		),
		// is.Int.Error(
		// 	Translate(
		// 		tr.TParam{
		// 			Key:          "validation.number",
		// 			TemplateData: map[string]interface{}{"Field": position},
		// 			PluralCount:  nil,
		// 		},
		// 		r,
		// 	),
		// ),
	}
}

// Validation Rule for scores
// position cannot be empty and must be an integer
func ScoreRule(r *http.Request) []validation.Rule {
	score := Translate(tr.TParam{Key: "general.score", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.NotNil.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": score},
					PluralCount:  nil,
				},
				r,
			),
		),
		// is.Int.Error(
		// 	Translate(
		// 		tr.TParam{
		// 			Key:          "validation.number",
		// 			TemplateData: map[string]interface{}{"Field": score},
		// 			PluralCount:  nil,
		// 		},
		// 		r,
		// 	),
		// ),
	}
}

// Validation Rule for mobile phone
// Phone cannot be empty, and the length must between 7 and 20
func PhoneRule(r *http.Request) []validation.Rule {
	phone := Translate(tr.TParam{Key: "general.phone", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": phone},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(7, 20).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": phone, "Min": 7, "Max": 20},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Credit Card No
// Card No cannot be empty and must be 16 digits
func CardNoRule(r *http.Request) []validation.Rule {
	cardNo := Translate(tr.TParam{Key: "general.card_no", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": cardNo},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(16, 16).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": cardNo, "Min": 16, "Max": 16},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for PIN
// PIN cannot be empty and must be 4 digits
func PinRule(r *http.Request) []validation.Rule {
	pin := Translate(tr.TParam{Key: "general.pin", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": pin},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(4, 4).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": pin, "Min": 4, "Max": 4},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for PIN
// PIN cannot be empty and must be 4 digits
func OTPRule(r *http.Request) []validation.Rule {
	otp := Translate(tr.TParam{Key: "general.otp", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": otp},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(6, 6).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": otp, "Min": 6, "Max": 6},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Credit Card Expiry month
// Card month cannot be empty and must be 2 digits long
func CardMonthRule(r *http.Request) []validation.Rule {
	cardMonth := Translate(tr.TParam{Key: "general.expiry_month", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": cardMonth},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(2, 2).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": cardMonth, "Min": 2, "Max": 2},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Credit Card Expiry year
// Card year cannot be empty and must be 2 digits long
func CardYearRule(r *http.Request) []validation.Rule {
	cardYear := Translate(tr.TParam{Key: "general.expiry_year", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": cardYear},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(2, 2).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": cardYear, "Min": 2, "Max": 2},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Credit Card Expiry year
// Card year cannot be empty and must be 2 digits long
func CardCvvRule(r *http.Request) []validation.Rule {
	cvv := "cvv"
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": cvv},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(2, 2).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": cvv, "Min": 3, "Max": 3},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for year
// The length must between 4 and 4
func YearRule(r *http.Request) []validation.Rule {
	year := Translate(tr.TParam{Key: "general.year", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Length(4, 4).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": year, "Min": 4, "Max": 4},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Customername
// Customername cannot be empty, and the length must between 3 and 20 and Alpha Numeric
func CustomernameRule(r *http.Request) []validation.Rule {
	customername := Translate(tr.TParam{Key: "general.customername", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": customername},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 20).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": customername, "Min": 3, "Max": 20},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for plan name
// SearchPlan name cannot be empty, and the length must between 3 and 50
func SearchPlanNameRule(r *http.Request) []validation.Rule {
	name := Translate(tr.TParam{Key: "general.name", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": name},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": name, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Name
// Name cannot be empty, and the length must between 3 and 50
func NameRule(r *http.Request) []validation.Rule {
	name := Translate(tr.TParam{Key: "general.name", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": name},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": name, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Short Code
// ShortCode cannot be empty, and the length must between 2 and 20
func ShortCodeRule(r *http.Request) []validation.Rule {
	name := Translate(tr.TParam{Key: "general.short_code", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": name},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(2, 20).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": name, "Min": 3, "Max": 20},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Question
// Question cannot be empty, and the length must between 10 and 100
func QuestionRule(r *http.Request) []validation.Rule {
	question := Translate(tr.TParam{Key: "general.question", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": question},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(10, 100).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": question, "Min": 10, "Max": 100},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for resource unique Key
// resource unique Key cannot be empty, and the length must between 3 and 50
func KeyRule(r *http.Request) []validation.Rule {
	langKey := Translate(tr.TParam{Key: "general.key", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": langKey},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": langKey, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Language Key
// Langauge Key cannot be empty, and the length must between 3 and 50
func LangKeyRule(r *http.Request) []validation.Rule {
	langKey := Translate(tr.TParam{Key: "general.lang_key", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": langKey},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": langKey, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for Description Language Key
// Langauge Key cannot be empty, and the length must between 3 and 50
func DescLangKeyRule(r *http.Request) []validation.Rule {
	langKey := Translate(tr.TParam{Key: "general.desc_lang_key", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": langKey},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": langKey, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Email Addresses
// Email cannot be empty, and must be a valid Email
func EmailRule(r *http.Request) []validation.Rule {
	email := Translate(tr.TParam{Key: "general.email", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		// is.Email.Error(
		// 	Translate(
		// 		tr.TParam{
		// 			Key:          "validation.email",
		// 			TemplateData: nil,
		// 			PluralCount:  nil,
		// 		},
		// 		r,
		// 	),
		// ),
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": email},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(5, 100).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": email, "Min": 5, "Max": 100},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Email Addresses
// Email cannot be empty, and must be a valid Email
func AboutRule(r *http.Request) []validation.Rule {
	about := Translate(tr.TParam{Key: "general.about", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": about},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(5, 150).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": about, "Min": 5, "Max": 150},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule for languagestring
// Lang cannot be empty, and the length must between 2 and 2c
func LangRule(r *http.Request) []validation.Rule {
	language := Translate(tr.TParam{Key: "general.language", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": language},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(2, 2).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": language, "Min": 2, "Max": 2},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation rule for Avatar
// Avatar length must between 5 and 200
func AvatarRule(r *http.Request) []validation.Rule {
	avatar := Translate(tr.TParam{Key: "general.avatar", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Length(5, 200).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": avatar, "Min": 5, "Max": 200},
					PluralCount:  nil,
				},
				r,
			),
		),
		is.URL.Error(
			Translate(
				tr.TParam{
					Key:          "validation.url",
					TemplateData: nil,
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation rule for Enum Fields
// Enum must be among the list enum
func EnumRule(r *http.Request, fieldNameKey string, enum ...interface{}) []validation.Rule {
	field := Translate(tr.TParam{Key: fieldNameKey, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": field},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.In(enum...).Error(
			Translate(
				tr.TParam{
					Key:          "validation.enum",
					TemplateData: map[string]interface{}{"Field": field, "Values": enum},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}
func EnumNotNilRule(r *http.Request, fieldNameKey string, enum ...interface{}) []validation.Rule {
	field := Translate(tr.TParam{Key: fieldNameKey, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.NotNil.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": field},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.In(enum...).Error(
			Translate(
				tr.TParam{
					Key:          "validation.enum",
					TemplateData: map[string]interface{}{"Field": field, "Values": enum},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Bank Name
// Bank name cannot be empty, Length between 3 t0 50
func BankNameRule(r *http.Request, key string) []validation.Rule {
	bankName := Translate(tr.TParam{Key: key, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": bankName},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 100).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": bankName, "Min": 3, "Max": 100},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Bank account
// account number cannot be empty, Length between 3 t0 50
func BankaccountRule(r *http.Request) []validation.Rule {
	acctNum := Translate(tr.TParam{Key: "general.account_number", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": acctNum},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": acctNum, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Routing Number
// Routing number cannot be empty, Length between 3 t0 50
func BankRoutingRule(r *http.Request, key string) []validation.Rule {
	rountingNum := Translate(tr.TParam{Key: key, TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		//  validation.Required.Error(
		// 		Translate(
		// 			tr.TParam{
		// 				Key:          "validation.required",
		// 				TemplateData: map[string]interface{}{"Field": rountingNum},
		// 				PluralCount:  nil,
		// 			},
		// 			r,
		// 		),
		// 	),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": rountingNum, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Security answer
// Answer cannot be empty, Length between 2 t0 50
func AnswerRule(r *http.Request) []validation.Rule {
	answer := Translate(tr.TParam{Key: "general.answer", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": answer},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": answer, "Min": 3, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Solution
// Solution cannot be empty, Length between 5 t0 200
func SolutionRule(r *http.Request) []validation.Rule {
	solution := Translate(tr.TParam{Key: "general.solution", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": solution},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(5, 200).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": solution, "Min": 5, "Max": 200},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Password
// Password cannot be empty, Length between 7 t0 50
func PasswordRule(r *http.Request) []validation.Rule {
	password := Translate(tr.TParam{Key: "general.password", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": password},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(7, 50).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": password, "Min": 7, "Max": 50},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Message
// Message cannot be empty, Length between 1 t0 1000
func MessageRule(r *http.Request) []validation.Rule {
	message := Translate(tr.TParam{Key: "general.message", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": message},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(1, 1000).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": message, "Min": 1, "Max": 1000},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Comment
// Comment Length between 3 to 200
func CommentRule(r *http.Request) []validation.Rule {
	comment := Translate(tr.TParam{Key: "general.comment", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Length(3, 200).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": comment, "Min": 3, "Max": 200},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Comment that is required
// Comment Length between 3 to 200
func RequiredCommentRule(r *http.Request) []validation.Rule {
	comment := Translate(tr.TParam{Key: "general.comment", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": comment},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Length(3, 200).Error(
			Translate(
				tr.TParam{
					Key:          "validation.length",
					TemplateData: map[string]interface{}{"Field": comment, "Min": 3, "Max": 200},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}

// Validation Rule For Money
// Money cannot be empty, must be digits and floats
func MoneyRule(r *http.Request) []validation.Rule {
	money := Translate(tr.TParam{Key: "general.money", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": money},
					PluralCount:  nil,
				},
				r,
			),
		),
		// is.Float.Error(
		// 	Translate(
		// 		tr.TParam{
		// 			Key:          "validation.float",
		// 			TemplateData: map[string]interface{}{"Field": money},
		// 			PluralCount:  nil,
		// 		},
		// 		r,
		// 	),
		// ),
	}
}

// Validation Rule For Duration
// Money cannot be empty, must be digits
func DurationRule(r *http.Request) []validation.Rule {
	duration := Translate(tr.TParam{Key: "general.duration", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": duration},
					PluralCount:  nil,
				},
				r,
			),
		),
		// is.Int.Error(
		// 	Translate(
		// 		tr.TParam{
		// 			Key:          "validation.number",
		// 			TemplateData: map[string]interface{}{"Field": duration},
		// 			PluralCount:  nil,
		// 		},
		// 		r,
		// 	),
		// ),
	}
}

// Validation Rule For Date
// Date cannot be empty, must be a valid date
func DateRule(r *http.Request) []validation.Rule {
	date := Translate(tr.TParam{Key: "general.date", TemplateData: nil, PluralCount: nil}, r)
	return []validation.Rule{
		validation.Required.Error(
			Translate(
				tr.TParam{
					Key:          "validation.required",
					TemplateData: map[string]interface{}{"Field": date},
					PluralCount:  nil,
				},
				r,
			),
		),
		validation.Date("2006-01-01 15:04:05").Error(
			Translate(
				tr.TParam{
					Key:          "validation.date",
					TemplateData: map[string]interface{}{"Field": date},
					PluralCount:  nil,
				},
				r,
			),
		),
	}
}
