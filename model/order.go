package model

import (
	"time"

	"gorm.io/gorm"
)

// 订单模型
type Order struct {
	gorm.Model
	UserID uint

	// 发件人与收件人
	SName    string
	SPhone   string
	SAddress string
	RName    string
	RPhone   string
	RAddress string

	// 货物信息
	Type   string // 货物种类
	Weight int
	Volume int
	Value  int
	Urgent bool
	Note   string // 备注

	// 订单状态
	OrderID  string    // 订单号（14位随机大数字，string类型）
	Time     time.Time // 预计送达时间（由距离计算）
	Status   int       // 0：已完成，1：未完成
	Allocate int       // 分配的车辆，为0则未分配
	Rating   int       // 评价，为0则未评价
}
