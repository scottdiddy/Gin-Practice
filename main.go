package main

import (
	"ginpractice/config"
	"ginpractice/logger"
	"ginpractice/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger initialized successfully")

	logger.InfoLn("Started router initialization")
	router.Init()
	logger.InfoLn("Router Initialized successfully")

}
	

