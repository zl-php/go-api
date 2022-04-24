package serializer

import (
	"longtu/model"
)

// CheckIn序列化器
type CheckIn struct {
	Openid    string `json:"openid"`
	DoorID    int    `json:"doorid"`
	Direction int    `json:"direction"`
	Qrstamp   string `json:"qrstamp"`
	Stamp     string `json:"stamp"`
	WorkID    string `json:"workid"`
	Name      string `json:"name"`
}

// BuildUser 序列化考勤数据
func BuildCheckIn(item model.CheckIn) CheckIn {
	return CheckIn{
		Name:      item.Name,
		Qrstamp:   item.Qrstamp,
		Stamp:     item.Stamp,
		Openid:    item.Openid,
		DoorID:    item.DoorID,
		Direction: item.Direction,
		WorkID:    item.WorkID,
	}
}

// BuildCheckIns 序列化考勤列表
func BuildCheckIns(items []model.CheckIn) (checkins []CheckIn) {
	for _, item := range items {
		checkin := BuildCheckIn(item)
		checkins = append(checkins, checkin)
	}
	return checkins
}
