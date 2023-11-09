package user

import (
	"encoding/json"
	"github.com/luyasr/mpush/app/common"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	*common.Meta
	Username  string                `json:"username" gorm:"not null;uniqueIndex:idx_username_deleted_at;types:varchar(20)"`
	Password  string                `json:"password" gorm:"not null;types:varchar(255)"`
	Nickname  string                `json:"nickname" gorm:"not null;uniqueIndex:idx_nickname_deleted_at;types:varchar(20)"`
	Email     string                `json:"email" gorm:"types:varchar(50)"`
	Role      Role                  `json:"role" gorm:"not null;types:tinyint"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"index:idx_username_deleted_at;index:idx_nickname_deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	bytes, _ := json.Marshal(u)
	return string(bytes)
}
