package encryption

import (
	"example/errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
}

func (BcryptService) Hash(password []byte) ([]byte, error) {
	generateFromPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return generateFromPassword, nil
}

func (BcryptService) Compare(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err != nil {
		return errors.ValidationError{Reason: "password is incorrect"}
	}

	return nil
}
