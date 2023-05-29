package util

import "encoding/json"

func MarshalByte(v interface{}) []byte {
	res, _ := json.Marshal(v)
	return res
}

func MarshalString(v interface{}) string {
	res, _ := json.Marshal(v)
	return string(res)
}
