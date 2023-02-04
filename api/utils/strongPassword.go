package utils

import (
	"fmt"
	"regexp"
)

var (
	uppercaseRegex = regexp.MustCompile(`[A-Z]+`)
	lowercaseRegex = regexp.MustCompile(`[a-z]+`)
	digitRegex     = regexp.MustCompile(`\d+`)
	specialRegex   = regexp.MustCompile(`[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]+`)
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if !uppercaseRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !lowercaseRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !digitRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one digit")
	}
	if !specialRegex.MatchString(password) {
		return fmt.Errorf("password must contain at least one special character")
	}
	return nil
}
