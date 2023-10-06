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

func (s *Service) Create(ctx context.Context, req *CreateChannelRequest) (*Channel, error) {
	return nil, nil
}

func (s *Service) Validate(ctx context.Context, req *CreateChannelRequest) error {

	return nil
}
