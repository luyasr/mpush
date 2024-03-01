package channel

import "encoding/json"

const (
	Name = "channel"
)

type Channel struct {
	// Id 频道ID
	Id int64 `json:"id"`
	// Name 频道名称
	Name string `json:"name"`
	// Token 频道token
	Token string `json:"token"`
	// UserId 用户ID
	UserId int64 `json:"user_id"`
	// Url 频道地址
	Url string `json:"url"`
	// Secret 频道密钥
	Secret string `json:"secret"`
	// CreatedAt 创建时间
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt int64 `json:"updated_at"`
}

func (Channel) TableName() string {
	return Name
}

func (c *Channel) String() string {
	bytes, _ := json.Marshal(c)

	return string(bytes)
}
