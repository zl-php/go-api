package model

// User 用户模型
type CheckIn struct {
	ID         uint
	Openid     string
	SID        int
	Controller string
	DoorID     int
	ReaderID   int
	Direction  int
	OID        int
	Qrstamp    string
	Stamp      string
	WorkID     string
	Name       string
}

// 自定义表名
func (CheckIn) TableName() string {
	return "kaoqin31day"
}
