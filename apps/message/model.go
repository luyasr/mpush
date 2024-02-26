package message

import "github.com/goccy/go-json"

const Name = "message"

type Message struct {
	// Id 消息ID
	Id int64 `json:"id"`
	// UserId 用户ID
	UserId int64 `json:"user_id"`
	// Channel 频道
	Channel Channel `json:"channel"`
	// Title 标题
	Title string `json:"title"`
	// Content 内容
	Content string `json:"content"`
	// Status 状态
	Status int `json:"status"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (m *Message) TableName() string {
	return Name
}

func (m *Message) String() string {
	bytes, _ := json.Marshal(m)

	return string(bytes)
}
