package chat

import (
	"ConfBackend/chat/archive"
	"ConfBackend/chat/unread"
)

// IncomingHTTPTextMsg 接受通过HTTP方式发来的文本消息
func IncomingHTTPTextMsg(fromUUID string, msgText string, isToGroup bool, toEntityUUID string) (string, error) {

	// 首先存档聊天
	msg := GenNewMsg(fromUUID, "text", msgText, "", toEntityUUID)

	go archive.ArchiveSingleMsg(msg)

	go func() {
		// 非群发
		if !isToGroup {
			err := DispatchToSingleOnlineUser(msg)
			if err != nil {
				//todo 不在线，存入未读消息
				go unread.SaveSingleUnreadMsg(msg)
			}
		} else {
			// 群发
			// 先不开发

		}
	}()

	return msg.UUID, nil
}

func IncomingHTTPFileMsg(fromUUID string, msgType string, msgFile string, isToGroup bool, toEntityUUID string) {

}

func IncomingWebSocketTextMsg() {

}

// InitChatServices 聊天部分入口初始化函数
func InitChatServices() {

	// 启动聊天WS管理
	{
		WsConnectionManager = &ConnectionManager{
			Users:      make(map[string]*_User),
			Register:   make(chan *_User),
			Unregister: make(chan *_User),
		}

		go WsConnectionManager.startUserManagement()

	}

}
