package main

import (
	"ginpractice/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Use(middleware.Authenticate)
	admin := router.Group("/admin", middleware.CustomBasicAuth("fidel", "pass"))
	{
		admin.GET("/getData2", middleware.AddHeader, getData2)
		admin.GET("/getData3", getData3)
	}

	router.GET("/getData", getData)

	router.Run(":8080")

}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I am get Data handler 1",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I am get Data handler 2",
	})
}
func getData3(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I am get Data handler 3",
	})
}
