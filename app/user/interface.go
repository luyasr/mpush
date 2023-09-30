package user

import "context"

type Interface interface {
	Create(context.Context, *CreateUserRequest) (*User, error)
	DeleteById(context.Context, int64) error
	Update(context.Context, int64, *UpdateUserRequest) error
	GetById(context.Context, int64) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	GetByNickname(context.Context, string) (*User, error)
}
