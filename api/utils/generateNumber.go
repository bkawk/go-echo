package utils

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
)

// GenerateNumber returns a random number with n digits or an error if it fails to generate one
func GenerateNumber(n int) (int, error) {
	var number int64
	err := binary.Read(rand.Reader, binary.LittleEndian, &number)
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}
	return int(number % int64(math.Pow10(n))), nil
}
