package emails

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
)

func SendWelcomeEmail(to, verificationLink string) error {

	var (
		smtpServer = os.Getenv("SMTP_SERVER")
		smtpPort   = os.Getenv("SMTP_PORT")
		username   = os.Getenv("EMAIL_FROM")
		password   = os.Getenv("SMTP_PASSWORD")
		from       = os.Getenv("EMAIL_FROM")
	)

	if smtpServer == "" || smtpPort == "" || username == "" || password == "" || from == "" {
		return fmt.Errorf("environment variable not set: SMTP_SERVER, SMTP_PORT, EMAIL_FROM, or SMTP_PASSWORD")
	}

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("failed to convert smtpPort to int: %v", err)
	}

	body := fmt.Sprintf(`
		<html>
			<body>
				<p>
					Welcome! Thank you for signing up.
				</p>
				<p>
					Please click the following link to verify your account:
					<br />
					<a href="%s">%s</a>
				</p>
				<p>
					<button style="background-color: #4CAF50; color: white; padding: 14px 20px; margin: 8px 0; border: none; cursor: pointer; width: 100%%;">Verify</button>
				</p>
				<p>
					Best regards,
					<br />
					The Team
				</p>
			</body>
		</html>
	`, verificationLink, verificationLink)
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Welcome\r\nContent-Type: text/html\r\n\r\n%s", from, to, body))
	auth := smtp.PlainAuth("", username, password, smtpServer)
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, port), auth, from, []string{to}, msg); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
