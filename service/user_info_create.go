package service

import (
	"j2ee/model"
)

type UserInfoCreateService struct {
	UserID   uint
	Username string
}

// UserInfoCreate 创建用户其它信息
func (service *UserInfoCreateService) UserInfoCreate() error {
	info := model.UserInfo{
		UserID:   service.UserID,
		Username: service.Username,

		MsgNumber:    0,
		OrderTotal:   0,
		SendTotal:    0,
		ReceiveTotal: 0,

		SendNumber:    0,
		ReceiveNumber: 0,
	}

	err := model.DB.Model(&model.UserInfo{}).Create(&info).Error

	return err
}
