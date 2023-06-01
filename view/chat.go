package view

import (
	"ConfBackend/chat"
	com "ConfBackend/common"
	S "ConfBackend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path/filepath"
)

func SendMsg(c *gin.Context) {
	//multipart/form-data
	msgType := c.PostForm("msgType")
	msgText := c.PostForm("msgText")
	msgFile, _ := c.FormFile("msgFile")
	isToGroup := c.PostForm("isToGroup")
	toEntityUUID := c.PostForm("toEntityUUID")

	// get fromUserUuid from middleware
	fromUserUuid := c.GetString("uuid")
	if fromUserUuid == "" {
		com.Error(c, "X-User-UUID 不能为空")
		return
	}
	if !checkMsgTypeAllowed(msgType) {
		com.Error(c, "msgType不合法，允许"+allowedMsgTypeToStr())
		return
	}
	if toEntityUUID == "" {
		com.Error(c, "toEntityUUID 收信者uuid不能为空")
		return
	}
	//log.Println(msgType, msgText, msgFile, isToGroup, toEntityUUID)
	if isToGroup != "0" && isToGroup != "1" {
		com.Error(c, "isToGroup只能为0或1")
		return
	}

	// todo 先不开发群聊功能
	if isToGroup == "1" {
		com.Error(c, "群聊功能暂不开发")
		return
	}

	//////// Param Check passed below
	err := error(nil)

	if fromUserUuid == toEntityUUID {
		com.Error(c, "不支持自己给自己发信")
		return
	}

	// 0 is false  1 is true
	isToGroupBool := false
	if isToGroup == "1" {
		isToGroupBool = true
	}

	// msg fromUserUuid init
	msgUuid := ""
	switch msgType {
	case "text":
		{
			if msgText == "" {
				com.Error(c, "类型为text时，msgText不能为空")
				return
			}

			msgUuid, err = chat.IncomingHTTPTextMsg(fromUserUuid, msgText, isToGroupBool, toEntityUUID)
			if err != nil {
				com.Error(c, err.Error())
				return
			}

		}
	case "image":
		{
			if msgFile.Size == 0 {
				com.Error(c, "类型为image时，msgFile不能为空")
				return
			}
			// save file

			fileType := filepath.Ext(msgFile.Filename)
			//uuid without "-"
			newFileName := uuid.New().String() + fileType
			newFileDir := filepath.Join(S.S.Conf.Chat.SaveStaticFileDirPrefix, newFileName)
			err := c.SaveUploadedFile(msgFile, newFileDir)
			if err != nil {
				com.Error(c, "文件保存失败")
				return
			}

			msgUuid, err = chat.IncomingHTTPFileMsg(fromUserUuid, msgType, isToGroupBool, toEntityUUID, newFileName, newFileDir)
			if err != nil {
				com.Error(c, err.Error())
				return
			}

		}
	case "audio":
		{
			if msgFile.Size == 0 {
				com.Error(c, "类型为image时，msgFile不能为空")
				return
			}
			// save file
			fileType := filepath.Ext(msgFile.Filename)
			newFileName := uuid.New().String() + fileType
			newFileDir := filepath.Join(S.S.Conf.Chat.SaveStaticFileDirPrefix, newFileName)
			err := c.SaveUploadedFile(msgFile, newFileDir)
			if err != nil {
				com.Error(c, "文件保存失败")
				return
			}

			msgUuid, err = chat.IncomingHTTPFileMsg(fromUserUuid, msgType, isToGroupBool, toEntityUUID, newFileName, newFileDir)
			if err != nil {
				com.Error(c, err.Error())
				return
			}

		}
	}

	com.OkD(c, msgUuid)

}

func AllOnline(c *gin.Context) {
	com.OkD(c, chat.GetAllOnlineUsers())
}

func ChatHistory(c *gin.Context) {
	fromUuid := c.GetString("uuid")
	sinceMid := c.PostForm("sinceMid")
	isGroupMsg := c.PostForm("isGroupMsg")
	objectEntityUuid := c.PostForm("objectEntityUUID")

	if fromUuid == "" {
		com.Error(c, "X-User-UUID 禁止为空")
		return
	}
	if isGroupMsg != "0" && isGroupMsg != "1" {
		com.Error(c, "isGroupMsg只能为0或1")
		return
	}
	if objectEntityUuid == "" {
		com.Error(c, "objectEntityUuid不能为空")
		return
	}

	if isGroupMsg == "1" {
		com.Error(c, "群聊功能暂不开发")
		return
	} else {
		// 私聊
		com.OkD(c, chat.GetChatHistory(fromUuid, objectEntityUuid, sinceMid))
	}

}
