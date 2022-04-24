package main

import (
	"goapi/config"
	"goapi/routes"
	"os"
)

func main() {

	// 加载项目，包含日志、mysql、redis等
	config.Init()

	// 加载路由 启动8888端口
	r := routes.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
