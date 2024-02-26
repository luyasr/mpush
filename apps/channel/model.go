package channel

type Channel struct {
	// Id 频道ID
	Id int64 `json:"id"`
	// Name 频道名称
	Name string `json:"name"`
	// UserId 用户ID
	UserId int64 `json:"user_id"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at"`
}
