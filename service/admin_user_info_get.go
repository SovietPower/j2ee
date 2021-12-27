package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AdminUserGetService 管理员获取所有用户信息服务
type AdminUserGetService struct {
	// Type int // 查询的用户状态
}

// AdminUserGet 管理员获取所有用户信息
func (service *AdminUserGetService) AdminUserGet() serializer.Response {
	var adminInfos []serializer.AdminUserInfo
	code := constant.SUCCESS

	var users []model.User
	err := model.DB.Model(&model.User{}).Find(&users).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var userInfos []model.UserInfo
	err = model.DB.Model(&model.UserInfo{}).Find(&userInfos).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	for i, user := range users {
		userInfo := userInfos[i]
		adminInfo := serializer.AdminUserInfo{
			UserID:    user.ID,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
			Username:  user.Username,

			Status:    user.Status,
			Authority: user.Authority,

			MsgNumber:    userInfo.MsgNumber,
			OrderTotal:   userInfo.OrderTotal,
			SendTotal:    userInfo.SendTotal,
			ReceiveTotal: userInfo.ReceiveTotal,

			SendNumber:    userInfo.SendNumber,
			ReceiveNumber: userInfo.ReceiveNumber,
		}
		adminInfos = append(adminInfos, adminInfo)
	}

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Data:   adminInfos,
	}
}
