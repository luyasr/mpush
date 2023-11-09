package common

type Meta struct {
	// 自增ID
	ID int64 `json:"id" gorm:"primaryKey"`
	// 创建时间
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" gorm:"autoCreateTime"`
}
