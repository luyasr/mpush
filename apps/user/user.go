package user

import (
	"context"
	"github.com/luyasr/gaia/crypto/bcrypt"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/stores/mysql"
	"github.com/luyasr/gaia/validator"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
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
	c.db = ioc.Container.Get(ioc.DbNamespace, mysql.Name).(*mysql.Mysql).Client

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
	findUserReq := &FindReq{
		FindBy: FindByUsername,
		Value:  req.Username,
	}
	byUsername, _ := c.Find(ctx, findUserReq)
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

func (c *Controller) Find(ctx context.Context, req *FindReq) (*User, error) {
	switch req.FindBy {
	case FindById:
		id, err := strconv.ParseInt(req.Value, 10, 64)
		if err != nil {
			return nil, errors.BadRequest("", err.Error())
		}
		return c.findById(ctx, id)
	case FindByUsername:
		return c.findByUsername(ctx, req.Value)
	}

	return nil, nil
}

func (c *Controller) Delete(ctx context.Context, id int64) error {
	now := time.Now().Unix()

	return c.delete(ctx, id, now)
}
