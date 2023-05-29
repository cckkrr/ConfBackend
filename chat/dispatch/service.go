package dispatch

import (
	"ConfBackend/chat"
	"ConfBackend/model"
	"errors"
	"github.com/google/uuid"
	"time"
)

// DispatchToSingleOnlineUser Dispatch 尝试将消息下发至一个用户，用户不在线则返回error
func DispatchToSingleOnlineUser(fromUUID string, msgText string, toEntityUUID string) error {
	newUuid := uuid.New()
	msg := model.ImMessage{
		UUID:         newUuid.String(),
		MsgType:      "text",
		TextTypeText: msgText,
		FileTypeURI:  "",
		FromUserUUID: fromUUID,
		ToEntityUUID: toEntityUUID,
		CreatedAt:    time.Now(),
	}
	err := chat.SendOnlineMsg(msg)
	if err != nil {
		return errors.New("send online msg failed")
		//todo send to unread channel
	}
	return nil
}
