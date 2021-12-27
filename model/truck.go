package model

import (
	"gorm.io/gorm"
)

type Truck struct {
	gorm.Model

	Status    int    // 货车状态。0：正常，1：不可用
	Allocated int    // 分配状态。0：未分配，1：已分配
	OrderID   string // 分配的订单号（如果已分配）

	// 保存部分订单信息，用于展示？
	SAddress string // 发货地
	RAddress string // 收货地
	Type     string // 货物种类
}
