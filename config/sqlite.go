package config

import (
	"gopportunities/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := NewLogger("sqlite")
	dbPath := "./db.dev"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating it")
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Error("failer to create db.dev", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("sqlite opening error: ", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Error("sqlite automigration error: ", err)
		return nil, err
	}

	return db, nil
}
