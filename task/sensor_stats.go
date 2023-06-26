package task

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/util"
	"encoding/json"
	"github.com/golang-module/carbon/v2"
)

func SetSensorStatsToRedis(nodeId string, light1, light2, voice int) {
	r := S.S.Redis
	b := dto.SensorStatsDTO{
		NodeId:     nodeId,
		UpdateTime: carbon.DateTime{carbon.Now()},
		Light1:     0,
		Light2:     0,
		Voice1:     0,
	}
	bstr, _ := json.Marshal(b)
	r.HSet(S.S.Context, util.GenNodeStatsKey(), nodeId, bstr)
}
