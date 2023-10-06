package user

import "context"

type Interface interface {
	Create(context.Context, *CreateUserRequest) (*User, error)
	DeleteUserById(context.Context, int64) error
	Update(context.Context, int64, *UpdateUserRequest) error
	GetUserById(context.Context, int64) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
	GetUserByNickname(context.Context, string) (*User, error)
}
