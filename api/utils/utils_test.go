package utils

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGenerateJWT_Success(t *testing.T) {
	// Setup
	id := "user123"
	os.Setenv("JWT_SECRET", "secret_key")

	// Test
	token, err := GenerateJWT(id)

	// Assertions
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if token == "" {
		t.Errorf("Expected token, got empty string")
	}
}

func TestGenerateJWT_MissingSecret(t *testing.T) {
	// Setup
	id := "user123"
	os.Unsetenv("JWT_SECRET")

	// Test
	_, err := GenerateJWT(id)

	// Assertions
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGenerateJWT_EmptyID(t *testing.T) {
	// Setup
	id := ""
	os.Setenv("JWT_SECRET", "secret_key")

	// Test
	_, err := GenerateJWT(id)

	// Assertions
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestValidatePasswordShortPassword(t *testing.T) {
	err := ValidatePassword("pass")
	expected := fmt.Errorf("password must be at least 8 characters long")

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("ValidatePassword(\"pass\") = %v; want %v", err, expected)
	}
}
func TestValidatePasswordNoUppercase(t *testing.T) {
	err := ValidatePassword("password")
	expected := fmt.Errorf("password must contain at least one uppercase letter")

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("ValidatePassword(\"password\") = %v; want %v", err, expected)
	}
}

func TestValidatePasswordNoLowercase(t *testing.T) {
	err := ValidatePassword("PASSWORD")
	expected := fmt.Errorf("password must contain at least one lowercase letter")

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("ValidatePassword(\"PASSWORD\") = %v; want %v", err, expected)
	}
}

func TestValidatePasswordNoDigit(t *testing.T) {
	err := ValidatePassword("Password")
	expected := fmt.Errorf("password must contain at least one digit")

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("ValidatePassword(\"Password\") = %v; want %v", err, expected)
	}
}

func TestValidatePasswordNoSpecialCharacter(t *testing.T) {
	err := ValidatePassword("Password1")
	expected := fmt.Errorf("password must contain at least one special character")

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("ValidatePassword(\"Password1\") = %v; want %v", err, expected)
	}
}

func TestValidatePasswordEmptyPassword(t *testing.T) {
	err := ValidatePassword("")
	expected := fmt.Errorf("password must be at least 8 characters long")

	if err == nil || err.Error() != expected.Error() {
		t.Errorf("ValidatePassword(\"\") = %v; want %v", err, expected)
	}
}

func TestGenerateRefreshTokenSpeed(t *testing.T) {
	start := time.Now()
	_, err := GenerateRefreshToken()
	if err != nil {
		t.Fatalf("failed to generate refresh token: %v", err)
	}
	elapsed := time.Since(start)

	if elapsed > time.Millisecond*50 {
		t.Fatalf("GenerateRefreshToken took too long: %v", elapsed)
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	token, err := GenerateRefreshToken()
	if err != nil {
		t.Errorf("GenerateRefreshToken returned error: %v", err)
	}

	if len(token) == 0 {
		t.Error("GenerateRefreshToken returned an empty string")
	}
}

func TestGenerateRefreshTokenSuccess(t *testing.T) {
	token, err := GenerateRefreshToken()
	if err != nil {
		t.Errorf("GenerateRefreshToken returned error: %v", err)
	}

	if len(token) != RefreshTokenLength {
		t.Errorf("GenerateRefreshToken returned a token of length %d, expected length %d", len(token), RefreshTokenLength)
	}
}

func TestGenerateRefreshTokenLength(t *testing.T) {
	token, err := GenerateRefreshToken()
	if err != nil {
		t.Errorf("GenerateRefreshToken returned error: %v", err)
	}

	if len(token) != RefreshTokenLength {
		t.Errorf("GenerateRefreshToken returned a token of length %d, expected length %d", len(token), RefreshTokenLength)
	}
}

func TestGenerateRefreshTokenUniqueness(t *testing.T) {
	tokens := make(map[string]bool)

	for i := 0; i < 100; i++ {
		token, err := GenerateRefreshToken()
		if err != nil {
			t.Errorf("GenerateRefreshToken returned error: %v", err)
		}

		if _, exists := tokens[token]; exists {
			t.Errorf("GenerateRefreshToken returned a duplicate token: %s", token)
			break
		}

		tokens[token] = true
	}
}

func TestGenerateUUID(t *testing.T) {
	ids := make(map[string]bool)

	for i := 0; i < 10000; i++ {
		id := GenerateUUID()

		if _, exists := ids[id]; exists {
			t.Errorf("GenerateUUID returned a duplicate identifier: %s", id)
			break
		}

		ids[id] = true
	}
}

func TestGenerateUUIDLength(t *testing.T) {
	expectedLength := 36
	for i := 0; i < 100; i++ {
		id := GenerateUUID()
		if len(id) != expectedLength {
			t.Errorf("GenerateUUID returned an unexpected length identifier: got %d, expected %d", len(id), expectedLength)
			break
		}
	}
}

func TestGenerateUUIDSpeed(t *testing.T) {
	numberOfIDs := 10000
	start := time.Now()
	for i := 0; i < numberOfIDs; i++ {
		GenerateUUID()
	}
	elapsed := time.Since(start)

	if elapsed > 100*time.Millisecond {
		t.Errorf("GenerateUUID took too long to generate %d UUIDs: %s", numberOfIDs, elapsed)
	}
}
