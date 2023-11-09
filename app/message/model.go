package message

import (
	"encoding/json"
	"github.com/luyasr/mpush/app/common"
	"gorm.io/plugin/soft_delete"
)

type Message struct {
	*common.Meta
	*Request
	Username    string                `json:"username" gorm:"not null;index"`
	ChannelType string                `json:"channel_type" gorm:"not null;index"`
	Status      Status                `json:"status" gorm:"not null;types:tinyint"` // pending, sent, failed
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"index:idx_username_deleted_at;index:idx_channel_deleted_at"`
}

type Request struct {
	Title   string `json:"title"`
	Text    string `json:"text"`    // 文本内容
	Content string `json:"content"` // markdown内容
	To      string `json:"to"`
}

func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) String() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}
