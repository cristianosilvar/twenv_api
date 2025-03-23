package main

import (
	config "twenv/config"
	router "twenv/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("error initializing config: %v", err)
		return
	}
	// Initialize the router
	router.Initialize()

}
