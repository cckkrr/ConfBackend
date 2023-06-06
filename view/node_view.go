package view

import (
	S "ConfBackend/services"
	_ "ConfBackend/services"
	"ConfBackend/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type locInfoType map[string]float64

// UpdateLocation 更新位置
func UpdateLocation(c *gin.Context) {
	// 获取header中X-Node-Id值
	nodeId := c.GetHeader("X-Node-Id")

	// 获取header中X-Packet-Id值
	packetId := c.GetHeader("X-Packet-Id")

	locInfo := locInfoType{}

	// get request body content
	err := json.NewDecoder(c.Request.Body).Decode(&locInfo)
	if err != nil {
		log.Println(err)
	}
	setToRedis(c, nodeId, packetId, locInfo)

}

// setToRedis 将距离信息存入redis
func setToRedis(c *gin.Context, nodeId, packetId string, info locInfoType) {
	r := S.S.Redis

	slice := make([]string, 2*len(info))
	i := 0
	for k, v := range info {
		slice[i] = k
		slice[i+1] = util.Float64ToString(v)
		i += 2
	}

	// 设置距离信息，key编号：dc_:pkt_{包编号}:nd_{节点编号}
	// hash key: 节点编号，hash value: 距离

	p := r.Pipeline()
	p.HSet(c, util.GenDistanceCacheKey(packetId, nodeId), slice)

	//r.HSet(c, util.GenDistanceCacheKey(packetId, nodeId), slice)
	// get current timestamp
	timestamp := time.Now().UnixMilli()
	b := redis.Z{
		Score:  float64(timestamp),
		Member: packetId,
	}

	// 在dc_下设置一个有序集合，记录首个该编号上传的时间。key编号：dc_:pkt_tm_
	// zset value: 包编号，score: 时间戳
	//r.ZAddNX(c, util.GenPacketTimelogPrefix(), b)
	p.ZAddNX(c, util.GenPacketTimelogPrefix(), b)
	_, err := p.Exec(c)
	if err != nil {
		return
	}

}
