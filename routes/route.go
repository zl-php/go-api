package routes

import (
	"github.com/gin-gonic/gin"
	"longtu/api"
	"longtu/middleware"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// 中间件
	r.Use(middleware.NewLogger(), middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("checkin/get", api.GetCheckinList)
		v1.GET("redis/test", api.RedisTest)
		v1.GET("jwt/create", api.CreateToken)
		v1.GET("jwt/update", api.UpdateToken)

		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("jwt/check", api.CheckToken)
		}
	}

	return r
}
