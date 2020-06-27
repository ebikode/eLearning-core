package utils

import (
	// "encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"html/template"
	"net/http"
)

type StaticMessage struct {
	CssClass string
	Message  string
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	render.JSON(w, r, data)
}

func ErrorRespond(errorCode int, w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	render.JSON(w, r, data)
}

// ServeStaticResponse for serving html response
func ServeStaticResponse(isSuccess bool, message string, w http.ResponseWriter) {
	t, err := template.ParseFiles("./templates/response.html")
	if err != nil {
		fmt.Println(err)
	}
	res := StaticMessage{
		CssClass: "success-color",
		Message:  message,
	}

	if !isSuccess {
		res.CssClass = "error-color"
		res.Message = message
	}

	t.Execute(w, res)
}
