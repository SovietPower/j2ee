package service

// import (
// 	"j2ee/model"
// 	"j2ee/serializer"
// )

// type AdminUserInfoCreateService struct {
// 	UserID   uint
// 	Username string
// }

// AdminUserInfoCreate 创建用于展示给管理员的用户信息
// func (service *AdminUserInfoCreateService) AdminUserInfoCreate() error {
// 	info := serializer.AdminUserInfo{
// 		UserID:   service.UserID,
// 		Username: service.Username,

// 		Status:    model.StatusNormal,
// 		Authority: model.AuthorityUser,

// 		MsgNumber:    0,
// 		OrderTotal:   0,
// 		SendTotal:    0,
// 		ReceiveTotal: 0,

// 		SendNumber:    0,
// 		ReceiveNumber: 0,
// 	}

// 	err := model.DB.Model(&model.AdminUserInfo{}).Create(&info).Error

// 	return err
// }
