package token

import (
	"context"
	"errors"
	"github.com/luyasr/mpush/app/user"
	"github.com/luyasr/mpush/app/utils"
	"github.com/luyasr/mpush/common/errs"
	"github.com/luyasr/mpush/common/validate"
	"github.com/luyasr/mpush/common/zerologger"
	"github.com/luyasr/mpush/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"time"
)

var _ Interface = (*Service)(nil)

type Service struct {
	db   *gorm.DB
	user *user.Service
	log  zerolog.Logger
}

func init() {
	NewService()
}

func NewService() *Service {
	return &Service{
		db:   config.C.Mysql.GetConn(),
		user: user.NewService(),
		log:  zerologger.NewFileLog("token"),
	}
}

func (s *Service) Login(ctx context.Context, req *LoginRequest) (*Token, error) {
	// 校验用户登录请求
	err := validate.Struct(req)
	if err != nil {
		s.log.Error().Err(err).Stack().Msg("用户登录请求校验失败")
		return nil, err
	}
	// 根据用户名称查询用户信息, 用户不存在或者密码不一致返回用户名或密码错误
	byUsername, err := s.user.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errs.NewAuthFailed("用户名或密码错误")
	}
	err = utils.PasswordCompare(byUsername.Password, req.Password)
	if err != nil {
		return nil, errs.NewAuthFailed("用户名或密码错误")
	}

	// 实例化一个token对象, 并把用户id赋值
	token := NewDefaultToken()
	token.UserId = byUsername.Id

	// 先查询用户是否登陆过, 如果存在登录记录直接更新token, 否则新建
	err = s.db.WithContext(ctx).Where("user_id = ?", token.UserId).First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = s.db.WithContext(ctx).Create(&token).Error
			if err != nil {
				return nil, err
			}
			return token, nil
		}
	} else {
		err := s.db.WithContext(ctx).Model(&token).Updates(token).Error
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

func (s *Service) Logout(ctx context.Context, req *Request) error {
	var token Token
	// 校验用户退出请求
	err := validate.Struct(req)
	if err != nil {
		return err
	}

	// 退出前查询用户是否存在
	err = s.db.WithContext(ctx).Where("user_id = ?", req.UserId).First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFound("用户id%d不存在", req.UserId)
		}
		return err
	}
	// 删除用户token
	tx := s.db.WithContext(ctx).
		Where("user_id = ? AND access_token = ? AND refresh_token = ?", req.UserId, req.AccessToken, req.RefreshToken).
		Delete(&token)
	if err = tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return errs.NewUpdateFailed("用户id%d退出失败", req.UserId)
	}

	return nil
}

func (s *Service) Refresh(ctx context.Context, req *Request) error {
	var token Token
	err := s.db.WithContext(ctx).
		Where("user_id = ? AND access_token = ? AND refresh_token = ?", req.UserId, req.AccessToken, req.RefreshToken).
		First(&token).Error
	if err != nil {
		return errs.NewAuthFailed("无效的token")
	}

	accessToken := token.Refresh()
	err = s.db.WithContext(ctx).Model(&token).Updates(
		Token{
			AccessToken:          accessToken,
			AccessTokenExpiredAt: time.Now().Add(2 * time.Hour).UnixMilli(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Validate(ctx context.Context, req *Request) error {
	var token Token
	err := s.db.WithContext(ctx).
		Where("user_id = ? AND access_token = ? AND refresh_token = ?", req.UserId, req.AccessToken, req.RefreshToken).
		First(&token).Error
	if err != nil {
		return errs.NewAuthFailed("无效的token")
	}
	if time.Now().UnixMilli() > token.AccessTokenExpiredAt {
		return errs.NewAuthFailed("无效的token")
	}
	if time.Now().UnixMilli() > token.RefreshTokenExpiredAt {
		return errs.NewAuthFailed("登录过期, 请重新登录")
	}
	return nil
}
