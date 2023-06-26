package util

import (
	"ConfBackend/model"
	S "ConfBackend/services"
)

func PadChatMsgFileUrl(msgs *[]model.ImMessage) {
	for i := range *msgs {
		if (*msgs)[i].MsgType != "text" && (*msgs)[i].FileTypeURI != "" {
			(*msgs)[i].FileTypeURI = ConcatFullFileUrl((*msgs)[i].FileTypeURI)
		}
	}
}

func PadSingleChatMsgFileUrl(msg *model.ImMessage) {
	if msg.MsgType != "text" && msg.FileTypeURI != "" {
		msg.FileTypeURI = ConcatFullFileUrl(msg.FileTypeURI)
	}
}

func ConcatFullFileUrl(fileUri string) string {
	pref := S.S.Conf.Chat.ServerFileUrlPrefix
	// if pref ends with /
	if pref[len(pref)-1] == '/' {
		return pref + fileUri
	} else {
		return pref + "/" + fileUri
	}

}
