package message

import (
	"context"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/stores/kafka"
	"github.com/luyasr/gaia/stores/mysql"
	"github.com/luyasr/gaia/validator"
	kg "github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"time"
)

var _ Service = (*Controller)(nil)

type Controller struct {
	db     *gorm.DB
	reader *kg.Reader
	writer *kg.Writer
}

func init() {
	ioc.Container.Registry(ioc.ControllerNamespace, &Controller{})
}

func (c *Controller) Init() error {
	c.db = mysql.DB()
	c.reader = kafka.Consumer(Name)
	c.writer = kafka.Producer(Name)
	return nil
}

func (c *Controller) Name() string {
	return Name
}

func (c *Controller) ClientSend(ctx context.Context, req *ClientSendReq) error {
	if err := validator.Struct(req); err != nil {
		return err
	}

	message := new(Message)
	message.Title = req.Title
	message.Content = req.Content
	message.Status = StatusUnsent
	now := time.Now().Unix()
	message.CreatedAt = now
	message.UpdatedAt = now

	return c.clientSend(ctx, message)
}

func (c *Controller) Query(ctx context.Context, req *QueryReq) (*Messages, error) {
	if err := validator.Struct(req); err != nil {
		return nil, err
	}

	query := c.db.WithContext(ctx).Model(&Message{})

	// 根据查询条件构建查询
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.Keywords != "" {
		query = query.Where("title like ? OR content like ?", "%"+req.Keywords+"%", "%"+req.Keywords+"%")
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.PageNumber == 0 {
		req.PageNumber = 1
	}

	messages := new(Messages)

	// 查询总数
	if err := query.Count(&messages.Total).Error; err != nil {
		return nil, err
	}

	// 查询分页数据
	if err := query.Offset(req.offset()).Limit(req.PageSize).Find(&messages.Items).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *Controller) Read() {
}
