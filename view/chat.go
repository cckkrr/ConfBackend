package view

import (
	"ConfBackend/chat"
	com "ConfBackend/common"
	"github.com/gin-gonic/gin"
	"log"
)

func SendMsg(c *gin.Context) {
	//multipart/form-data
	msgType := c.PostForm("msgType")
	msgText := c.PostForm("msgText")
	msgFile := c.PostForm("msgFile")
	isToGroup := c.PostForm("isToGroup")
	toEntityUUID := c.PostForm("toEntityUUID")

	// get uuid from middleware
	uuid := c.GetString("uuid")
	if uuid == "" {
		com.Error(c, "uuid is empty")
		return
	}
	if !checkMsgTypeAllowed(msgType) {
		com.Error(c, "msgType不合法，允许"+allowedMsgTypeToStr())
		return
	}
	if toEntityUUID == "" {
		com.Error(c, "toEntityUUID不能为空")
	}
	log.Println(msgType, msgText, msgFile, isToGroup, toEntityUUID)
	if isToGroup != "0" && isToGroup != "1" {
		com.Error(c, "isToGroup只能为0或1")
		return
	}
	// 0 is false  1 is true
	isToGroupBool := false
	if isToGroup == "1" {
		isToGroupBool = true
	}

	switch msgType {
	case "text":
		{
			if msgText == "" {
				com.Error(c, "类型为text时，msgText不能为空")
				return
			}
			go chat.IncomingHTTPTextMsg(uuid, msgText, isToGroupBool, toEntityUUID)

		}
	case "image":
		{
			if msgFile == "" {
				com.Error(c, "类型为image时，msgFile不能为空")
				return
			}
			go chat.IncomingHTTPFileMsg(uuid, msgType, msgFile, isToGroupBool, toEntityUUID)

		}
	case "audio":
		{
			if msgFile == "" {
				com.Error(c, "类型为audio时，msgFile不能为空")
				return
			}
			go chat.IncomingHTTPFileMsg(uuid, msgType, msgFile, isToGroupBool, toEntityUUID)

		}
	}

	com.Ok(c)

}

func AllOnline(c *gin.Context) {
	com.OkD(c, chat.GetAllOnlineUsers())
}
