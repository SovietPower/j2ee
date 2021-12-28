package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	PasswordDigest string

	Status    int // 状态。0：正常，1：限制功能，2：封禁
	Authority int // 权限。0：管理员，1：普通用户
	// Email     string `gorm:"unique"`
	// Avatar    string
}

const (
	// 密码加密难度
	PassWordCost = 12

	StatusNormal  = 0
	StatusLimited = 1
	StatusBanned  = 2

	AuthorityAdmin = 0
	AuthorityUser  = 1
)

// 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// 生成并设置加密后的密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 校验用户密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
