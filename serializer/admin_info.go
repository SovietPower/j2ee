package serializer

import "j2ee/model"

// AdminInfo 管理员其它信息序列化器
type AdminInfo struct {
	Username string `gorm:"unique"`

	MsgNumber  uint // 未读消息数
	OrderTotal uint // 总处理订单数
	UserTotal  uint // 总用户数
	TruckTotal uint // 总车辆数

	TruckAvailale  uint // 当前可用车辆数
	OrderUnhandled uint // 未处理订单数
}

// BuildAdminInfo 序列化管理员信息
func BuildAdminInfo(item model.AdminInfo) AdminInfo {
	return AdminInfo{
		Username: item.Username,

		MsgNumber:  item.MsgNumber,
		OrderTotal: item.OrderTotal,
		UserTotal:  item.UserTotal,
		TruckTotal: item.TruckTotal,

		TruckAvailale:  item.TruckAvailale,
		OrderUnhandled: item.OrderUnhandled,
	}
}
