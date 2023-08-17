package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type logFormatLocal struct {
	ClientIP     string
	TimeStamp    time.Time
	Method       string
	Path         string
	RequestProto string
	StatusCode   int
	Latency      time.Duration
	ErrorMessage string
}

// { ::1 - [Thu, 17 Aug 2023 17:33:38 WAT] " GET /getData HTTP/1.1 200
// 100.3µs "PostmanRuntime/7.32.3" "}
func FormatLogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf("{ %s - [%s] \" %s %s %s %d %s \"%s\" %s\"} \n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

// {"ClientIP":"::1","TimeStamp":"2023-08-17T17:30:27.8205963+01:00",
// "Method":"GET","Path":"/getData","RequestProto":"HTTP/1.1",
// "StatusCode":200,"Latency":532100,"ErrorMessage":""}
func FormatLogsJson(param gin.LogFormatterParams) string {
	params := &logFormatLocal{
		ClientIP:     param.ClientIP,
		TimeStamp:    param.TimeStamp,
		Method:       param.Method,
		Path:         param.Path,
		RequestProto: param.Request.Proto,
		StatusCode:   param.StatusCode,
		Latency:      param.Latency,
		ErrorMessage: param.ErrorMessage,
	}
	j, _ := json.Marshal(params)
	return string(j)
}
func DebugRoute(httpMethod string, absolutePath string, handlerName string, nuHandlers int) {
	log.Printf("Endpoint formatted as %v http method, %v path, %v handlerName, %v nuHandlers", httpMethod, absolutePath, handlerName, nuHandlers)
}
