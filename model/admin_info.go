package model

import (
	"gorm.io/gorm"
)

// 管理员其它信息（常变化的）
type AdminInfo struct {
	gorm.Model
	UserID   uint   `gorm:"unique"`
	Username string `gorm:"unique"`

	MsgNumber  uint // 未读消息数
	OrderTotal uint // 总处理订单数
	UserTotal  uint // 总用户数
	TruckTotal uint // 总车辆数

	TruckAvailale  uint // 当前可用车辆数
	OrderUnhandled uint // 未处理订单数
}
