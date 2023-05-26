package server

func StartTimed() {

	// 定时更新位置
	go StartUpdateLocationTask()

}
