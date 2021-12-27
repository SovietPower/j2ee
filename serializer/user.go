package serializer

import "j2ee/model"

// User 用户序列化器。将模型转为前端存储的结构
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`

	Status    int `json:"status"`
	Authority int `json:"authority"` // 用户权限，0：管理员，1：普通用户
	// Email     string `json:"email"`
}

// BuildUser 序列化用户。将模型转为前端存储的结构
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Unix(),

		Status:    user.Status,
		Authority: user.Authority,
		// Email:     user.Email,
	}
}
