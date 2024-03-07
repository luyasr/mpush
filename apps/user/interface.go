package user

import "context"

type Service interface {
	Create(ctx context.Context, req *CreateReq) (*User, error)
	Query(ctx context.Context) (*User, error)
	QueryByUsername(ctx context.Context, username string) (*User, error)
	Delete(ctx context.Context, id int64) error
}

// CreateReq 创建用户请求
type CreateReq struct {
	// 用户名
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	// 密码
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

// QueryReq 查询用户请求
type QueryReq struct {
	// 查询条件
	QueryBy QueryBy `json:"query_by"`
	// 查询值
	Value string `json:"value"`
}
