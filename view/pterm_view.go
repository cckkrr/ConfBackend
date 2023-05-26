package view

import (
	com "ConfBackend/common"
	"github.com/gin-gonic/gin"
)

func PTerm(c *gin.Context) {
	com.OkM(c, "pong")
}

func FileReceived(c *gin.Context) {
	f, err := c.MultipartForm()
	if err != nil {
		com.Error(c, err.Error())
	}
	print(len(f.Value))
	com.Ok(c)

}
