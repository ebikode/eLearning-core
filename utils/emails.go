package utils

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailData encapsulates email sending data
type EmailData struct {
	To             []*mail.Email
	PageTitle      string
	Preheader      string
	Subject        string
	BodyTitle      string
	FirstBodyText  string
	SecondBodyText string
	Button         struct {
		URL  string
		Text string
	}
}

// SendEmail - sends email to customers
func SendEmail(emailBody []byte, apiKey string) {
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = emailBody
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// ProcessEmail Email utilizing dynamic transactional templates
// Note: you must customize subject line of the dynamic template itself
// Note: you may not use substitutions with dynamic templates
func ProcessEmail(emailData EmailData) []byte {
	m := mail.NewV3Mail()

	address := "noreply@kimberly-ryan.com"
	name := "Kimberly Ryan"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)
	m.Subject = emailData.Subject

	m.SetTemplateID("d-5cad13d382184d6c913caa58ea0944b9")

	p := mail.NewPersonalization()
	tos := emailData.To

	p.AddTos(tos...)
	p.Subject = emailData.Subject

	title := emailData.BodyTitle
	firstBody := emailData.FirstBodyText
	secondBody := emailData.SecondBodyText
	url := emailData.Button.URL
	button := make(map[string]string)
	button["text"] = emailData.Button.Text
	button["url"] = url

	p.SetDynamicTemplateData("page_title", emailData.PageTitle)
	p.SetDynamicTemplateData("subject", emailData.Subject)
	p.SetDynamicTemplateData("preheader", emailData.Preheader)
	p.SetDynamicTemplateData("title", title)
	p.SetDynamicTemplateData("first_body", firstBody)
	p.SetDynamicTemplateData("second_body", secondBody)
	p.SetDynamicTemplateData("button", button)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}
