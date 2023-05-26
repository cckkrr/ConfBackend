package main

import (
	"ConfBackend/server"
	S "ConfBackend/services"
)

func init() {
	S.InitServices()
}

// main 入口函数
func main() {

	// 单独的协程监听车的socket端口
	go server.StartListenHeroPort()

	// 协程，启动所有其他的定时/周期或其它线程任务
	go server.StartTimed()

	// 设置gin的运行模式 调试/生产
	//gin.SetMode(gin.ReleaseMode)
	// web 服务器
	server.StartApi()

}
