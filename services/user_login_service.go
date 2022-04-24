package services

import (
	"goapi/model"
	"goapi/pkg/util"
	"goapi/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// 用户登录
func (service *UserLoginService) Login() serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 下发token
	token, err := util.GenerateToken(user.ID, service.UserName)
	if err != nil {
		return serializer.ParamErr("生成token失败", err)
	}

	// 返回数据
	return serializer.Response{
		ErrCode: 0,
		Message: "success",
		Data:    serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}

}
