package message

type Status int

const (
	StatusUnknown Status = iota
	StatusPending
	StatusSent
	StatusFailed
)

type Message struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	UserId    int64  `json:"user_id" gorm:"not null;uniqueIndex"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Content   string `json:"content"`
	To        string `json:"to"`
	Time      string `json:"time"`
	Timestamp int64  `json:"timestamp"`
	Status    Status `json:"status" gorm:"index"` // pending, sent, failed
}
