package services

import (
	"goapi/model"
	"goapi/serializer"
)

// 管理用户修改的服务
type UserUpdateService struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
}

// valid 验证表单
func (service *UserUpdateService) valid(userId interface{}) *serializer.Response {

	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			ErrCode: serializer.CodeParamErr,
			Message: "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			ErrCode: serializer.CodeParamErr,
			Message: "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserUpdateService) Update(userId interface{}) serializer.Response {
	var user model.User

	// 表单验证
	if err := service.valid(userId); err != nil {
		return *err
	}

	if err := model.DB.First(&user, userId).Error; err != nil {
		return serializer.ParamErr("用户不存在", err)
	}

	user.Nickname = service.Nickname
	user.UserName = service.UserName

	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.DBErr("", err)
	}

	// 返回数据
	return serializer.Response{
		ErrCode: 0,
		Message: "success",
		Data:    serializer.BuildUser(user),
	}

}
