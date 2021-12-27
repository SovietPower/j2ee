package serializer

import "j2ee/model"

// Truck 货车序列化器
type Truck struct {
	ID        uint
	UpdatedAt int64 // 上次状态更新时间

	Status    int    // 货车状态。0：正常，1：不可用
	Allocated int    // 分配状态。0：未分配，1：已分配
	OrderID   string // 分配的订单号（如果已分配）

	// 保存部分订单信息，用于展示
	SAddress string // 发货地
	RAddress string // 收货地
	Type     string // 货物种类
}

// BuildTruck 序列化货车
func BuildTruck(item model.Truck) Truck {
	return Truck{
		ID:        item.ID,
		UpdatedAt: item.UpdatedAt.Unix(),

		Status:    item.Status,
		Allocated: item.Allocated,
		OrderID:   item.OrderID,

		SAddress: item.SAddress,
		RAddress: item.RAddress,
		Type:     item.Type,
	}
}

// BuildTrucks 序列化货车列表
func BuildTrucks(items []model.Truck) (trucks []Truck) {
	for _, item := range items {
		truck := BuildTruck(item)
		trucks = append(trucks, truck)
	}
	return trucks
}
