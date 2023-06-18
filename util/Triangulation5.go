package util

import "fmt"

func getTargetPoint() {
	var knownPoints = [5][3]float64{
		{39.92330, 116.40350, 1.0}, // 故宫
		{39.90966, 116.43361, 3.6}, // 北京站
		{39.88766, 116.41766, 3.4}, // 天坛公园
		{39.90827, 116.36568, 3.4}, // 中央音乐学院
		{39.92237, 116.35876, 3.9}, // 月坛公园
	}

	var targetPoint = calculateTargetPoint(knownPoints) // 目标位置大概在天安门
	fmt.Println(targetPoint)
}

func calculateTargetPoint(knownPoints [5][3]float64) []float64 {
	var latitudeSum = 0.0
	var longitudeSum = 0.0
	var weightSum = 0.0
	for i := 0; i < len(knownPoints); i++ {
		var latitude = knownPoints[i][0]
		var longitude = knownPoints[i][1]
		var distance = knownPoints[i][2]

		var weight = 1.0 / distance

		latitudeSum += weight * latitude
		longitudeSum += weight * longitude
		weightSum += weight
	}

	var targetLatitude = latitudeSum / weightSum
	var targetLongitude = longitudeSum / weightSum

	var res = []float64{targetLatitude, targetLongitude}
	return res
}
