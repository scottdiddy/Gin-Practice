package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.POST("/getDataPost", getDataPost)
	auth := gin.BasicAuth(gin.Accounts{
		"Fidelis" : "Pass2",
	})
	
	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)

	}
	client := router.Group("/client")
	{
		client.GET("/getQueryString", getQueryString)
	}
	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
// func auth(c *gin.Context) {
// 	if c.GetHeader("Authorization") != "Bearer token1234" {
// 		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
// 		return
// 	}
// 	c.Next()
// }

// http://localhost:8080/getData
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, I am a gin get Data bobo",
	})
}

// http://localhost:8080//getUrlData/:name/:age
func getUrlData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"message": "Hi, I am a get Url Data",
		"name":    name,
		"age":     age,
	})
}

// http://localhost:8080/getQueryString?name=Fidel&age=2
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Hi, I am a gin get Query string",
		"name":    name,
		"age":     age,
	})
}

// http://localhost:8080/getDataPost
func getDataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := io.ReadAll(body)

	c.JSON(200, gin.H{
		"message": "Hi I am a gin post Data gee",
		"body":    string(value),
	})
}
