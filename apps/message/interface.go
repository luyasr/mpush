package message

import "context"

type Service interface {
	// ClientSend 客户端发送消息
	ClientSend(ctx context.Context, req *ClientSendReq) error
	Query(ctx context.Context, req *QueryReq) (*Messages, error)
}

type ClientSendReq struct {
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
