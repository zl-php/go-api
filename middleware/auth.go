package middleware

import (
	"github.com/gin-gonic/gin"
	"goapi/pkg/util"
	"goapi/serializer"
	"strings"
	"time"
)

// JWT 中间件
func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(200, gin.H{"errcode": serializer.CodeParamErr, "message": "未登录，请先登录"})
			c.Abort() //阻止执行
			return
		}

		tokenSlice := strings.SplitN(token, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			c.JSON(200, gin.H{"errcode": serializer.CodeParamErr, "message": "token格式错误"})
			c.Abort()
			return
		}

		claims, err := util.ParseToken(tokenSlice[1])
		if err != nil {
			c.JSON(200, gin.H{"errcode": serializer.CodeParamErr, "message": "用户信息解析失败，请重新登录"})
			c.Abort()
			return
		}

		//token超时
		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(200, gin.H{"errcode": serializer.CodeParamErr, "message": "登录状态已失效，请重新登录"})
			c.Abort()
			return
		}

		// 此处待改成根据用户id获取用户信息并set
		c.Set("user_id", claims.UserId)
		c.Set("username", claims.UserName)

		c.Next()
	}
}
