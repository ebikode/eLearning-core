package utils

import (
	"net/http"

	tr "github.com/ebikode/eLearning-core/translation"
)

// Translate Tranlates string based on customer request lang param
func Translate(tParam tr.TParam, r *http.Request) string {
	params, ok := r.URL.Query()["lang"]
	lang := "en"

	// Change the language if lang param is provided
	if ok && len(params[0]) > 0 {
		lang = params[0]
	}
	// fmt.Printf("lang:: %v \n", lang)

	msg, _ := tr.Translate(tParam, lang)

	return msg
}

// TranslateString Tranlates string based on customer lang param provided
func TranslateString(tParam tr.TParam, lang string) string {
	msg, _ := tr.Translate(tParam, lang)
	return msg
}
