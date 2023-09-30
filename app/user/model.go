package user

import (
	"encoding/json"
	"github.com/luyasr/mpush/common"
)

type User struct {
	common.Meta
	Username string `json:"username" gorm:"not null;uniqueIndex;type:varchar(20)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
	Nickname string `json:"nickname" gorm:"not null;uniqueIndex;type:varchar(20)"`
	Email    string `json:"email" gorm:"type:varchar(50)"`
	Role     Role   `json:"role" gorm:"not null;type:tinyint"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	jsonData, _ := json.Marshal(u)
	return string(jsonData)
}

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password   string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	RePassword string `json:"re_password" validate:"required,min=6,max=20,eqfield=Password" label:"确认密码"`
	Email      string `json:"email" validate:"required,email" label:"邮箱"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"omitempty,min=6,max=20" label:"密码"`
	Nickname string `json:"nickname" validate:"omitempty,min=3,max=20" label:"昵称"`
	Email    string `json:"email" validate:"omitempty,email" label:"邮箱"`
}
