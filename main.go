package main

import (
	"fmt"
	"ginpractice/logger"
	"ginpractice/middleware"
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	//Enforces colour for other terminals 
	//Must be put above router definition if it should work
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	router := gin.Default()

	//Was: GET    /getData      --> main.getData (3 handlers)
	//Now: Endpoint formatted as GET http method, /getData path, main.getData handlerName, 4 nuHandlers
	//Doesn't print to file no matter where put
	gin.DebugPrintRouteFunc = logger.DebugRoute

	f, _ := os.Create("ginlogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//Should be put below default writer initialization if you want it to also write to the file
	router.Use(gin.LoggerWithFormatter(logger.FormatLogs))
	router.GET("/getData", middleware.CustomBasicAuth3, getData)
	router.Run(":8080")
}
func getData(c *gin.Context) {
	username, _ := c.Get(gin.AuthUserKey)
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("You've been authorized user '%s'", username),
	})
}
