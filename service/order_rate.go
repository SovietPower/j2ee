package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// OrderRateService 订单评价服务
type OrderRateService struct {
	ID     int `json:"id"`
	Rating int `json:"rating"`
}

// OrderRate 评价订单
func (service *OrderRateService) OrderRate() serializer.Response {
	var order model.Order
	code := constant.SUCCESS

	db := model.DB.Model(&model.Order{}).Where("id = ?", service.ID).First(&order)
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

	// 可以用update
	order.Rating = service.Rating
	err = db.Save(&order).Error
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
