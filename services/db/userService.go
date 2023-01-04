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

func (userService) FindByUsername(username string, fields ...interface{}) (*models.UserDao, error) {
	db, err := config.DB()

	if err != nil {
		return nil, err
	}

	inst := &models.UserDao{}
	db = db.Where("username = ?", username)

	if fields != nil {
		db = db.Select(fields)
	}

	err = db.First(inst).Error

	if err != nil {
		return nil, err
	}

	return inst, nil

}
