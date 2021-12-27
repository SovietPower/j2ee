package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// TruckGetService 获取所有货车信息服务
type TruckGetService struct {
	Type int `form:"type"` // 0：所有，1：未分配，2：已分配，3：不可用

	// 展示顺序由前端对UpdatedAt进行排序，不重复请求
	// Order int `form:"order"` // 0：按编号顺序显示，1：按更新时间显示
}

// TruckGet 获取所有货车信息
func (service *TruckGetService) TruckGet() serializer.Response {
	var trucks []model.Truck
	code := constant.SUCCESS

	db := model.DB.Model(&model.Truck{})
	if service.Type == 1 {
		db = db.Where("allocated = ?", 0)
	} else if service.Type == 2 {
		db = db.Where("allocated = ?", 1)
	} else if service.Type == 3 {
		db = db.Where("status = ?", 1)
	}

	// if service.Order == 1 {
	// 	db = db.Order("updated_at desc")
	// }
	err := db.Find(&trucks).Error
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

// TruckGetByUserService 用户获取可用货车数量服务
type TruckGetByUserService struct {
}

// TruckGetByUser 用户获取可用货车数量
func (service *TruckGetByUserService) TruckGetByUser() serializer.Response {
	var total int64
	code := constant.SUCCESS

	err := model.DB.Model(&model.Truck{}).Where("allocated = ?", 0).Count(&total).Error
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
		Data:   total,
	}
}
