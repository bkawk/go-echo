package email

import (
	"net/smtp"
	"os"

	"bkawk/go-echo/api/models"
)

// Send welcome email
func WelcomeEmail(u *models.User) error {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	to := u.Email
	message := `Subject: Welcome to Our System
	<p>Dear ` + u.Username + `,</p>
	<p>Welcome to our system!</p>
	<p>Best regards,</p>
	<p>Support Team</p>`

	smtpServer := os.Getenv("SMTP_SERVER")
	auth := smtp.PlainAuth("", from, password, smtpServer)
	err := smtp.SendMail(smtpServer, auth, from, []string{to}, []byte(message))
	return err
}
