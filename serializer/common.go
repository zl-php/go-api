package serializer

import (
	"github.com/gin-gonic/gin"
)

// 错误码定义
const (
	//CodeParamErr 各种参数错误
	CodeParamErr = 40001
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
)

// Response json序列化
type Response struct {
	ErrCode int         `json:"errcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

//TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// 带有总数的Data结构
type ListData struct {
	List  interface{} `json:"list"`
	Total uint        `json:"total"`
}

// 带有总数的列表构建器
func ListTotalData(items interface{}, total uint) Response {
	return Response{
		ErrCode: 0,
		Data: ListData{
			List:  items,
			Total: total,
		},
		Message: "success",
	}
}

// ParamErr 参数错误
func ParamErr(message string, err error) Response {
	if message == "" {
		message = "参数错误"
	}
	return Err(CodeParamErr, message, err)
}

// DBErr 数据库操作失败
func DBErr(message string, err error) Response {
	if message == "" {
		message = "数据库操作失败"
	}
	return Err(CodeDBError, message, err)
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		ErrCode: errCode,
		Message: msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}
