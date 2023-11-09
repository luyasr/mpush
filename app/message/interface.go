package message

import (
	"context"
	"github.com/luyasr/mpush/pkg/utils"
)

type Interface interface {
	CreateMessage(context.Context, *Request) (*Message, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) error
	QueryMessage(context.Context, *QueryMessageRequest) (*Messages, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) error
}

type UpdateMessageRequest struct {
	ID     int64   `json:"id"`
	Status *Status `json:"status"`
}

type QueryMessageRequest struct {
	Status     *Status  `json:"status"`      // 消息状态
	PageSize   int      `json:"page_size"`   // 消息分页页码
	PageNumber int      `json:"page_number"` // 消息分页大小
	Usernames  []string `json:"usernames"`   // 消息发送用户
	Channels   []string `json:"channels"`    // 消息发送通道
	Keywords   string   `json:"keywords"`    // 消息关键词
}

func (q *QueryMessageRequest) SetStatus(status Status) {
	q.Status = &status
}

func (q *QueryMessageRequest) Offset() int {
	return q.PageSize * (q.PageNumber - 1)
}

func (q *QueryMessageRequest) ParsePageSize(pageSize string) {
	q.PageSize = utils.StringToInt(pageSize)
}

func (q *QueryMessageRequest) ParsePageNumber(pageNumber string) {
	q.PageNumber = utils.StringToInt(pageNumber)
}

func NewQueryMessageRequest() *QueryMessageRequest {
	return &QueryMessageRequest{
		PageNumber: 1,
		PageSize:   10,
	}
}

type Messages struct {
	Total int64     `json:"total"`
	Items []Message `json:"items"`
}

type DeleteMessageRequest struct {
	ID int64 `json:"id"`
}
