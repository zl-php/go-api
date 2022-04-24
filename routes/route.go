package routes

import (
	"github.com/gin-gonic/gin"
	"goapi/api"
	"goapi/middleware"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// 中间件
	r.Use(middleware.NewLogger(), middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.GET("redis/test", api.RedisTest)

		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/info", api.GetUserInfo)
			auth.POST("user/update", api.UserUpdate)
		}
	}

	return r
}
