package coord

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/task"
	"ConfBackend/util"
	"encoding/json"
	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
)

// CalculateCoordinate 计算节点坐标
// b是一个map，key是termID，value是一个数组，数组中的元素是PTermDistanceDTO
// {NodeId: str, Distance: float64(mm)}
// 要求：conf定义的color｜nodeId和updateLocation api中的nodeId严格一致
func CalculateCoordinate(b map[string][]dto.PTermDistanceDTO, timeInUnixMilli float64) {
	// 获取已知节点的坐标
	nodeCoords := task.GetNodeCoord()

	// 检查已知节点个数是否符合要求
	if len(nodeCoords) < 4 {
		S.S.Logger.WithFields(logrus.Fields{
			"task":      "计算节点坐标时，已知节点个数不足。需等待小车返回足够的已知节点坐标。",
			"nodeCount": len(nodeCoords),
		}).Info()

		return
	}

	// 获取已知节点的ID slice
	knownNodeIds := []string{}
	for k, _ := range nodeCoords {
		knownNodeIds = append(knownNodeIds, k)
	}

	toSaveToRedis := make([]dto.PTermCalcedCoordDTO, 0)

	// 逐个处理每个人/pterm 的数据
	// ptermId, distanceDTOS 就是一个人的数据
	for ptermId, distanceDTOS := range b {
		// ptermId是人/pterm的ID, v是一个数组，数组中的元素是PTermDistanceDTO
		if len(distanceDTOS) < 4 {
			// 跳过不足4个距节点距离的pterm数据
			S.S.Logger.WithFields(logrus.Fields{
				"task":      "计算节点坐标时，pterm数据中有不足4个距节点距离的数据",
				"distCount": len(distanceDTOS),
				"ptermId":   ptermId,
			}).Info()
			continue
		}

		///// 检查已知的节点的坐标是否都已经已知（小车已回传）
		knownDistNodeIds := []string{}
		for _, dist := range distanceDTOS {
			knownDistNodeIds = append(knownDistNodeIds, dist.NodeId)
		}

		// nodeIdsInters 是已知坐标node和入参中已知距离pterm nodeId交集
		nodeIdsInters := util.Intersection(knownNodeIds, knownDistNodeIds)

		if len(nodeIdsInters) < 4 {
			// 跳过不足4个已知节点的pterm数据
			S.S.Logger.WithFields(logrus.Fields{
				"task":               "计算节点坐标时，pterm距离数据和已知节点数据的交集不足4个",
				"knownNodeCount":     len(knownNodeIds),
				"knownDistNodeCount": len(knownDistNodeIds),
			}).Info()
			continue
		}
		// todo 修改下面，需要根据intersection的结果获取已知，进行拼接
		// todo for _, dist := range nodeIdsInters { 之类的

		// 计算坐标，此处距离/1000，单位转换为m

		nodeCoordArray := [][]float64{}
		distArrayInMeter := [][]float64{}
		for _, eachNodeIdInInters := range nodeIdsInters {
			for _, distDTO := range distanceDTOS {
				if distDTO.NodeId == eachNodeIdInInters {
					nodeCoordArray = append(nodeCoordArray,
						[]float64{nodeCoords[eachNodeIdInInters].X,
							nodeCoords[eachNodeIdInInters].Y,
							nodeCoords[eachNodeIdInInters].Z})
					distArrayInMeter = append(distArrayInMeter, []float64{distDTO.Distance / 1000})
				}
			}
		}

		// Slices padded. Then calculate the coords.
		coord, err := doCalcCoord(ptermId, nodeCoordArray, distArrayInMeter, timeInUnixMilli)
		if err != nil {

		} else {
			toSaveToRedis = append(toSaveToRedis, coord)
		}

	}

	// save toSaveToRedis to Redis
	{
		r := S.S.Redis
		pipe := r.Pipeline()
		rkey := util.GenCalcedPTermCoordKey()
		for _, coordDTO := range toSaveToRedis {
			hkey := coordDTO.PTermId
			hvalue, _ := json.Marshal(coordDTO)
			pipe.HSet(S.S.Context, rkey, hkey, string(hvalue))
		}
		_, err := pipe.Exec(S.S.Context)
		if err != nil {
			S.S.Logger.WithFields(logrus.Fields{
				"task": "计算节点坐标时，保存计算结果到Redis时出错",
			}).Error(err)
		}
	}

}

// doCalcCoord 计算坐标
// 传进来的坐标和距离必须是匹配好的，下标对应.
// 距离单位必须已经转成单位 米
func doCalcCoord(termId string, nodeCoordArray, distArrayInMeter [][]float64, timeInUnixMilli float64) (dto.PTermCalcedCoordDTO, error) {
	x, y, z, err := doCalcCoordMath(nodeCoordArray, distArrayInMeter)
	if err != nil {
		S.S.Logger.WithFields(logrus.Fields{
			"task":   "计算节点坐标时，计算出错-doCalcCoord-doCalcCoordMath",
			"termId": termId,
		}).Error(err)
		return dto.PTermCalcedCoordDTO{}, err
	}

	body := dto.PTermCalcedCoordDTO{
		PTermId:    termId,
		UpdateTime: carbon.DateTimeMilli{Carbon: carbon.CreateFromTimestampMilli(int64(timeInUnixMilli))},
		X:          x,
		Y:          y,
		Z:          z,
	}
	return body, nil

}
