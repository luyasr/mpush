package channel

type Channel struct {
	Id      int64  `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Type    string `json:"type" gorm:"not null" validate:"required" label:"通道类型"`
	UserId  int64  `json:"userId" gorm:"index" validate:"required" label:"用户id"`
	Token   string `json:"token" gorm:"uniqueIndex"`
	Url     string `json:"url"`
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
	AppId   string `json:"app_id"`
}

type CreateChannelRequest struct {
	Channel
}
