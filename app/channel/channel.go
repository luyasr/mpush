package channel

type Channel struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"userId" gorm:"index"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Token  string `json:"token" gorm:"uniqueIndex"`
	Url    string `json:"url"`
	Secret string `json:"secret"`
}
