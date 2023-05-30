package archive

import (
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
)

func ArchiveSingleMsg(msg model.ImMessage) error {
	//todo 保存消息，redis和mysql

	onlyToRedis := false

	if onlyToRedis {
		// 只保存到redis
		rKey := util.GenMsgListStaticKey()
		_, err := S.S.Redis.HSet(S.S.Context, rKey, msg.UUID, util.MarshalString(msg)).Result()
		if err != nil {
			return err
		}
	} else {
		// 保存到mysql
		S.S.Mysql.Create(&msg)

	}

	return nil
}
