package user

import "context"

type Service interface {
	Create(ctx context.Context, req *CreateReq) (*User, error)
	QueryById(ctx context.Context, id int64) (*User, error)
	QueryByUsername(ctx context.Context, username string) (*User, error)
	QueryByUsernameAndPassword(ctx context.Context, username, password string) (*User, error)
	Delete(ctx context.Context, id int64) error
}

// CreateReq 创建用户请求
type CreateReq struct {
	// 用户名
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	// 密码
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}
