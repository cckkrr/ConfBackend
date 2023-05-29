package chat

import (
	"ConfBackend/chat/dispatch"
)

func IncomingHTTPTextMsg(fromUUID string, msgText string, isToGroup bool, toEntityUUID string) {

	err := dispatch.DispatchToSingleOnlineUser(fromUUID, msgText, toEntityUUID)
	if err != nil {
		//todo 不在线，存入未读消息
	}
}

func IncomingHTTPFileMsg(fromUUID string, msgType string, msgFile string, isToGroup bool, toEntityUUID string) {

}

func IncomingWebSocketTextMsg() {

}
