package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// AdminOrderGetService 管理员获取订单服务
type AdminOrderGetService struct {
	UserID uint `form:"user_id"` // 当不为0时，表示查询某用户的订单
	Limit  int  `form:"limit"`   // 显示数量
	Offset int  `form:"offset"`  // 从符合条件的第几个开始显示
	Type   int  `form:"type"`    // 0：所有，1：未分配，2：紧急（但未完成）订单
}

// AdminOrderGet 管理员获取订单
func (service *AdminOrderGetService) AdminOrderGet() serializer.Response {
	var orders []model.Order
	var total int64
	code := constant.SUCCESS

	if service.Limit == 0 {
		service.Limit = 5
	}

	db := model.DB.Model(&model.Order{})
	if service.UserID != 0 {
		db = db.Where("user_id = ?", service.UserID)
	}

	if service.Type == 1 {
		db = db.Where("allocate = ?", 0)
	} else if service.Type == 2 {
		db = db.Where("urgent = ? AND status = ?", 1, 1)
	}

	err := db.Count(&total).Error
	if err == nil {
		db = db.Limit(service.Limit).Offset(service.Offset).Order("created_at desc").Find(&orders)
		err = db.Error
	}
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	OrdersUpdate(db, orders)

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Data:   serializer.BuildDataWithTotal(serializer.BuildOrders(orders), uint(total)),
	}
}
