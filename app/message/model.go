package message

import "encoding/json"

type Message struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	UserID    int64  `json:"user_id" gorm:"not null;index"`
	Username  string `json:"username" gorm:"not null;index"`
	ChannelID int64  `json:"channel_id" gorm:"not null;index"`
	Channel   string `json:"channel" gorm:"not null;index"`
	Title     string `json:"title"`
	Text      string `json:"text"`    // 文本内容
	Content   string `json:"content"` // markdown内容
	To        string `json:"to"`
	Timestamp int64  `json:"timestamp"`
	Status    Status `json:"status" gorm:"not null;type:tinyint"` // pending, sent, failed
}

func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) String() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}
