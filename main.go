package main

import (
	"os"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint: true,
	})
	
	logrus.WithField("Debug", "Creating File").Debug("Starting file creation")
	file, err := os.Create("logrus.log")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "os.Create",
			"error": true,
		}).Error(err)
	}
	logrus.WithField("Info", "Created File").Debug("End file creation")
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	

	router := gin.Default()
	router.GET("/getData", getData)
	logrus.Infof("starting server on port 8080")
	router.Run(":8080")
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi I am a get data method",
	})
}
