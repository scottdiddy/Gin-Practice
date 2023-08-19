package main

import (
	"ginpractice/src/config"
	"ginpractice/src/logger"
	"ginpractice/src/router"
	// "ginpractice/src/database"

)

func main() {
	config.Init()
	logger.Init()
	logger.InfoLn("Logger initialized successfully")
	// database.Init()
	// if config.Appconfig.GetBool("seeddata"){
	// 	//Logic to seed data to database
	// 	logger.InfoLn("Data seeded successfully")
	// }
	logger.InfoLn("Started router initialization")
	router.Init()
	logger.InfoLn("Router Initialized successfully")

}
	

