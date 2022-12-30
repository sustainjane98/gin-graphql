package models

import (
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	gorm.Model
	Email     string
	Firstname string
	Lastname  string
	Age       int
	Username  string
	Password  string
	Birthdate time.Time
}

// TableName overrides the table name used by User to `profiles`
func (UserDao) TableName() string {
	return "users"
}
