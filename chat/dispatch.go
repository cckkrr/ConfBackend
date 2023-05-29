package chat

import (
	"ConfBackend/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func SendOnlineMsg(msg model.ImMessage) error {
	if !IsUserOnline(msg.ToEntityUUID) {
		return fmt.Errorf("user %s is not online", msg.ToEntityUUID)
	}
	//put in that user's message queue
	msgStr, _ := json.Marshal(msg)
	WsConnectionManager.Users[msg.ToEntityUUID].Messages <- string(msgStr)
	return nil

}

// DispatchToSingleOnlineUser Dispatch 尝试将消息下发至一个用户，用户不在线则返回error
func DispatchToSingleOnlineUser(fromUUID string, msgText string, toEntityUUID string) error {
	newUuid := uuid.New()
	msg := model.ImMessage{
		UUID:         "mid-" + newUuid.String(),
		MsgType:      "text",
		TextTypeText: msgText,
		FileTypeURI:  "",
		FromUserUUID: fromUUID,
		ToEntityUUID: toEntityUUID,
		CreatedAt:    time.Now(),
	}
	err := SendOnlineMsg(msg)
	if err != nil {
		return errors.New("send online msg failed")

	}
	return nil
}
