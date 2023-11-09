package message

import (
	"context"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var _ Interface = (*Service)(nil)

type Service struct {
	db  *gorm.DB
	log zerolog.Logger
}

func (s *Service) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*Message, error) {
	return nil, nil
}
func (s *Service) UpdateMessage(ctx context.Context, req *UpdateMessageRequest) error {
	return nil
}
func (s *Service) QueryMessage(ctx context.Context, req *QueryMessageRequest) (*Message, error) {
	messages := NewMessages()

	// 根据条件查询
	query := s.db.WithContext(ctx).Model(&Message{})
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}
	if req.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+req.Keywords+"%")
	}
	if len(req.Usernames) > 0 {
		query = query.Where("username IN ?", req.Usernames)
	}
	if len(req.Channels) > 0 {
		query = query.Where("channel IN ?", req.Channels)
	}

	// 查询total总数, 分页必须
	if err := query.Count(&messages.Total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	query.Offset(req.Offset()).Limit(req.PageSize)

	return nil, nil
}
func (s *Service) DeleteMessage(ctx context.Context, req *DeleteMessageRequest) error {
	return nil
}
