package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AdminInfoGetService 获取管理员其它信息服务
type AdminInfoGetService struct {
	UserID uint `form:"user_id" binding:"required"`
}

// AdminInfoGet 获取管理员其它信息
func (service *AdminInfoGetService) AdminInfoGet() serializer.Response {
	var info model.AdminInfo
	code := constant.SUCCESS

	err := model.DB.Model(&model.AdminInfo{}).Where("user_id = ?", service.UserID).Find(&info).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Data:   serializer.BuildAdminInfo(info),
	}
}
