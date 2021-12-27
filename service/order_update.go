package service

import (
	"j2ee/model"
	"time"

	"gorm.io/gorm"
)

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
	}
	// 更新未完成订单
	if flag {
		err = db.Save(&order).Error
	}

	return err
}
