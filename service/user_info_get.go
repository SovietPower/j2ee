package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// UserInfoGetService 获取用户其它信息服务
type UserInfoGetService struct {
	UserID uint `form:"user_id" binding:"required"`
}

// UserInfoGet 获取用户其它信息
func (service *UserInfoGetService) UserInfoGet() serializer.Response {
	var info model.UserInfo
	code := constant.SUCCESS

	err := model.DB.Model(&model.UserInfo{}).Where("user_id = ?", service.UserID).Find(&info).Error
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
		Data:   serializer.BuildUserInfo(info),
	}
}
