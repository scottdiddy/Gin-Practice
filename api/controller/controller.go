package controller

import (
	"encoding/json"
	"ginpractice/api/model"
	"ginpractice/logger"

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