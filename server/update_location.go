package server

import (
	S "ConfBackend/services"
	"ConfBackend/util"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	pkgRetrieveStride int64 = 10
	ctx                     = context.Background()
)

func StartUpdateLocationTask() {

	for {
		// 将任务提交到协程池来执行
		//S.S.TaskPool.Submit(func() {
		//	updateTask(ctx)
		//})

		updateTask(ctx)
		// 休眠一段时间，秒数由配置文件中的update_interval_in_second指定
		time.Sleep(time.Duration(S.S.Conf.Location.UpdateIntervalInSecond) * time.Second)
	}
}

func updateTask(c context.Context) {
	// 获取服务示例
	r := S.S.Redis
	lastUpdateTimestamp := r.Get(c, util.GenLatestUpdatePackageTimeKey()).Val()
	if lastUpdateTimestamp == "" {
		r.Set(c, util.GenLatestUpdatePackageTimeKey(), "0", 0)
		lastUpdateTimestamp = "0"
	}

	//p := S.S.TaskPool

	queryRedisTimeLogKey := util.GenPacketTimelogPrefix()

	var offset int64 = 0

	for {
		pkgTimes := r.ZRevRangeByScoreWithScores(c, queryRedisTimeLogKey, &redis.ZRangeBy{
			Min:    lastUpdateTimestamp,
			Max:    "+inf",
			Offset: offset,
			Count:  pkgRetrieveStride,
		}).Val()

		// len = 0 means no data in redis
		if len(pkgTimes) == 0 {
			break
		}

		hasFoundValid := false

		for _, pkgItem := range pkgTimes {
			pkgNo := pkgItem.Member.(string)
			pkgTime := pkgItem.Score
			queryCountKey := util.GenDistanceQueryKey(pkgNo)
			pkgNodeList := r.Keys(c, queryCountKey).Val()
			if len(pkgNodeList) < 3 {

				// 少于3个数据，无法计算位置
				continue
			} else {

				// 有足够的数据，计算位置
				// 找到了有效的数据，设置为true
				hasFoundValid = true
				r.Set(c, util.GenLatestUpdatePackageTimeKey(), pkgTime, 0)
				//todo 计算位置
				updateLocation(pkgNodeList)
				break
			}

		}

		if hasFoundValid {
			break
		} else {
			offset += pkgRetrieveStride
		}

	}

}

func updateLocation(pkgNodeList []string) {

}
