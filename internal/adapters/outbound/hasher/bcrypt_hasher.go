package hasher

import (
	"errors"
	"seven-solutions-challenge/internal/app/ports"

	e "seven-solutions-challenge/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
}

func NewBcryptHasher() ports.IHasher {
	return &BcryptHasher{}
}

// HashPassword implements ports.IHasher.
func (b *BcryptHasher) HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(bytes), err
}

// ComparePassword implements ports.IHasher.
func (b *BcryptHasher) ComparePassword(hashed string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return errors.New(e.ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD)
	}
	return nil
}
