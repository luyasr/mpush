package message

import "context"

type Interface interface {
	Create(context.Context, *CreateMessageRequest) (*Message, error)
	UpdateMessageStatusByUserId(context.Context, int64, Status) error
	UpdateMessageStatusByUserIds(context.Context, []int64, Status) error
	GetMessageByUserId(context.Context, int64) (*Message, error)
	GetMessageByUserIds(context.Context, []int64) ([]*Message, error)
}

type CreateMessageRequest struct {
}
