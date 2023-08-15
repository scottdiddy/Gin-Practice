package middleware

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context){
	// c.GetHeader()
	if c.Request.Header.Get("Token") != "auth" {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Token not present",
		})
		return
	}
	c.Next()
}
func CustomBasicAuth(username, password string) gin.HandlerFunc {
    return func(c *gin.Context) {
		// gin.BasicAuth
        user, pass, hasAuth := c.Request.BasicAuth()

        if !hasAuth || user != username || pass != password {
            c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
            c.AbortWithStatus(401)
            return
        }

        c.Next()
    }
}

func Authenticate2() gin.HandlerFunc{
	return func (c *gin.Context){
		if c.GetHeader("Token") != "auth" {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Token not present",
			})
			return
		}
		c.Next()
	}
}
func AddHeader(c *gin.Context){
	//c.Header
	c.Writer.Header().Set("Namesss", "Fidel")
	c.Next()
}