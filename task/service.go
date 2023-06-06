package task

import (
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
	"log"
)

func HaveValidUser(userUuid string) bool {
	// 先查询redis
	r := S.S.Redis

	// 是否有key
	hasKey := r.Keys(S.S.Context, util.GenMemberInfoNicknameKey(userUuid)).Val()

	// 有key，直接返回
	if len(hasKey) != 0 {
		log.Println("HaveValidUser: hasKey", userUuid)
		return true
	}

	// 没有key，查询数据库
	mber := getUserInMysql(userUuid)

	if len(mber) != 0 {
		// 有记录，写入redis
		// todo check if this is correct after changing k-v type
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
func LoadMysqlUserToRedis(mber ...model.Member) {
	if len(mber) == 0 {
		return
	}

	// todo change nickname type
	go func() {
		//r.HSet(S.S.Context, util.GenMemberInfoNicknameKey(mber.UUID), com.NicknameKey, mber.Nickname)
		//// 5 min expire
		//r.Expire(S.S.Context, util.GenMemberInfoNicknameKey(mber.UUID), time.Minute*5)

		//r.Set(S.S.Context, util.GenMemberInfoNicknameKey(mber.UUID), mber.Nickname, time.Minute*5)
		data := []string{}
		keys := []string{}
		for _, m := range mber {
			key := util.GenMemberInfoNicknameKey(m.UUID)
			data = append(data, key, m.Nickname)
			keys = append(keys, key)
		}
		p := S.S.Redis.TxPipeline()
		p.MSet(S.S.Context, data)

		{
			// 设置是否key过期，注释掉则不过期
			/*			for _, key := range keys {
						p.Expire(S.S.Context, key, time.Minute*5)
					}*/
		}

		exec, err := p.Exec(S.S.Context)
		if err != nil {
			return
		}
		log.Println("exec", exec)

		//r.MSet(S.S.Context, data)

	}()

}

// GetBatchNicknamesFromMysql 批量从mysql中获取用户信息。
// 不一定每个uuid都存在，需要判断
func GetBatchNicknamesFromMysql(uuids []string) map[string]interface{} {
	if len(uuids) == 0 {
		return map[string]interface{}{}
	}
	foundRes := make([]model.Member, 0)
	S.S.Mysql.Where("uuid in (?) and deleted_at is null", uuids).Find(&foundRes)
	foundMembers := make([]string, 0)
	for _, m := range foundRes {
		foundMembers = append(foundMembers, m.UUID)
	}
	notFoundMembers := util.Difference(uuids, foundMembers)

	LoadMysqlUserToRedis(foundRes...)

	///// padding returning results
	ret := make(map[string]interface{}, 0)
	for _, m := range foundRes {
		//ret = append(ret, map[string]interface{}{m.UUID: m.Nickname})
		ret[m.UUID] = m.Nickname
	}
	for _, m := range notFoundMembers {
		//ret = append(ret, map[string]interface{}{m: nil})
		ret[m] = nil
	}
	///// end padding

	return ret

}
