package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// TruckCreateService 货车添加服务
type TruckCreateService struct {
}

// TruckCreate 创建一个货车信息，并返回更新后的货车信息列表
func (service *TruckCreateService) TruckCreate() serializer.Response {
	truck := model.Truck{
		Status:    0,
		Allocated: 0,
		OrderID:   "",

		SAddress: "",
		RAddress: "",
		Type:     "",
	}
	code := constant.SUCCESS

	err := model.DB.Model(&model.Truck{}).Create(&truck).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var trucks []model.Truck
	err = model.DB.Model(&model.Truck{}).Find(&trucks).Error
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
		Data:   serializer.BuildTrucks(trucks),
	}
}
