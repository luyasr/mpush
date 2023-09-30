package user

import (
	"context"
	"errors"
	"github.com/luyasr/mpush/app/utils"
	"github.com/luyasr/mpush/common/errs"
	"github.com/luyasr/mpush/common/validate"
	"github.com/luyasr/mpush/common/zerologger"
	"github.com/luyasr/mpush/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var _ Interface = (*Service)(nil)

type Service struct {
	db  *gorm.DB
	log zerolog.Logger
}

func init() {
	NewService()
}

func NewService() *Service {
	return &Service{
		db:  config.C.Mysql.GetConn(),
		log: zerologger.NewFileLog("user"),
	}
}

func (s *Service) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
	var user User
	// verify user create request
	err := validate.Struct(req)
	if err != nil {
		return nil, err
	}
	// instantiate
	user.Username = req.Username
	user.Password = utils.PasswordHash(req.Password)
	user.Nickname = req.Username
	user.Email = req.Email
	user.Role = RoleMember

	// check if the user exists
	byUsername, _ := s.GetByUsername(ctx, user.Username)
	if byUsername != nil {
		return nil, errs.NewExists("创建用户失败,用户%s已存在", user.Username)
	}

	err = s.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) DeleteById(ctx context.Context, id int64) error {
	byId, err := s.GetById(ctx, id)
	if err != nil {
		return err
	}
	err = s.db.WithContext(ctx).Delete(&byId).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, id int64, req *UpdateUserRequest) error {
	var user User
	// verify user delete request
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	user.Id = id

	if req.Username != "" {
		byUsername, _ := s.GetByUsername(ctx, req.Username)
		if byUsername != nil {
			return errs.NewExists("更新用户失败, 用户名称%s已存在", req.Username)
		}
	}
	if req.Nickname != "" {
		byNickname, _ := s.GetByNickname(ctx, req.Nickname)
		if byNickname != nil {
			return errs.NewExists("更新用户失败, 用户昵称%s已存在", req.Nickname)
		}
	}

	fields, err := utils.UpdateNoneZeroFields(req)
	if err != nil {
		return err
	}

	if len(fields) == 0 {
		return errs.NewUpdateFailed("更新用户id%d失败, 无更新字段", user.Id)
	}

	updates := s.db.WithContext(ctx).Model(&user).Updates(fields)
	err = updates.Error
	if err != nil {
		return err
	}
	affected := updates.RowsAffected
	if affected == 0 {
		return errs.NewUpdateFailed("更新用户id%d失败, 受影响的行记录为0", user.Id)
	}

	return nil
}

func (s *Service) GetById(ctx context.Context, id int64) (*User, error) {
	var user User
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound("用户id%d不存在", id)
		}
		return nil, err
	}
	return &user, nil
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := s.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound("用户名称%s不存在", username)
		}
		return nil, err
	}
	return &user, nil
}

func (s *Service) GetByNickname(ctx context.Context, nickname string) (*User, error) {
	var user User
	err := s.db.WithContext(ctx).Where("nickname = ?", nickname).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFound("用户昵称%s不存在", nickname)
		}
		return nil, err
	}
	return &user, nil
}
