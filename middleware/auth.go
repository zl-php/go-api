package middleware

import (
	"github.com/gin-gonic/gin"
	"longtu/pkg/util"
	"net/http"
	"strings"
	"time"
)

// JWT 中间件
func AuthRequired() gin.HandlerFunc {

	return func(cxt *gin.Context) {
		token := cxt.GetHeader("Authorization")

		if token == "" {
			cxt.JSON(http.StatusOK, gin.H{"errcode": 40001, "message": "未登录，请先登录"})
			cxt.Abort()
			return
		}

		tokenSlice := strings.SplitN(token, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			cxt.JSON(http.StatusOK, gin.H{"errcode": 40001, "message": "token格式错误"})
			cxt.Abort() //阻止执行
			return
		}

		claims, err := util.ParseToken(tokenSlice[1])
		if err != nil {
			cxt.JSON(http.StatusOK, gin.H{"errcode": 40001, "message": "用户信息解析失败，请重新登录"})
			cxt.Abort() //阻止执行
			return
		}

		//token超时
		if time.Now().Unix() > claims.ExpiresAt {
			cxt.JSON(http.StatusOK, gin.H{"errcode": 40001, "message": "登录状态已失效，请重新登录"})
			cxt.Abort() //阻止执行
			return
		}

		// 此处待改成根据用户id获取用户信息并set
		cxt.Set("user_id", claims.UserId)
		cxt.Set("username", claims.UserName)

		cxt.Next()
	}
}
