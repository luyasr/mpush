package user

import "github.com/luyasr/mpush/pkg/errorhandler"

type Role int

const (
	RoleMember Role = iota
	RoleAdmin
)

type QueryBy int

const (
	QueryById QueryBy = iota
	QueryByUsername
	QueryByNickName
	QueryByUsernameOrNickname
)

var (
	ErrUserNotFound   = errorhandler.NewNotFound("用户没找到")
	ErrUserExists     = errorhandler.NewExists("用户已存在")
	ErrUsernameExists = errorhandler.NewExists("用户名已存在")
	ErrNicknameExists = errorhandler.NewExists("用户昵称已存在")
)
