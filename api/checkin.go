package api

import (
	"github.com/gin-gonic/gin"
	"longtu/services"
)

// 获取考勤数据
func GetCheckinList(ctx *gin.Context) {
	var service services.SearchCheckInService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.GetCheckinDataList()
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, ErrorResponse(err))
	}
}
