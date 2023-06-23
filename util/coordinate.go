package util

import (
	"ConfBackend/dto"
)

func CalculateCoordinate(b map[string][]dto.NodeDistanceDTO, timeInUnixMilli float64) {
	// 获取已知节点的坐标

	for _, v := range b {
		if len(v) < 4 {
			continue
		}

	}
}
