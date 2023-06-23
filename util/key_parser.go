package util

import (
	com "ConfBackend/common"
	"strings"
)

// ParseNodeIdFromPktKey parses the node id from the key of a packet
func ParseNodeIdFromPktKey(key string) string {
	// first split by ":", and find the slice that starts with "nd_", return the second element
	// e.g. "pkt:nd_1:1234567890" -> "1"
	// e.g. "pkt:nd_1:1234567890:1234567890" -> "1"

	// split by ":"
	keySlice := strings.Split(key, ":")
	// find the slice that starts with "nd_"
	for _, s := range keySlice {
		if strings.HasPrefix(s, "nd_") {
			return strings.TrimPrefix(s, com.NodePrefix)
		}
	}
	return ""
}
