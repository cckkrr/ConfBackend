package task

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/util"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strings"
)

func SetNodeCoord(inferredColor string, x, y, z float64) {
	nodeInfo := S.S.Conf.Node.NodeInfo

	color2NodeId := map[string]string{}
	for _, node := range nodeInfo {
		res := strings.Split(node, "|")
		color2NodeId[strings.ToLower(res[0])] = res[1]
	}

	if !util.ContainKey(inferredColor, color2NodeId, false) {
		S.S.Logger.WithFields(logrus.Fields{
			"inferredColor": inferredColor,
			"x":             x,
			"y":             y,
			"z":             z,
		}).Errorf("操作失败：conf中未定义此颜色节点")
		return
	}

	targetNodeId := color2NodeId[strings.ToLower(inferredColor)]
	S.S.Logger.WithFields(logrus.Fields{
		"inferredColor": inferredColor,
		"x":             x,
		"y":             y,
		"z":             z,
	}).Infof("准备设置节点坐标")

	setNodeCoordToRedis(targetNodeId, x, y, z)

}

func setNodeCoordToRedis(nodeId string, x, y, z float64) {
	r := S.S.Redis

	// check if node coord already exists in redis
	{
		e := r.HExists(S.S.Context, util.GenNodeCoordKey(), nodeId).Val()
		if e {
			S.S.Logger.WithFields(logrus.Fields{
				"nodeId": nodeId,
			}).Warn("警告：尝试修改坐标的节点，节点坐标已存在。操作仍会继续。")

		}
	}
	rkey := util.GenNodeCoordKey()
	nodeBody := dto.NodeCoordDTO{
		NodeId: nodeId,
		X:      x,
		Y:      y,
		Z:      z,
	}
	nbstr, err := json.Marshal(nodeBody)
	if err != nil {
		return
	}

	// set to redis
	r.HSet(S.S.Context, rkey, nodeId, nbstr)
}

// GetNodeCoord get node coord from redis
// return map[nodeId]nodeCoord
func GetNodeCoord() map[string]dto.NodeCoordDTO {
	r := S.S.Redis
	rkey := util.GenNodeCoordKey()

	res := r.HGetAll(S.S.Context, rkey).Val()

	nodeCoord := map[string]dto.NodeCoordDTO{}

	for _, v := range res {
		var nodeBody dto.NodeCoordDTO
		err := json.Unmarshal([]byte(v), &nodeBody)
		if err != nil {
			continue
		}
		nodeCoord[nodeBody.NodeId] = nodeBody
	}

	return nodeCoord
}
