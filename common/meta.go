package common

import "gorm.io/plugin/soft_delete"

type Meta struct {
	Id        int64                 `json:"id" gorm:"primaryKey"`
	CreatedAt int64                 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64                 `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"index"`
}
