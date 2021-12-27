package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// TruckAllocateService 分配货车服务
type TruckAllocateService struct {
	ID uint `json:"id"`

	OrderID string `json:"order_id"` // 分配的订单号
}

// TruckAllocate 分配货车
func (service *TruckAllocateService) TruckAllocate() serializer.Response {
	var truck model.Truck
	code := constant.SUCCESS

	db := model.DB.Model(&model.Truck{}).Where("id = ?", service.ID).First(&truck)
	err := db.Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 订单状态更新
	var order model.Order
	db2 := model.DB.Model(&model.Order{}).Where("order_id = ?", service.OrderID).First(&order)
	err = db2.Error
	if err != nil {
		// 请求订单号不合法
		code = constant.ERROR_NO_SUCH_ORDER
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}
	if order.Allocate != 0 {
		// 订单已分配
		code = constant.INVALID_PARAMS
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	order.Allocate = int(service.ID)
	order.Time = TimeEvaluate(order.SAddress, order.RAddress)

	err = db2.Save(&order).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_UPDATING_INFO
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 货车状态更新
	truck.Allocated = 1
	truck.OrderID = service.OrderID
	truck.SAddress = order.SAddress
	truck.RAddress = order.RAddress
	truck.Type = order.Type

	err = db.Save(&truck).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_UPDATING_INFO
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
