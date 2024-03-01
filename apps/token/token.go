package token

import (
	"context"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/stores/mysql"
	"github.com/luyasr/gaia/validator"
	"github.com/luyasr/mpush/apps/user"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"time"
)

const (
	tokenInvalid  = "invalid token"
	tokenNotFound = "%v token not found"
	tokenExpired  = "%v token expired"
)

var _ Service = (*Controller)(nil)

type Controller struct {
	db  *gorm.DB
	log *log.Helper

	user user.Service
}

func init() {
	ioc.Container.Registry(ioc.ControllerNamespace, &Controller{})
}

func (c *Controller) Init() error {
	c.db = mysql.DB()
	c.log = log.NewHelper(zerolog.DefaultLogger)
	c.user = ioc.Container.Get(ioc.ControllerNamespace, user.Name).(user.Service)

	return nil
}

func (c *Controller) Name() string {
	return Name
}

func (c *Controller) QueryByUserId(ctx context.Context, userId int64) (*Token, error) {
	return c.findByUserId(ctx, userId)
}

func (c *Controller) QueryByToken(ctx context.Context, token string) (*Token, error) {
	return c.findByToken(ctx, token)
}

func (c *Controller) Login(ctx context.Context, req *LoginReq) (*Tk, error) {
	// 查询用户是否存在
	findUserReq := &user.QueryReq{
		QueryBy: user.QueryByUsername,
		Value:   req.Username,
	}
	byUsername, err := c.user.Query(ctx, findUserReq)
	if err != nil {
		return nil, err
	}

	// 查询用户是否已经登录
	byUserId, _ := c.QueryByUserId(ctx, byUsername.Id)

	// 如果用户已经登录, 则更新token
	if byUserId != nil {
		// TODO: 更新token
		byUserId.AccessToken = xid.New().String()
		byUserId.UpdatedAt = time.Now().Unix()
		if err := c.update(ctx, byUserId); err != nil {
			return nil, err
		}
		return &Tk{
			AccessToken:  byUserId.AccessToken,
			RefreshToken: byUserId.RefreshToken,
		}, nil
	}

	// 创建token
	token := NewToken(byUsername.Id)
	// 创建token
	return c.login(ctx, token)
}

func (c *Controller) Logout(ctx context.Context, req *Tk) error {
	return nil
}

func (c *Controller) Refresh(ctx context.Context, req *Tk) (string, error) {
	if err := validator.Struct(req); err != nil {
		return "", errors.BadRequest(tokenInvalid, err.Error())
	}

	return "", nil
}

func (c *Controller) Validate(ctx context.Context, token string) (*Token, error) {
	byToken, err := c.findByToken(ctx, token)
	if err != nil {
		return nil, errors.Unauthorized(tokenInvalid, tokenNotFound, token)
	}

	if byToken.IsExpired() {
		return nil, errors.Unauthorized(tokenExpired, tokenExpired, token)
	}

	return byToken, nil
}
