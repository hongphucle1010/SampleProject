package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost (currently 10) is secure and performant for most cases.
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword compares a hashed password with a plaintext password.
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
