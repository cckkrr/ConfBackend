package task

import (
	com "ConfBackend/common"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
)

func HaveValidUser(userUuid string) bool {
	// 先查询redis
	r := S.S.Redis

	// 是否有key
	hasKey, _ := r.Exists(S.S.Context, util.GenMemberInfoKey(userUuid)).Result()

	// 有key，直接返回
	if hasKey == 1 {
		return true
	}

	// 没有key，查询数据库
	mber := GetUserInMysql(userUuid)

	if len(mber) != 0 {
		// 有记录，写入redis
		LoadMysqlUserToRedis(mber[0])
		return true
	} else {
		// 没有记录，返回false
		return false
	}

}

// GetUserInMysql 从mysql中获取用户信息
// 有则返回一个，没有则返回空数组
func GetUserInMysql(uuid string) []model.Member {
	d := S.S.Mysql
	members := make([]model.Member, 0)
	d.Where("uuid = ?", uuid).Find(&members)
	return members

}
func LoadMysqlUserToRedis(mber model.Member) {
	r := S.S.Redis
	go r.HSet(S.S.Context, util.GenMemberInfoKey(mber.UUID), com.NicknameKey, mber.Nickname)
}
