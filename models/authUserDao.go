package models

import (
	"example/config"
	"gorm.io/gorm"
)

type UserDao struct {
	gorm.Model
	Age             int
	Username        string
	Password        string
	PasswordConfirm string
}

func (d *UserDao) Create() {

	db, err := config.DB()

	if err != nil {
		return
	}

	db.Create(d)
}
