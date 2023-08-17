package middleware

import (
	"encoding/base64"
	"strings"
	"github.com/gin-gonic/gin"
)

var verifiedAccounts = map[string]string{
	"fidelis":  "fidexxx",
	"febe":     "fineboy",
	"benjamin": "freshboy",
}

func CustomBasicAuth1(c *gin.Context){
	user, passedUserPassword, hasAuth := c.Request.BasicAuth()
	if !hasAuth{
		c.AbortWithStatusJSON(401, gin.H{
			"error" : "Invalid authorization header, user or password",
		})
		return
	}
	actualUserPassword, ok := verifiedAccounts[strings.ToLower(user) ] 
	if !ok || actualUserPassword != passedUserPassword {
		c.AbortWithStatusJSON(401, gin.H{
			"error" : "Invalid username or password",
		})
		return
	}
	c.Set(gin.AuthUserKey, user)
	c.Next()
}
func CustomBasicAuth2() gin.HandlerFunc{
	return gin.BasicAuth(gin.Accounts{
		"fidelis":  "fidexxx",
		"febe":     "fineboy",
		"benjamin": "freshboy",
	})
}
func CustomBasicAuth3(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Authorization header missing",
		})
		return
	}
	authParts := strings.SplitN(authHeader, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Basic" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Invalid authorization header format",
		})
		return
	}
	decodedBytes, err:= base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Failed to decode authorization header",
		})
		return
	}

	credential := string(decodedBytes)
	user, passedUserPassword, found:= strings.Cut(credential, ":")
	if !found {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Invalid credentials format",
		})
		return
	}
	actualUserPassword, ok := verifiedAccounts[strings.ToLower(user) ] 
	if !ok || actualUserPassword != passedUserPassword {
		c.AbortWithStatusJSON(401, gin.H{
			"error" : "Invalid username or password",
		})
		return
	}
	c.Set(gin.AuthUserKey, user)
	c.Next()

}



func AddHeader(c *gin.Context) {

	//c.Header is shortcut for this
	c.Writer.Header().Set("Namesss", "Fidel")
	c.Next()
}
