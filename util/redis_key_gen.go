package util

import com "ConfBackend/common"

func GenDistanceCacheKey(packetId, nodeId string) string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PacketPrefix + packetId + ":" + com.NodePrefix + nodeId
}

func GenDistanceQueryKey(packetId string) string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PacketPrefix + packetId + ":" + com.NodePrefix + "*"
}

func GenPacketTimelogPrefix() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.PacketTimelogPrefix
}

// GenLatestUpdatePackageTimeKey 生成最新更新的包的时间的key
// 为固定的"dc_:latest_package_time_"，值是一个时间戳
func GenLatestUpdatePackageTimeKey() string {
	return com.ProjectPref + ":" + com.DistanceCachePrefix + ":" + com.LatestPackageTimePrefix
}
