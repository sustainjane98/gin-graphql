package db

import (
	"example/config"
	"example/models"
)

type userService struct {
}

func (u userService) Create(d *models.UserDao) error {
	db, err := config.DB()

	if err != nil {
		return err
	}

	db.Create(d)

	return nil
}
