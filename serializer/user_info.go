package serializer

import "j2ee/model"

// UserInfo 用户其它信息序列化器
type UserInfo struct {
	Username string `gorm:"unique"`

	MsgNumber    uint // 未读消息数
	OrderTotal   uint // 总订单数
	SendTotal    uint // 共发出订单数
	ReceiveTotal uint // 共接收订单数

	SendNumber    uint // 已发出（未被接收）订单数
	ReceiveNumber uint // 待接收订单数
}

// BuildUserInfo 序列化用户信息
func BuildUserInfo(item model.UserInfo) UserInfo {
	return UserInfo{
		Username: item.Username,

		MsgNumber:    item.MsgNumber,
		OrderTotal:   item.OrderTotal,
		SendTotal:    item.SendTotal,
		ReceiveTotal: item.ReceiveTotal,

		SendNumber:    item.SendNumber,
		ReceiveNumber: item.ReceiveNumber,
	}
}
