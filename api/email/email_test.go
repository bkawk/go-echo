package email

import (
	"net/smtp"
	"testing"

	"bkawk/go-echo/api/models"
)

type mockEmailSender struct{}

func (m mockEmailSender) Send(from, password, server string, auth smtp.Auth, to []string, msg []byte) error {
	return nil
}

func TestWelcomeEmail(t *testing.T) {
	// Prepare test data
	u := &models.User{
		Username: "John Doe",
		Email:    "john.doe@example.com",
	}
	from := "test@example.com"
	password := "password"
	server := "smtp.example.com:587"
	sender := mockEmailSender{}

	// Call the function
	err := WelcomeEmail(u, from, password, server, sender)
	if err != nil {
		t.Errorf("WelcomeEmail returned unexpected error: %v", err)
	}
}
