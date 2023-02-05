package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// GenerateUUID returns a unique identifier or an error if it fails to generate one
func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %v", err)
	}
	return id.String(), nil
}
