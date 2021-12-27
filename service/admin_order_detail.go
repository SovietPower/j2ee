package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AdminOrderDetailService 获取订单详情服务
type AdminOrderDetailService struct {
}

// AdminOrderDetail 获取订单详情
func (service *AdminOrderDetailService) AdminOrderDetail(order_id string) serializer.Response {
	var order model.Order
	code := constant.SUCCESS

	db := model.DB.Model(&model.Order{}).Where("order_id = ?", order_id).First(&order)
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

	// 检查并更新状态
	OrderUpdate(db, order)

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Data:   serializer.BuildOrderDetail(order),
	}
}
