package serializer

// 管理员获取的用户信息序列化器
type AdminUserInfo struct {
	UserID    uint
	Username  string
	CreatedAt int64
	UpdatedAt int64

	Status    int // 状态。0：正常，1：限制功能，2：封禁
	Authority int // 权限。0：管理员，1：普通用户

	MsgNumber    uint // 未读消息数
	OrderTotal   uint // 总订单数
	SendTotal    uint // 共发出订单数
	ReceiveTotal uint // 共接收订单数

	SendNumber    uint // 已发出（未被接收）订单数
	ReceiveNumber uint // 待接收订单数
}

// BuildAdminInfo 序列化管理员获取的用户信息
// func BuildAdminUserInfo(item model.AdminUserInfo) AdminUserInfo {
// 	return AdminUserInfo{
// 		UserID:    item.UserID,
// 		CreatedAt: item.CreatedAt.Unix(),
// 		UpdatedAt: item.UpdatedAt.Unix(),
// 		Username:  item.Username,

// 		Status:    item.Status,
// 		Authority: item.Authority,

// 		MsgNumber:    item.MsgNumber,
// 		OrderTotal:   item.OrderTotal,
// 		SendTotal:    item.SendTotal,
// 		ReceiveTotal: item.ReceiveTotal,

// 		SendNumber:    item.SendNumber,
// 		ReceiveNumber: item.ReceiveNumber,
// 	}
// }

// BuildAdminUserInfos 序列化管理员获取的用户信息列表
// func BuildAdminUserInfos(items []model.AdminUserInfo) (users []AdminUserInfo) {
// 	for _, item := range items {
// 		user := BuildAdminUserInfo(item)
// 		users = append(users, user)
// 	}
// 	return users
// }
