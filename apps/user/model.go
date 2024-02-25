package user

import (
	"encoding/json"
	"time"
)

const (
	Name = "user"
)

type User struct {
	// ID 用户ID
	Id int64 `json:"id"`
	// Username 用户名
	Username string `json:"username"`
	// Password 密码
	Password string `json:"password"`
	// Nickname 昵称
	Nickname string `json:"nickname"`
	// Email 邮箱
	Email string `json:"email"`
	// Phone 手机号
	Phone string `json:"phone"`
	// Status 状态
	Status *Status `json:"status"`
	// Role 角色
	Role *Role `json:"role"`
	// CreatedAt 创建时间
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt int64 `json:"updated_at"`
	// DeletedAt 删除时间
	DeletedAt int64 `json:"deleted_at"`
}

func (User) TableName() string {
	return Name
}

func (u *User) String() string {
	bytes, _ := json.Marshal(u)

	return string(bytes)
}

func NewUser() *User {
	now := time.Now().Unix()
	status := StatusNormal
	role := RoleAnonymous

	return &User{
		Status:    &status,
		Role:      &role,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
