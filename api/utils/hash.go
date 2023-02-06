package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash(data []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		return "", fmt.Errorf("error writing to hash: %v", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
