package chat

// IncomingHTTPTextMsg 接受通过HTTP方式发来的文本消息
func IncomingHTTPTextMsg(fromUUID string, msgText string, isToGroup bool, toEntityUUID string) error {

	// 非群发
	if !isToGroup {
		err := DispatchToSingleOnlineUser(fromUUID, msgText, toEntityUUID)
		if err != nil {
			//todo 不在线，存入未读消息
		}
	} else {
		// 群发
		// todo

	}

	// todo archive

	return nil
}

func IncomingHTTPFileMsg(fromUUID string, msgType string, msgFile string, isToGroup bool, toEntityUUID string) {

}

func IncomingWebSocketTextMsg() {

}
