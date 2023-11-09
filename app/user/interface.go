package user

import "context"

type Interface interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) error
	QueryUser(context.Context, *QueryUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
}

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password   string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	RePassword string `json:"re_password" validate:"required,min=6,max=20,eqfield=Password" label:"确认密码"`
	Email      string `json:"email" validate:"required,email" label:"邮箱"`
}

type UpdateUserRequest struct {
	ID       int64  `json:"id" validate:"omitempty" label:"用户id"`
	Username string `json:"username" validate:"omitempty,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"omitempty,min=6,max=20" label:"密码"`
	Nickname string `json:"nickname" validate:"omitempty,min=3,max=20" label:"昵称"`
	Email    string `json:"email" validate:"omitempty,email" label:"邮箱"`
}

type QueryUserRequest struct {
	QueryBy      QueryBy `json:"query_by"`
	QueryByValue []any   `json:"query_by_value"`
}

func NewQueryUserByIdRequest(id int64) *QueryUserRequest {
	return &QueryUserRequest{
		QueryBy:      QueryById,
		QueryByValue: []any{id},
	}
}

func NewQueryUserByUsernameRequest(username string) *QueryUserRequest {
	return &QueryUserRequest{
		QueryBy:      QueryById,
		QueryByValue: []any{username},
	}
}

func NewQueryUserByUsernameOrNicknameRequest(username, nickname string) *QueryUserRequest {
	return &QueryUserRequest{
		QueryBy:      QueryById,
		QueryByValue: []any{username, nickname},
	}
}

type DeleteUserRequest struct {
	ID int64 `json:"id" validate:"omitempty" label:"用户id"`
}
