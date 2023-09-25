package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	//初始化配置
	global.Gva.GvaConfigs = initialize.Config()
	//初始化日志Logger
	global.Gva.GvaLogger = initialize.Logger()
	//启动server
	core.StartServer()

}
