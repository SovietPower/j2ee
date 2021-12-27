package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AddressGetService 获取地址服务
type AddressGetService struct {
}

// AddressGet 获取地址
func (service *AddressGetService) AddressGet(id string) serializer.Response {
	var addresses []model.Address
	code := constant.SUCCESS

	err := model.DB.Model(&model.Address{}).Where("user_id = ?", id).Order("pinned desc, updated_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}
