package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error(err)
		return
	}

	router.Init()
}
