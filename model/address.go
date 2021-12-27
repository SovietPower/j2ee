package model

import (
	"gorm.io/gorm"
)

// Address 收货地址模型
type Address struct {
	gorm.Model
	UserID  uint   // 属于哪个用户
	Name    string // 姓名
	Phone   string
	Address string
	Pinned  bool // 是否置顶
}
