package handler

import (
	"gopportunities/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.NewLogger("handler")
	db = config.GetSQLite()
}
