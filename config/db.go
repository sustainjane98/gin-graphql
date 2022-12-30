package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() (*gorm.DB, error) {

	if db != nil {
		return db, nil
	}

	dsn := "host=localhost user=jannik password=gorm dbname=ktor port=5432 sslmode=disable TimeZone=Europe/Berlin"

	dbI, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	db = dbI

	return db, err

}
