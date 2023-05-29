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

// GenUserInboxKey 用于获取一个用户的inbox的key，输入user uuid，返回key
func GenUserInboxKey(useruuid string) string {
	return com.ImInboxWholePrefix + ":" + useruuid
}

// GenMsgListStaticKey 返回一个静态的消息列表list，即 tr_:im_:msg_
func GenMsgListStaticKey() string {
	return com.MsgListStaticKey
}

func GenMemberInfoNicknameKey(uuid string) string {
	return com.ProjectPref + ":" + com.MemberStr + ":" + com.NicknameKey + ":" + uuid
}
