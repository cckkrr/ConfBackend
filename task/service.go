package task

import (
	com "ConfBackend/common"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
	"time"
)

func HaveValidUser(userUuid string) bool {
	// 先查询redis
	r := S.S.Redis

	// 是否有key
	hasKey, _ := r.Keys(S.S.Context, util.GenMemberInfoKey(userUuid)).Result()

	// 有key，直接返回
	if len(hasKey) != 0 {
		return true
	}

	// 没有key，查询数据库
	mber := getUserInMysql(userUuid)

	if len(mber) != 0 {
		// 有记录，写入redis
		go LoadMysqlUserToRedis(mber[0])
		return true
	} else {
		// 没有记录，返回false
		return false
	}

}

// getUserInMysql 从mysql中获取用户信息
// 有则返回一个，没有则返回空数组
func getUserInMysql(uuid string) []model.Member {
	d := S.S.Mysql
	members := make([]model.Member, 0)
	d.Where("uuid = ? and deleted_at is null", uuid).Find(&members)
	return members

}
func LoadMysqlUserToRedis(mber model.Member) {
	r := S.S.Redis
	go func() {
		r.HSet(S.S.Context, util.GenMemberInfoKey(mber.UUID), com.NicknameKey, mber.Nickname)
		// 5 min expire
		r.Expire(S.S.Context, util.GenMemberInfoKey(mber.UUID), time.Minute*5)
	}()

}
