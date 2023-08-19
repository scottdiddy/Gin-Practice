//Package controller contains the handler func for the different endpoints 
//called when the server is running
package controller

import (
	"encoding/json"
	"ginpractice/src/api/model"
	"ginpractice/src/logger"

	// "net/http"
	"github.com/gin-gonic/gin"
)
func GetData(c *gin.Context){
	model := model.GetData{
		Name: "Mark",
		Age: 30,
		City: "NY",
		Pincode: 777,
	}
	j, _ := json.Marshal(model)
	logger.LogInfo("In GetData", c)
	c.JSON(200, gin.H{
		"Data": string(j),
	})
}
func GetQueryStringData(c *gin.Context){
	name := c.Query("name")
	age:= c.Query("age")
	logger.LogInfo("In GetQueryStringData", c)
	c.JSON(200, gin.H{
		"Data": "In GetQueryStringData method",
		"name": name,
		"age": age,
	})
}
