package user

import (
	"encoding/json"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        int64                 `json:"id" gorm:"primaryKey"`
	Username  string                `json:"username" gorm:"not null;uniqueIndex:idx_username_deleted_at;type:varchar(20)"`
	Password  string                `json:"password" gorm:"not null;type:varchar(255)"`
	Nickname  string                `json:"nickname" gorm:"not null;uniqueIndex:idx_nickname_deleted_at;type:varchar(20)"`
	Email     string                `json:"email" gorm:"type:varchar(50)"`
	Role      Role                  `json:"role" gorm:"not null;type:tinyint"`
	CreatedAt int64                 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64                 `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"index:idx_username_deleted_at;index:idx_nickname_deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	bytes, _ := json.Marshal(u)
	return string(bytes)
}
