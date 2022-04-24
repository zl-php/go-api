package serializer

import "goapi/model"

// User 用户序列化器
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		Nickname: user.Nickname,
		Status:   user.Status,
	}
}

// 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Message: "success",
		Data:    BuildUser(user),
	}
}
