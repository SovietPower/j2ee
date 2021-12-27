package serializer

import (
	"j2ee/model"
)

// Address 地址序列化器
type Address struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Pinned  bool   `json:"pinned"`
	Seen    bool   `json:"seen"` // 用于前端展示，全部在前端处理
	// CreatedAt int64  `json:"created_at"`
}

// BuildAddress 序列化地址
func BuildAddress(item model.Address) Address {
	return Address{
		ID:      item.ID,
		UserID:  item.UserID,
		Name:    item.Name,
		Phone:   item.Phone,
		Address: item.Address,
		Pinned:  item.Pinned,
		Seen:    false,
		// CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildAddresses 序列化地址列表
func BuildAddresses(items []model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return addresses
}
