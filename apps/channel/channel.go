package channel

import (
	"context"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/reflection"
	"github.com/luyasr/gaia/stores/mysql"
	"github.com/luyasr/gaia/validator"
	"github.com/luyasr/mpush/apps/token"
	"gorm.io/gorm"
	"time"
)

type Controller struct {
	db *gorm.DB

	tokenController *token.Controller
}

func init() {
	ioc.Container.Registry(ioc.ControllerNamespace, &Controller{})
}

func (c *Controller) Init() error {
	c.db = mysql.DB()
	c.tokenController = ioc.Container.Get(ioc.ControllerNamespace, token.Name).(*token.Controller)
	return nil
}

func (c *Controller) Name() string {
	return Name
}

func (c *Controller) Create(ctx context.Context, req *CreateReq) (*Channel, error) {
	if err := validator.Struct(req); err != nil {
		return nil, err
	}

	tk := ctx.Value(token.Name).(*token.Token)

	channel := new(Channel)
	now := time.Now().Unix()
	channel.UserId = tk.UserId
	channel.Name = req.Name
	channel.Url = req.Url
	channel.Secret = req.Secret
	channel.CreatedAt = now
	channel.UpdatedAt = now

	create, err := c.create(ctx, channel)
	if err != nil {
		return nil, err
	}

	return create, nil
}

func (c *Controller) Update(ctx context.Context, req *UpdateReq) error {
	structToMap := reflection.StructToMap(req)

	return c.update(ctx, structToMap)
}

func (c *Controller) Delete(ctx context.Context, id int64) error {
	return c.delete(ctx, id)
}

func (c *Controller) Query(ctx context.Context, req *QueryReq) (*Channels, error) {
	// 根据查询条件构建查询
	query := c.db.WithContext(ctx).Model(&Channel{})

	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.PageNumber == 0 {
		req.PageNumber = 1
	}
	if req.Keywords != "" {
		query = query.Where("name like ?", "%"+req.Keywords+"%")
	}

	channels := new(Channels)
	if err := query.Count(&channels.Total).Error; err != nil {
		return nil, err
	}

	if err := query.Offset(req.offset()).Limit(req.PageSize).Find(&channels.Items).Error; err != nil {
		return nil, err
	}

	return channels, nil
}
