package email

import (
	"net/smtp"

	"bkawk/go-echo/api/models"
)

type EmailSender interface {
	Send(from, password, server string, auth smtp.Auth, to []string, msg []byte) error
}

type SMTPSender struct{}

func (s SMTPSender) Send(from, password, server string, auth smtp.Auth, to []string, msg []byte) error {
	return smtp.SendMail(server, auth, from, to, msg)
}

func WelcomeEmail(u *models.User, from, password, server string, sender EmailSender) error {
	to := u.Email
	message := `Subject: Welcome to Our System
	<p>Dear ` + u.Username + `,</p>
	<p>Welcome to our system!</p>
	<p>Best regards,</p>
	<p>Support Team</p>`

	auth := smtp.PlainAuth("", from, password, server)
	return sender.Send(from, password, server, auth, []string{to}, []byte(message))
}
