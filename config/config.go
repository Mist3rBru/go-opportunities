package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = initSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
