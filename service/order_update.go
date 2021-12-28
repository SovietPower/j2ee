package service

import (
	"errors"
	"j2ee/model"
	"time"

	"gorm.io/gorm"
)

// UserInfoUpdate 当新订单完成时，更新用户信息
func UserInfoUpdate(UserID uint, RNname string) (err error) {
	var info model.UserInfo
	db := model.DB.Model(&model.UserInfo{}).Where("user_id = ?", UserID).First(&info)
	if err = db.Error; err != nil {
		return err
	}

	err = db.Update("send_number", info.SendNumber-1).Error
	if err = db.Error; err != nil {
		return err
	}

	db = model.DB.Model(&model.UserInfo{}).Where("username = ?", RNname).First(&info)
	if err = db.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}

	err = db.Update("receive_number", info.ReceiveNumber-1).Error
	if err = db.Error; err != nil {
		return err
	}

	return nil
}

// OrdersUpdate 检查并更新一部分订单完成状态
func OrdersUpdate(db *gorm.DB, orders []model.Order) (err error) {
	err = nil

	// 检查是否有刚好完成订单（目前只会在这里检查）
	flag := false
	now := time.Now()
	for index, order := range orders {
		if order.Status == 1 && order.Allocate != 0 && now.After(order.Time) {
			flag = true
			// 更新（不能更新order？）
			orders[index].Status = 0
			// 释放货车
			truckID := order.Allocate
			model.DB.Model(&model.Truck{}).Where("id = ?", truckID).Update("allocated", 0)
			// 更新UserInfo
			if err := UserInfoUpdate(order.UserID, order.RName); err != nil {
				return err
			}
			// 管理员信息更新
			adminInfo := AdminInfoUpdateService{
				DMsgNumber:      0,
				DOrderTotal:     0,
				DUserTotal:      0,
				DTruckTotal:     0,
				DTruckAvailable: 1,
				DOrderUnhandled: 0,
			}
			err, _ = adminInfo.AdminInfoUpdate()
			if err != nil {
				return err
			}
		}
	}
	// 更新未完成订单
	if flag {
		err = db.Save(&orders).Error
	}

	return err
}

// OrderUpdate 检查并更新单个订单完成状态
func OrderUpdate(db *gorm.DB, order model.Order) (err error) {
	err = nil

	// 检查是否有刚好完成订单（目前只会在这里检查）
	flag := false
	now := time.Now()
	if order.Status == 1 && order.Allocate != 0 && now.After(order.Time) {
		flag = true
		// 更新
		order.Status = 0
		// 释放货车
		truckID := order.Allocate
		model.DB.Model(&model.Truck{}).Where("id = ?", truckID).Update("allocated", 0)
		// 更新UserInfo
		if err := UserInfoUpdate(order.UserID, order.RName); err != nil {
			return err
		}
		// 管理员信息更新
		adminInfo := AdminInfoUpdateService{
			DMsgNumber:      0,
			DOrderTotal:     0,
			DUserTotal:      0,
			DTruckTotal:     0,
			DTruckAvailable: 1,
			DOrderUnhandled: 0,
		}
		err, _ = adminInfo.AdminInfoUpdate()
		if err != nil {
			return err
		}
	}

	// 更新未完成订单
	if flag {
		err = db.Save(&order).Error
	}

	return err
}
