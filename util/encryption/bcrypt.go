package encryption

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var PASSWORD_FORMAT = regexp.MustCompile(`^.{8,}$`)

const (
	COST_PASSWORD int = 10
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), COST_PASSWORD)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePasswordRequirement(password string) bool {
	return PASSWORD_FORMAT.MatchString(password)
}
