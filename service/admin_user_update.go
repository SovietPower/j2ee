package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AdminUserStatusService 设置用户状态服务
type AdminUserStatusService struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

// AdminUserStatus 设置用户状态
func (service *AdminUserStatusService) AdminUserStatus() serializer.Response {
	code := constant.SUCCESS

	err := model.DB.Model(&model.User{}).Where("id = ?", service.ID).Update("status", service.Status).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// err = model.DB.Model(&model.AdminUserInfo{}).Where("user_id = ?", service.ID).Update("status", service.Status).Error
	// if err != nil {
	// 	logging.Info(err)
	// 	code = constant.ERROR_DATABASE
	// 	return serializer.Response{
	// 		Status: code,
	// 		Msg:    constant.GetMsg(code),
	// 		Error:  err.Error(),
	// 	}
	// }

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
	}
}
