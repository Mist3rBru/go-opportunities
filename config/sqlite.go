package config

import (
	"gopportunities/internal/domain"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSQLite() (*gorm.DB, error) {
	logger := NewLogger("sqlite")
	dbPath := "./db.dev"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating it")
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("failer to create %s: %v", dbPath, err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("sqlite opening error: ", err)
		return nil, err
	}

	err = db.AutoMigrate(&domain.Opening{})
	if err != nil {
		logger.Error("sqlite automigration error: ", err)
		return nil, err
	}

	return db, nil
}
