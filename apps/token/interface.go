package token

import "context"

type Service interface {
	FindByUserId(ctx context.Context, userId int64) (*Token, error)
	FindByToken(ctx context.Context, token string) (*Token, error)
	// Login 登录
	Login(ctx context.Context, req *LoginReq) (*Tk, error)
	// Logout 登出
	Logout(ctx context.Context, req *Tk) error
	// Refresh 刷新
	Refresh(ctx context.Context, req *Tk) (string, error)
	Validate(ctx context.Context, req *ValidateReq) (*Token, error)
}

// LoginReq 登录请求
type LoginReq struct {
	// 用户名
	Username string `json:"username" validate:"required" label:"用户名"`
	// 密码
	Password string `json:"password" validate:"required" label:"密码"`
}

// Tk 登录token 刷新token
type Tk struct {
	// 登录token
	Token string `json:"token" validate:"required" label:"token"`
	// 刷新token
	RefreshToken string `json:"refresh_token" validate:"required" label:"refresh_token"`
}

// ValidateReq 验证token请求
type ValidateReq struct {
	// 登录token
	Token string `json:"token" validate:"required" label:"token"`
}
