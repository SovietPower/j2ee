package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// OrderDetailService 获取订单详情服务
type OrderDetailService struct {
	UserID   uint   `form:"user_id"`  // 用于限制查看权限
	Username string `form:"username"` // 用于限制查看权限（接收方可查看订单详情）
}

// OrderDetail 获取订单详情
func (service *OrderDetailService) OrderDetail(order_id string) serializer.Response {
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

	if order.UserID != service.UserID && order.RName != service.Username {
		code = constant.ERROR_NO_PERMISSION
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
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
