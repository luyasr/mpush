package message

type Status int

type Message struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id" gorm:"index"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Content   string `json:"content"`
	To        string `json:"to"`
	Time      string `json:"time"`
	Timestamp int64  `json:"timestamp"`
	Status    int    `json:"status" gorm:"default:1;index"` // pending, sent, failed
}
