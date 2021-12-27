package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AddressCreateService 地址创建服务
type AddressCreateService struct {
	UserID  uint   `form:"user_id" json:"user_id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

// AddressCreate 创建一条地址，并返回更新后的地址列表（按Pinned、创建时间排序）
func (service *AddressCreateService) AddressCreate() serializer.Response {
	address := model.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
		Pinned:  false,
	}
	code := constant.SUCCESS

	err := model.DB.Model(&model.Address{}).Create(&address).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var addresses []model.Address
	err = model.DB.Model(&model.Address{}).Where("user_id = ?", service.UserID).Order("pinned desc, updated_at desc").Find(&addresses).Error
	// Pinned = true 的结果在上
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
