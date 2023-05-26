package com

import (
	"github.com/gin-gonic/gin"
)

const (
	ok       = 200
	error    = 400
	notFound = 404
)

type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OkM(c *gin.Context, msg string) {
	resp := CommonResp{
		Code: ok,
		Msg:  msg,
		Data: nil,
	}

	c.JSON(200, resp)
}

func Ok(c *gin.Context) {
	resp := CommonResp{
		Code: ok,
		Msg:  "ok",
		Data: nil,
	}

	c.JSON(200, resp)

}

func OkD(c *gin.Context, data any) {
	resp := CommonResp{
		Code: ok,
		Msg:  "ok",
		Data: data,
	}

	c.JSON(200, resp)

}

func OkF(c *gin.Context, msg string, data interface{}) {
	resp := CommonResp{
		Code: ok,
		Msg:  msg,
		Data: data,
	}
	c.JSON(200, resp)
}

func Error(c *gin.Context, msg string) {
	resp := CommonResp{
		Code: error,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(200, resp)
}

func ErrorD(c *gin.Context, msg string, data any) {
	resp := CommonResp{
		Code: error,
		Msg:  msg,
		Data: data,
	}
	c.JSON(200, resp)
}
