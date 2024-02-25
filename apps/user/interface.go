package user

import "context"

type Service interface {
	Create(ctx context.Context, req *CreateReq) (*User, error)
	Find(ctx context.Context, req *FindReq) (*User, error)
	Delete(ctx context.Context, id int64) error
}

// CreateReq 创建用户请求
type CreateReq struct {
	// 用户名
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	// 密码
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

// FindReq 查询用户请求
type FindReq struct {
	// 查询条件
	FindBy FindBy `json:"find_by"`
	// 查询值
	Value string `json:"value"`
}
