package channel

import (
	"encoding/json"
	"github.com/luyasr/mpush/app/common"
)

type Channel struct {
	*common.Meta
	*Request
	Username int64  `json:"username" gorm:"not null;index"`
	Token    string `json:"token" gorm:"not null;uniqueIndex"`
}

type Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ChannelType string `json:"channel_type" gorm:"not null" validate:"required" label:"通道类型"`
	Url         string `json:"url"`
	Webhook     string `json:"webhook"`
	Secret      string `json:"secret"`
	AppId       string `json:"app_id"`
}

func (c *Channel) TableName() string {
	return "channels"
}

func (c *Channel) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
