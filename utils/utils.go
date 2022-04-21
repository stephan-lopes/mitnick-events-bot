package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// IsEmptyString returns true if the string is empty
func IsEmptyString(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// HaveMinimumLength returns true if the string has a minimum length
func HaveMinLength(value string, min int) bool {
	return len(strings.TrimSpace(value)) >= min
}

// HaveMaximumLength returns true if the string has a maximum length
func HaveMaxLength(value string, max int) bool {
	return len(strings.TrimSpace(value)) <= max
}

// GenerateToken returns a random token
func GenerateToken(positions int) (string, error) {
	bytes := make([]byte, positions)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
