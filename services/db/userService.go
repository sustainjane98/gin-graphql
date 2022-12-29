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

	err = db.Create(d).Error

	if err != nil {
		return err
	}

	return nil
}
