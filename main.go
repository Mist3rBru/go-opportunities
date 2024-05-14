package main

import (
	"gopportunities/config"
	"gopportunities/handler"
	"gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config init error: ", err)
		return
	}

	handler.Init()
	router.Init()
}
