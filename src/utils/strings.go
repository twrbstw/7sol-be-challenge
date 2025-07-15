package utils

import "golang.org/x/crypto/bcrypt"

func HashString(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes), err
}
