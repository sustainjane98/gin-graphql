package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	db *gorm.DB
}

func (m DatabaseManager) Connect() (*gorm.DB, error) {

	if m.db != nil {
		return m.db, nil
	}

	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	m.db = db

	return db, err
}

var dbm *DatabaseManager

func DB() *DatabaseManager {

	if dbm == nil {
		dbm = &DatabaseManager{}
	}

	return dbm

}
