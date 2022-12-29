package models

import (
	"gorm.io/gorm"
)

type UserDao struct {
	gorm.Model
	Age      int
	Username string
	Password string
}
