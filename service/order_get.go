package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// OrderGetService 获取订单服务
type OrderGetService struct {
	UserID   uint   `form:"user_id" binding:"required"`
	Username string `form:"username"` // 用户RealName，用于查询待接收订单
	Limit    int    `form:"limit"`    // 显示数量
	Offset   int    `form:"offset"`   // 从符合条件的第几个开始显示
	Type     int    `form:"type"`     // 0：所有，1：发出，2：待接收，3：已完成但未评价
}

// OrderGet 获取订单
func (service *OrderGetService) OrderGet() serializer.Response {
	var orders []model.Order
	var total int64
	code := constant.SUCCESS

	if service.Limit == 0 {
		service.Limit = 5
	}

	db := model.DB.Model(&model.Order{})
	if service.Type == 0 {
		db = db.Where("user_id = ? OR r_name = ?", service.UserID, service.Username)
	} else if service.Type == 1 {
		db = db.Where("user_id = ? AND status = ?", service.UserID, 1)
	} else if service.Type == 2 {
		db = db.Where("r_name = ? AND status = ?", service.Username, 1)
	} else if service.Type == 3 {
		db = db.Where("user_id = ? AND rating = ?", service.UserID, 0)
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
