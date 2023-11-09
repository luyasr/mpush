package message

import (
	"context"
	"github.com/luyasr/mpush/config"
	"github.com/luyasr/mpush/pkg/zerologger"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var _ Interface = (*Service)(nil)

type Service struct {
	db  *gorm.DB
	log zerolog.Logger
}

func NewService() *Service {
	return &Service{
		db:  config.C.Mysql.GetConn(),
		log: zerologger.NewFileLog("message", zerologger.LogWithOptions{Dir: config.C.Log.Dir}),
	}
}

func (s *Service) CreateMessage(ctx context.Context, req *Request) (*Message, error) {
	var message *Message

	message.Request = req
	message.Status = StatusPending

	err := s.db.WithContext(ctx).Create(message).Error
	if err != nil {
		return nil, err
	}

	return message, nil
}
func (s *Service) UpdateMessage(ctx context.Context, req *UpdateMessageRequest) error {
	var message *Message

	message.ID = req.ID
	tx := s.db.WithContext(ctx).Model(message).Updates(req)
	if err := tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return ErrUpdateMessageFailed
	}

	return nil
}
func (s *Service) QueryMessage(ctx context.Context, req *QueryMessageRequest) (*Messages, error) {
	var messages *Messages

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
	err := query.Offset(req.Offset()).Limit(req.PageSize).Order("created_at DESC").Find(&messages.Items).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
func (s *Service) DeleteMessage(ctx context.Context, req *DeleteMessageRequest) error {
	var message *Message

	message.ID = req.ID
	tx := s.db.WithContext(ctx).Delete(message)
	if err := tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return ErrDeleteMessageFailed
	}
	return nil
}
