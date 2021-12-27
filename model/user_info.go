package model

import (
	"gorm.io/gorm"
)

// 用户其它信息（常变化的）
type UserInfo struct {
	gorm.Model
	UserID   uint   `gorm:"unique"`
	Username string `gorm:"unique"`

	MsgNumber    uint // 未读消息数
	OrderTotal   uint // 总订单数
	SendTotal    uint // 共发出订单数
	ReceiveTotal uint // 共接收订单数

	SendNumber    uint // 已发出（未被接收）订单数
	ReceiveNumber uint // 待接收订单数
}
