package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goapi/cache"
	"goapi/model"
	"goapi/pkg/util"
	"os"
)

func Init() {

	// 从本地读取 .env
	godotenv.Load()

	// 设置gin的模式
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 读取翻译文件
	if err := LoadLocales("config/lang/zh-cn.yaml"); err != nil {
		util.Logger().Info(err) //日志
		panic(err)
	}

	// 连接数据库
	model.Mysql(os.Getenv("MYSQL_DSN"))

	// 连接redis
	cache.Redis(os.Getenv("REDIS_DB"))

}
