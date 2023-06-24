package coord

import (
	"ConfBackend/dto"
	S "ConfBackend/services"
	"ConfBackend/task"
	"github.com/sirupsen/logrus"
)

// CalculateCoordinate 计算节点坐标
// b是一个map，key是termID，value是一个数组，数组中的元素是PTermDistanceDTO
// {NodeNo: str, Distance: float64(mm)}
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

	// 逐个处理每个人/pterm 的数据
	// ptermId, v 就是一个人的数据
	for ptermId, v := range b {
		// ptermId是人/pterm的ID, v是一个数组，数组中的元素是PTermDistanceDTO
		if len(v) < 4 {
			// 跳过不足4个距节点距离的pterm数据
			S.S.Logger.WithFields(logrus.Fields{
				"task":      "计算节点坐标时，pterm数据中有不足4个距节点距离的数据",
				"distCount": len(v),
				"ptermId":   ptermId,
			}).Info()
			continue
		}

		// 检查已知的节点的坐标是否都已经已知（小车已回传）

		// 计算坐标
		{
			nodeCoordArray := [][]float64{}
			distArray := []float64{}
			for _, dist := range v {
				nodeCoordArray = append(nodeCoordArray, []float64{nodeCoords[dist.NodeNo].X, nodeCoords[dist.NodeNo].Y, nodeCoords[dist.NodeNo].Z})
				distArray = append(distArray, dist.Distance)
			}
		}

	}
}
