package service

import (
	"j2ee/constant"
	"j2ee/model"
)

// AdminInfoUpdateService 更新管理员其它信息服务
type AdminInfoUpdateService struct {
	DMsgNumber  int
	DOrderTotal int
	DUserTotal  int
	DTruckTotal int

	DTruckAvailable int
	DOrderUnhandled int
}

// AdminInfoUpdate 获取管理员其它信息
func (service *AdminInfoUpdateService) AdminInfoUpdate() (err error, code int) {
	var info model.AdminInfo
	code = constant.SUCCESS

	db := model.DB.Model(&model.AdminInfo{}).First(&info)
	err = db.Error
	if err != nil {
		return err, constant.ERROR_DATABASE
	}

	info.MsgNumber = uint(int(info.MsgNumber) + service.DMsgNumber)
	info.OrderTotal = uint(int(info.OrderTotal) + service.DOrderTotal)
	info.UserTotal = uint(int(info.UserTotal) + service.DUserTotal)
	info.TruckTotal = uint(int(info.TruckTotal) + service.DTruckTotal)

	info.TruckAvailale = uint(int(info.TruckAvailale) + service.DTruckAvailable)
	info.OrderUnhandled = uint(int(info.OrderUnhandled) + service.DOrderUnhandled)

	err = db.Save(&info).Error
	if err != nil {
		return err, constant.ERROR_UPDATING_INFO
	}

	return nil, code
}
