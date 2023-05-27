package view

var allowedMsgTypes = []string{
	"text",
	"image",
	"audio",
}

func checkMsgTypeAllowed(msgType string) bool {
	for _, v := range allowedMsgTypes {
		if v == msgType {
			return true
		}
	}
	return false
}

func allowedMsgTypeToStr() string {
	s := ""
	for _, v := range allowedMsgTypes {
		s += v + " "
	}
	return s
}
