package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AddressDeleteService 地址删除服务
type AddressDeleteService struct {
	AddressID uint `json:"address_id" binding:"required"`
}

// AddressDelete 删除地址。只需在前端删除，无需返回更新后的列表
func (service *AddressDeleteService) AddressDelete() serializer.Response {
	var address model.Address
	code := constant.SUCCESS

	err := model.DB.Where("id = ?", service.AddressID).Find(&address).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&address).Error
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
	}
}
