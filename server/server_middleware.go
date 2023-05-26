package server

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
)

func printRequest(c *gin.Context) {

	r, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		return
	}
	println(string(r))

}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Next()
	}
}
