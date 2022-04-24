package api

import (
	"github.com/gin-gonic/gin"
	"goapi/model"
	"goapi/serializer"
	"goapi/services"
)

// 用户注册
func UserRegister(c *gin.Context) {
	var service services.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 用户登录
func UserLogin(c *gin.Context) {
	var service services.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取用户信息
func GetUserInfo(c *gin.Context) {
	if userId, _ := c.Get("user_id"); userId != nil {
		user, err := model.GetUser(userId)
		if err == nil {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	}
}
