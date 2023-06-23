package server

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/util"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	pkgRetrieveStride int64 = 1
	ctx                     = context.Background()
)

func StartUpdateLocationTask() {

	for {
		// 将任务提交到协程池来执行
		//S.S.TaskPool.Submit(func() {
		//	updateTask(ctx)
		//})

		S.S.Logger.WithFields(logrus.Fields{
			"task": "定时执行一次更新位置任务",
		}).Infof("")
		updateTask(ctx)
		// 休眠一段时间，秒数由配置文件中的update_interval_in_second指定
		time.Sleep(time.Duration(S.S.Conf.Location.UpdateIntervalInSecond) * time.Second)
	}
}

func updateTask(c context.Context) {
	// 获取redis实例
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
			// float64 (time) UnixMilli()
			pkgTime := pkgItem.Score
			queryCountKey := util.GenDistanceQueryKey(pkgNo)
			pkgNodeList := r.Keys(c, queryCountKey).Val()
			if len(pkgNodeList) < 4 {

				// 少于4个数据，无法计算位置
				continue
			} else {

				// 如果算过，则也不用再算
				if pkgTime == util.StringToFloat64(lastUpdateTimestamp) {
					continue
				}

				// 有足够的数据，计算位置
				// 找到了有效的数据，设置为true
				hasFoundValid = true
				r.Set(c, util.GenLatestUpdatePackageTimeKey(), pkgTime, 0)
				//todo 计算位置
				S.S.Logger.WithFields(logrus.Fields{
					"task":         "找到可更新的数据，更新操作",
					"pkgNodeCount": len(pkgNodeList),
				}).Infof("最新包时间:%s", pkgTime)
				updateLocation(pkgNodeList, pkgTime)
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

func updateLocation(pkgNodeList []string, timeInUnixMilli float64) {
	r := S.S.Redis
	p := r.Pipeline()
	for _, nodeKey := range pkgNodeList {
		p.HGetAll(ctx, nodeKey)
	}
	res, err := p.Exec(S.S.Context)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"task":      "更新位置时，从redis中获取数据失败",
			"triedKeys": pkgNodeList,
		}).Error()
		return
	}

	b := make(map[string][]dto.NodeDistanceDTO)

	for _, cmd := range res {
		nodeNo := util.ParseNodeIdFromPktKey(cmd.(*redis.MapStringStringCmd).Args()[1].(string))
		nodeInfo := cmd.(*redis.MapStringStringCmd).Val()

		for k, v := range nodeInfo {
			// k is termid, v is distance in mm, if k not in b, add it in
			if _, ok := b[k]; !ok {
				b[k] = make([]dto.NodeDistanceDTO, 0)
			}
			b[k] = append(b[k], dto.NodeDistanceDTO{
				NodeNo:   nodeNo,
				Distance: util.StringToFloat64(v),
			})

		}

	}

	//todo 计算位置
	util.CalculateCoordinate(b, timeInUnixMilli)

}
