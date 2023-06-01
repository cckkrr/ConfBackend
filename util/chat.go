package util

import (
	"ConfBackend/model"
	S "ConfBackend/services"
)

func PadChatMsgFileUrl(msgs *[]model.ImMessage) {
	for i := range *msgs {
		if (*msgs)[i].MsgType != "text" && (*msgs)[i].FileTypeURI != "" {
			(*msgs)[i].FileTypeURI = S.S.Conf.Chat.ServerFileUrlPrefix + (*msgs)[i].FileTypeURI
		}
	}
}

func PadSingleChatMsgFileUrl(msg *model.ImMessage) {
	if msg.MsgType != "text" && msg.FileTypeURI != "" {
		msg.FileTypeURI = S.S.Conf.Chat.ServerFileUrlPrefix + msg.FileTypeURI
	}
}
