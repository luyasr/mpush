package channel

import (
	"context"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	log zerolog.Logger
}

func (s *Service) CreateChannel(ctx context.Context, req *Request) (*Channel, error) {
	var channel *Channel

	channel.Request = req
	// TODO: 在上下文中传递用户信息
	if err := s.db.WithContext(ctx).Create(channel).Error; err != nil {
		return nil, err
	}

	return channel, nil
}
