package actions

import (
	"example/models"
	"example/services/encryption"
)

func (ActionServices) EncryptPassword(c *models.UserDao) error {
	hashedPW, err := encryption.Bcrypt().Hash([]byte(c.Password))

	if err != nil {
		return err
	}

	passwordString := string(hashedPW)
	c.Password = passwordString

	return nil
}
