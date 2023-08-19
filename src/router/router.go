//Package router helps to define a new router engine using gin.
//It works on whatever port that it specified in the config.yaml file and
//uses the middleware func specified in the middleware package
package router

import (
	"ginpractice/src/api/controller"
	"ginpractice/src/config"
	"ginpractice/src/middleware"

	"github.com/gin-gonic/gin"
)
func Init()  {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}
func NewRouter()  *gin.Engine{
	router := gin.New()
	resource := router.Group("/api")
	resource.Use(middleware.LogRequestInfo)
	{
		resource.GET("/GetData", controller.GetData)
		resource.GET("/GetQueryStringData", controller.GetQueryStringData)
	}
	return router
}