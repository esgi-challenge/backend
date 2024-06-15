package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

const port = "587"

type emailManager struct {
	username string
	password string
	host     string

	smtpAuth smtp.Auth
}

func InitEmailManager(username string, password string, host string) *emailManager {
	return &emailManager{username, password, host, smtp.PlainAuth("", username, password, host)}
}

func (e *emailManager) sendEmail(to []string, subject string, template *template.Template, templateData any) error {
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: %s\nSubject: %s \n%s\n\n", e.username, subject, mimeHeaders)))

	template.Execute(&body, templateData)

	return smtp.SendMail(fmt.Sprintf("%s:%s", e.host, port), e.smtpAuth, e.username, to, body.Bytes())
}

func (e *emailManager) SendInvitationEmail(to []string, name string, lastname string) error {
	t, err := template.ParseFiles("templates/emails/school-invitation.html")

	if err != nil {
		return err
	}

	templateData := struct {
		Name           string
		Lastname       string
		InvitationLink string
	}{
		Name:           name,
		Lastname:       lastname,
		InvitationLink: "https://www.youtube.com/watch?v=xvFZjo5PgG0",
	}

	err = e.sendEmail(to, "Studies Invitation", t, templateData)
	if err != nil {
		return err
	}

	return nil
}
