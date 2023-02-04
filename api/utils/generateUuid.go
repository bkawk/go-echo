package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID returns a unique identifier
func GenerateUUID() string {
	return uuid.New().String()
}
