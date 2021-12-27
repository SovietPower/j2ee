package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AddressUpdateService 地址修改服务
type AddressUpdateService struct {
	ID      uint   `form:"id" json:"id"`
	UserID  uint   `form:"user_id" json:"user_id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

// AddressUpdate 修改地址信息，并返回更新后的地址列表
func (service *AddressUpdateService) AddressUpdate() serializer.Response {
	address := model.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	// address.ID = service.ID // gorm.Model中的赋值更新
	code := constant.SUCCESS

	// 更新多个值，用Updates（用结构体，不更新零值，但这样更新不了Pinned bool）
	var tmp model.Address
	err := model.DB.Where("id = ?", service.ID).First(&tmp).Updates(address).Error
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
	err = model.DB.Where("user_id = ?", service.UserID).Order("pinned desc, updated_at desc").Find(&addresses).Error
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

// AddressUpdatePinService 地址置顶修改服务
type AddressUpdatePinService struct {
	ID     uint `form:"id" json:"id"`
	UserID uint `form:"user_id" json:"user_id"`
	Pinned bool `form:"pinned" json:"pinned"`
}

// AddressUpdatePin 修改地址是否置顶，并返回更新后的地址列表
func (service *AddressUpdatePinService) AddressUpdatePin() serializer.Response {
	code := constant.SUCCESS

	// 更新单个值，用Update
	err := model.DB.Model(&model.Address{}).Where("id = ?", service.ID).Update("pinned", service.Pinned).Error
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
	err = model.DB.Where("user_id = ?", service.UserID).Order("pinned desc, updated_at desc").Find(&addresses).Error
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
