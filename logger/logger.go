package logger

import (
	"fmt"
	"ginpractice/config"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() {
	customFormatter := new(logrus.JSONFormatter)
	customFormatter.PrettyPrint = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)
	logLevel := config.Appconfig.GetString("Logging.level")
	setLogLevel(logLevel)
	if config.Appconfig.GetBool("Logging.stdout") {
		logrus.New().Out = os.Stdout
	} else {
		options := os.O_WRONLY | os.O_CREATE | os.O_APPEND
		file, err := os.OpenFile(config.Appconfig.GetString("Logging.path"), options, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			fmt.Println("Failed to log to file", err.Error())
		}
	}
}
func setLogLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}
}
func LogError(message string, err error, c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"path":         c.Request.RequestURI,
		"error":        err.Error(),
		"version":      c.Request.Header.Get("version"),
		"x-request-id": c.Request.Header.Get("x-request-id"),
	}).Error(message)
}
func LogInfo(message string, c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"path":         c.Request.RequestURI,
		"version":      c.Request.Header.Get("version"),
		"x-request-id": c.Request.Header.Get("x-request-id"),
	}).Info(message)
}
func LogFatal(message string, err error, c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"path":         c.Request.RequestURI,
		"error":        err.Error(),
		"version":      c.Request.Header.Get("version"),
		"x-request-id": c.Request.Header.Get("x-request-id"),
	}).Fatal(message)
}
func LogDebug(message string, path string, xRequestID string, errors error) {
	logrus.WithFields(logrus.Fields{
		"path":         path,
		"error":        errors.Error(),
		"version":      config.Appconfig.GetString("version"),
		"x-request-id": xRequestID,
	}).Debug(message)
}
//Panic will exit with a status code of 2
func PanicLn(message string)  {
	logrus.Panicln(message)
}

//Fatal will exit with a status code of 1
func FatalLn(message string)  {
	logrus.Fatalln(message)
}

//Just log the message as Info
func InfoLn(message string)  {
	logrus.Infoln(message)
}

//Just log the message as a warning
func WarnLn(message string)  {
	logrus.Warnln(message)
}

//Just log the message as Debug
func DebugLn(message string)  {
	logrus.Debugln(message)
}

//Just print the message using print
func PrintLn(message string)  {
	logrus.Println(message)
}