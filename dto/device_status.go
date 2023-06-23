package dto

import (
	"github.com/golang-module/carbon/v2"
)

type DeviceInfo struct {
	DeviceName   string `json:"deviceName"`
	DeviceStatus string `json:"deviceStatus"`
	Time         any    `json:"time"`
}

func DeviceStatus() DeviceInfo {
	resp := DeviceInfo{
		DeviceName:   "Device Id",
		DeviceStatus: "好",
		Time:         carbon.Now().String(),
	}
	return resp
}
