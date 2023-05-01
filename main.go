package main

import (
	"github.com/Wtheodoro/gopportunities/config"
	"github.com/Wtheodoro/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger(("main"))

	//  initialize config
	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: %v", err)
		return
	}

	//  initialize router

	router.Initializ()
}