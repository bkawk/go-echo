package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

const (
	RefreshTokenLength = 32
	RefreshTokenExpiry = 24 * 7 * time.Hour // Refresh token is valid for 7 days
)

// GenerateRefreshToken generates a new refresh token
func GenerateRefreshToken() (string, error) {
	b := make([]byte, RefreshTokenLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// ValidateRefreshToken checks if a refresh token is valid and has not expired
func ValidateRefreshToken(refreshToken string, createdAt time.Time) error {
	if len(refreshToken) == 0 {
		return errors.New("refresh token is missing")
	}

	if time.Since(createdAt) > RefreshTokenExpiry {
		return errors.New("refresh token has expired")
	}

	return nil
}
