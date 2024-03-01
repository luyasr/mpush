package message

import "context"

type Service interface {
	// Producer 客户端发送消息
	Producer(ctx context.Context, req *ProducerReq) error
	// Consumer 服务端消费消息
	Consumer(ctx context.Context) error
	// Create 创建消息
	Create(ctx context.Context, message *Message) error
	// Query 查询消息
	Query(ctx context.Context, req *QueryReq) (*Messages, error)
}

type ProducerReq struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type QueryReq struct {
	Status     *Status `json:"status"`
	PageSize   int     `json:"page_size"`
	PageNumber int     `json:"page_number"`
	Keywords   string  `json:"keywords"`
}

func (q *QueryReq) offset() int {
	return (q.PageNumber - 1) * q.PageSize
}

type Messages struct {
	Total int64      `json:"total"`
	Items []*Message `json:"items"`
}
