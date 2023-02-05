package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

// VerifyJwt verifies that the JWT is valid and has been signed with the secret
func VerifyJwt(jwtToken string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}
	if jwtToken == "" {
		return "", errors.New("jwtToken cannot be empty")
	}

	// Decode and verify the token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Convert the claims to a JSON string
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Return the claims
	return string(claimsJson), nil
}
