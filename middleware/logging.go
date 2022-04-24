package middleware

import (
	"github.com/gin-gonic/gin"
	"longtu/pkg/util"
	"time"
)

func NewLogger() gin.HandlerFunc {
	logger := util.Logger()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()               // 处理请求时间
		c.Next()                              // 处理下一步请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 该中间件执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqUri := c.Request.RequestURI        // 请求路由
		statusCode := c.Writer.Status()       // 状态码
		clientIP := c.ClientIP()              // 请求ip
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
