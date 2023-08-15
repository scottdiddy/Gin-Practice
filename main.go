package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)

	router.POST("/getDataPost", getDataPost)
	router.Run(":8080")
}

//http://localhost:8080/getData
func getData(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "Hi, I am a gin get Data bobo",
	})
}

//http://localhost:8080//getUrlData/:name/:age
func getUrlData(c *gin.Context)  {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"message": "Hi, I am a get Url Data",
		"name": name,
		"age" : age,
	})
}

//http://localhost:8080/getQueryString?name=Fidel&age=2
func getQueryString(c *gin.Context)  {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Hi, I am a gin get Query string",
		"name": name,
		"age" : age,
	})
}

//http://localhost:8080/getDataPost
func getDataPost(c *gin.Context)  {
	body := c.Request.Body
	value, _ := io.ReadAll(body)

	c.JSON(200, gin.H{
		"message": "Hi I am a gin post Data gee",
		"body" : string(value),
	})
}