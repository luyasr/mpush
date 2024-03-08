package user

import (
	"context"
	"github.com/luyasr/gaia/crypto/bcrypt"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/stores/mysql"
	"github.com/luyasr/gaia/validator"
	"gorm.io/gorm"
	"time"
)

const (
	invalid           = "无效的用户名或密码"
	userNotFound      = "user %v not found"
	userAlreadyExists = "user %v already exists"
)

var _ Service = (*Controller)(nil)

type Controller struct {
	db *gorm.DB
}

func init() {
	ioc.Container.Registry(ioc.ControllerNamespace, &Controller{})
}

func (c *Controller) Init() error {
	c.db = mysql.DB()

	return nil
}

func (c *Controller) Name() string {
	return Name
}

func (c *Controller) Create(ctx context.Context, req *CreateReq) (*User, error) {
	// 校验请求参数
	if err := validator.Struct(req); err != nil {
		return nil, errors.BadRequest("", err.Error())
	}

	// 查询用户是否已经存在
	byUsername, _ := c.queryByUsername(ctx, req.Username)
	if byUsername != nil {
		return nil, errors.BadRequest("", userAlreadyExists, req.Username)
	}

	// 构造用户实例
	user := NewUser()

	user.Username = req.Username
	hashPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Internal("", err.Error())
	}
	user.Password = hashPassword

	// 创建用户
	return c.create(ctx, user)
}

func (c *Controller) QueryById(ctx context.Context, id int64) (*User, error) {
	return c.queryById(ctx, id)
}

func (c *Controller) QueryByUsername(ctx context.Context, username string) (*User, error) {
	return c.queryByUsername(ctx, username)
}

func (c *Controller) QueryByUsernameAndPassword(ctx context.Context, username, password string) (*User, error) {
	user, err := c.queryByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.ComparePassword(password, user.Password); err != nil {
		return nil, errors.BadRequest(invalid, invalid)
	}

	return user, nil
}

func (c *Controller) Delete(ctx context.Context, id int64) error {
	now := time.Now().Unix()

	return c.delete(ctx, id, now)
}
