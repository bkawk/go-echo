package utils

import (
	"errors"
	"os"
	"regexp"

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

	// Validate the format of the JWT using a regular expression
	jwtRegex := "^[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*$"
	match, err := regexp.MatchString(jwtRegex, jwtToken)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("jwtToken does not match the expected format")
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

	// Verify that the token was correctly signed and not tampered with
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Check that the claims contain an ID
	id, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("claims do not contain an ID")
	}

	// Return the ID
	return id, nil
}
