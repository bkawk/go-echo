package utils

import (
	"regexp"
)

type PasswordError struct {
	Password string `json:"password,omitempty"`
}

func (e *PasswordError) Error() string {
	return e.Password
}

var (
	uppercaseRegex = regexp.MustCompile(`[A-Z]+`)
	lowercaseRegex = regexp.MustCompile(`[a-z]+`)
	digitRegex     = regexp.MustCompile(`\d+`)
	specialRegex   = regexp.MustCompile(`[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]+`)
)

func CheckPasswordStrength(password string) error {
	if len(password) < 8 {
		return &PasswordError{Password: "password must be at least 8 characters long"}
	}
	if !uppercaseRegex.MatchString(password) {
		return &PasswordError{Password: "password must contain at least one uppercase letter"}
	}
	if !lowercaseRegex.MatchString(password) {
		return &PasswordError{Password: "password must contain at least one lowercase letter"}
	}
	if !digitRegex.MatchString(password) {
		return &PasswordError{Password: "password must contain at least one digit"}
	}
	if !specialRegex.MatchString(password) {
		return &PasswordError{Password: "password must contain at least one special character"}
	}
	return nil
}
