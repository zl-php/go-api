package services

import (
	"longtu/model"
	"longtu/pkg/util"
	"longtu/serializer"
	"strconv"
)

// 全局变量，声明可调用该接口的appid和签名秘钥
var app = map[string]string{
	"oa": "8JAde9mdLJLhpxQSPGFhiV6gEjUoq7PPEeIlJO3B8qZXRneNl0XqRK0ChdFV4mNU",
}

// SearchCheckInService 考勤查询的服务
type SearchCheckInService struct {
	Appid     string `form:"appid" json:"appid" binding:"required"`
	Starttime string `form:"starttime" json:"starttime" binding:"required"`
	Endtime   string `form:"endtime" json:"endtime" binding:"required"`
	Workid    string `form:"workid" json:"workid" binding:"required"`
	Sign      string `form:"sign" json:"sign" binding:"required"`
}

// 获取考勤数据 list
func (service *SearchCheckInService) GetCheckinDataList() serializer.Response {

	checkins := []model.CheckIn{}

	// 验证appid
	secretKey, ok := app[service.Appid]
	if !ok {
		return serializer.ParamErr("无效的参数 -- [appid]", nil)
	}

	// 拼装签名参数
	params := map[string]string{
		"appid":     service.Appid,
		"starttime": service.Starttime,
		"endtime":   service.Endtime,
		"workid":    service.Workid,
		"sign":      service.Sign,
	}

	// 生成签名
	sign := util.GenParmSign(params, secretKey)

	// 效验签名
	if service.Sign != sign {
		return serializer.ParamErr("签名错误", nil)
	}

	// 判断查询时间范围，将字符串类型的时间戳转为int64类型
	start, _ := strconv.ParseInt(service.Starttime, 10, 64)
	end, _ := strconv.ParseInt(service.Endtime, 10, 64)
	if util.GetDiffDaysBySecond(start, end) > 7 {
		return serializer.ParamErr("查询开始时间和结束时间不能超过7天", nil)
	}

	// 查询考勤数据库
	if service.Workid == "all" {
		if err := model.DB.Where("stamp BETWEEN ? AND ?", service.Starttime, service.Endtime).Find(&checkins).Error; err != nil {
			return serializer.ParamErr("该查询条件内没有考勤数据", err)
		}
	} else {
		if err := model.DB.Where("WorkID = ? AND stamp BETWEEN ? AND ?", service.Workid, service.Starttime, service.Endtime).Find(&checkins).Error; err != nil {
			return serializer.ParamErr("该查询条件内没有考勤数据", err)
		}
	}

	// 返回考勤列表数据
	data := serializer.BuildCheckIns(checkins)

	return serializer.Response{
		ErrCode: 0,
		Message: "success",
		Data:    data,
	}

}
