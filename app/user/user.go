package user

import (
	"context"
	"errors"
	"github.com/luyasr/mpush/config"
	"github.com/luyasr/mpush/pkg/errorhandler"
	"github.com/luyasr/mpush/pkg/utils"
	"github.com/luyasr/mpush/pkg/validate"
	"github.com/luyasr/mpush/pkg/zerologger"
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
		log: zerologger.NewFileLog("user", zerologger.LogWithOptions{Dir: config.C.Log.Dir}),
	}
}

func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	var user *User
	// 校验创建用户请求
	err := validate.Struct(req)
	if err != nil {
		return nil, err
	}

	user.Username = req.Username
	user.Password = utils.PasswordHash(req.Password)
	user.Nickname = req.Username
	user.Email = req.Email
	user.Role = RoleMember

	// 检查用户是否存在
	queryUser, _ := s.QueryUser(ctx, NewQueryUserByUsernameRequest(req.Username))
	if queryUser != nil {
		return nil, ErrUserExists
	}

	err = s.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	var user *User
	// 校验用户更新请求
	if err := validate.Struct(req); err != nil {
		return err
	}

	// 更新前查询用户
	queryUser, _ := s.QueryUser(ctx, NewQueryUserByUsernameOrNicknameRequest(req.Username, req.Nickname))
	if queryUser.Username == req.Username {
		return ErrUsernameExists
	}
	if queryUser.Nickname == req.Nickname {
		return ErrNicknameExists
	}

	// TODO 合并用户的更新请求, 只允许字段的不同的更新

	user.ID = req.ID
	tx := s.db.WithContext(ctx).Model(&user).Updates(req)
	if err := tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return errorhandler.NewUpdateFailed("更新用户id%d失败, 受影响的行记录为0", user.ID)
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, req *DeleteUserRequest) error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	queryUser, err := s.QueryUser(ctx, NewQueryUserByIdRequest(req.ID))
	if err != nil {
		return err
	}

	err = s.db.WithContext(ctx).Delete(queryUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) QueryUser(ctx context.Context, req *QueryUserRequest) (*User, error) {
	// 校验查询请求
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	// 根据条件查询
	var user *User
	query := s.db.WithContext(ctx)
	switch req.QueryBy {
	case QueryById:
		query = query.Where("user = ?", req.QueryByValue...)
	case QueryByUsername:
		query = query.Where("username = ?", req.QueryByValue...)
	case QueryByUsernameOrNickname:
		query = query.Where("username = ? OR nickname = ?", req.QueryByValue...)
	}

	if err := query.First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
	}

	return user, nil
}
