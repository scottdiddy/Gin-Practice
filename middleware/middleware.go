package middleware

import (
	"encoding/json"
	"fmt"
	"ginpractice/config"
	"ginpractice/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LogRequestInfo(c *gin.Context) {
	accessLogMap := make(map[string]string)
	accessLogMap["request_time"] = time.Now().String()
	accessLogMap["request_method"] = c.Request.Method
	accessLogMap["request_url"] = c.Request.RequestURI
	accessLogMap["request_proto"] = c.Request.Proto
	accessLogMap["request_ua"] = c.Request.UserAgent()
	accessLogMap["request_referer"] = c.Request.Referer()
	accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
	accessLogMap["request_client_ip"] = c.ClientIP()
	accessLogMap["x-request-id"] = c.Request.Header.Get("x-request-id")
	accessLogMap["version"] = config.Appconfig.GetString("version")
	logData, err := json.Marshal(accessLogMap)
	if err != nil {
		fmt.Println(err)
	}
	accessLogJSON := string(logData)
	logger.PrintLn(accessLogJSON)
	c.Next()

}
