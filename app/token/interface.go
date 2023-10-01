package token

import (
	"context"
)

type Interface interface {
	Login(context.Context, *LoginRequest) (*Token, error)
	Logout(context.Context, *Request) error
	Refresh(context.Context, *Request) error
	Validate(context.Context, *Request) error
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

type Request struct {
	UserId       int64  `json:"user_id" validate:"required" label:"用户id"`
	AccessToken  string `json:"access_token" validate:"required" label:"登录token"`
	RefreshToken string `json:"refresh_token" validate:"required" label:"刷新token"`
}
