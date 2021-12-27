package serializer

import (
	"j2ee/model"
)

// Order 订单序列化器
type Order struct {
	ID uint `json:"id"` // 订单在数据库中的编号，用于查询订单详情（用大数order_id查好像有点影响效率？）

	// 发件人与收件人（分别展示给收件人、发件人）
	SName    string `json:"s_name"`
	SPhone   string `json:"s_phone"`
	SAddress string `json:"s_address"`
	RName    string `json:"r_name"`
	RPhone   string `json:"r_phone"`
	RAddress string `json:"r_address"`

	// 货物信息
	Type string `json:"type"`
	// Weight int    `json:"weight"`
	// Volume int    `json:"volume"`
	// Value  int    `json:"value"`
	Urgent bool `json:"urgent"`
	// Note   string `json:"note"`

	// 订单状态
	OrderID string `json:"order_id"`
	// Time     int64  `json:"time"`
	Status   int `json:"status"`
	Allocate int `json:"allocate"`
	Rating   int `json:"rating"`

	CreatedAt int64 `json:"created_at"`
}

// BuildOrder 序列化订单
func BuildOrder(item model.Order) Order {
	return Order{
		ID: item.ID,

		SName:    item.SName,
		SPhone:   item.SPhone,
		SAddress: item.SAddress,
		RName:    item.RName,
		RPhone:   item.RPhone,
		RAddress: item.RAddress,

		Type: item.Type,
		// Weight: item.Weight,
		// Volume: item.Volume,
		// Value:  item.Value,
		Urgent: item.Urgent,
		// Note:   item.Note,

		OrderID: item.OrderID,
		// Time:     item.Time.Unix(),
		Status:   item.Status,
		Allocate: item.Allocate,
		Rating:   item.Rating,

		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildOrders 序列化订单列表
func BuildOrders(items []model.Order) (orders []Order) {
	for _, item := range items {
		order := BuildOrder(item)
		orders = append(orders, order)
	}
	return orders
}

// OrderDetail 订单详情序列化器
type OrderDetail struct {
	ID uint `json:"id"` // 订单在数据库中的编号，便于查询或更新

	// 发件人与收件人
	SName    string `json:"s_name"`
	SPhone   string `json:"s_phone"`
	SAddress string `json:"s_address"`
	RName    string `json:"r_name"`
	RPhone   string `json:"r_phone"`
	RAddress string `json:"r_address"`

	// 货物信息
	Type   string `json:"type"`
	Weight int    `json:"weight"`
	Volume int    `json:"volume"`
	Value  int    `json:"value"`
	Urgent bool   `json:"urgent"`
	Note   string `json:"note"`

	// 订单状态
	OrderID  string `json:"order_id"`
	Time     int64  `json:"time"`
	Status   int    `json:"status"`
	Allocate int    `json:"allocate"`
	Rating   int    `json:"rating"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// BuildOrderDetail 序列化订单详情
func BuildOrderDetail(item model.Order) OrderDetail {
	return OrderDetail{
		ID: item.ID,

		SName:    item.SName,
		SPhone:   item.SPhone,
		SAddress: item.SAddress,
		RName:    item.RName,
		RPhone:   item.RPhone,
		RAddress: item.RAddress,

		Type:   item.Type,
		Weight: item.Weight,
		Volume: item.Volume,
		Value:  item.Value,
		Urgent: item.Urgent,
		Note:   item.Note,

		OrderID:  item.OrderID,
		Time:     item.Time.Unix(),
		Status:   item.Status,
		Allocate: item.Allocate,
		Rating:   item.Rating,

		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}
