package unread

import (
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
	"github.com/redis/go-redis/v9"
	"log"
)

// SaveSingleUnreadMsg 保存单挑消息到用户的inbox
func SaveSingleUnreadMsg(msg model.ImMessage) {
	log.Println("SaveSingleUnreadMsg", msg)
	userInboxRedisKey := util.GenUserInboxKey(msg.ToEntityUUID)
	zobj := redis.Z{
		Score:  float64(msg.CreatedAt.UnixNano()),
		Member: msg.UUID,
	}

	_, err := S.S.Redis.ZAdd(S.S.Context, userInboxRedisKey, zobj).Result()

	if err != nil {
		log.Println("save single msg FAILED", err)
		return
	}
}

func GetUnreadMessage(registerUserUuid string) []model.ImMessage {
	unreadKey := util.GenUserInboxKey(registerUserUuid)
	msgList, _ := S.S.Redis.ZRange(S.S.Context, unreadKey, 0, -1).Result()
	// init msg bodies to hit mysql
	msgBodies := make([]model.ImMessage, 0)
	// 批量填充msg的uuid，然后批量查询mysql
	for _, msgUuid := range msgList {
		msg := model.ImMessage{}
		msg.UUID = msgUuid
		msgBodies = append(msgBodies, msg)
	}
	// 批量查询mysql
	S.S.Mysql.Where("uuid in ?", msgList).Find(&msgBodies)

	// 清空redis
	S.S.Redis.Del(S.S.Context, unreadKey)
	// 返回
	return msgBodies
}
