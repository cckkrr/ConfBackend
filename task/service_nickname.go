package task

import (
	S "ConfBackend/services"
	"ConfBackend/util"
)

type nickNameBody struct {
	UserUUID string `json:"userUuid"`
	NickName string `json:"nickName"`
}

func GetNickNames(uuids []string) map[string]interface{} {
	r := S.S.Redis
	//gen redis keys
	keys := make([]string, 0)
	for _, uuid := range uuids {
		keys = append(keys, util.GenMemberInfoNicknameKey(uuid))
	}
	// 批量获取keys的hash中的nickname值
	validRes := make(map[string]interface{}, 0)

	// 首先查询是否能够从redis中获取
	res := r.MGet(S.S.Context, keys...).Val()

	nilUuids := []string{}
	for i, v := range res {
		if v == nil {
			nilUuids = append(nilUuids, uuids[i])
		}
		if v != nil {
			validRes[uuids[i]] = v.(string)
			//validRes = append(validRes, map[string]interface{}{keys[i]: v.(string)})
		}
	}

	// 从mysql中获取memberBody
	mysqlFound := GetBatchNicknamesFromMysql(nilUuids)
	// 拼接mysql和redis的结果
	//validRes = append(validRes, mysqlFound...)
	for k, v := range mysqlFound {
		//validRes = append(validRes, v)
		validRes[k] = v
	}
	return validRes
}
