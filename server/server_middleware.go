package server

import (
	com "ConfBackend/common"
	"ConfBackend/task"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
)

// 打印请求中间件
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

func MustHasUserUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.GetHeader("X-User-UUID")
		if uuid == "" {
			com.Error(c, "HTTP请求头中缺少X-User-UUID")
			c.Abort()
			return
		}

		if !task.HaveValidUser(uuid) {
			com.Error(c, "找不到X-User-UUID对应的用户："+uuid)
			c.Abort()
			return
		}

		c.Set("uuid", uuid)
	}
}
