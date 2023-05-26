package view

import (
	com "ConfBackend/common"
	"ConfBackend/model"
	S "ConfBackend/services"
	"ConfBackend/util"
	"github.com/gin-gonic/gin"
)

type d struct {
	TotalPage int64 `json:"totalPage"`
	Result    any   `json:"result"`
}

func TestDb(c *gin.Context) {
	memberMgr := model.MemberMgr(S.S.Mysql)
	memberMgr.Omit("nickName")
	//q := c.Query("q")
	cur := c.Query("pg")
	size := c.Query("psize")
	page := model.Page{}
	// cur to int
	curNo := util.StringToInt64(cur)
	page.SetCurrent(curNo)
	page.SetSize(util.StringToInt64(size))
	//order := model.BuildDesc("id")
	//page.AddOrderItem(order)
	// ignore nickname
	Member, err := memberMgr.SelectPage(&page, memberMgr.WithNickname("a"))

	res := d{TotalPage: page.GetPages(), Result: Member.GetRecords()}

	if err != nil {
		return
	}
	com.OkD(c, res)
}
