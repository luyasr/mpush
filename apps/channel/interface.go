package channel

import "context"

type Service interface {
	// Create 创建频道
	Create(ctx context.Context, req *CreateReq) (*Channel, error)
	// Update 更新频道
	Update(ctx context.Context, req *UpdateReq) error
	// Delete 删除频道
	Delete(ctx context.Context, id int64) error
	// Query 查询频道
	Query(ctx context.Context, req *QueryReq) (*Channels, error)
}

type CreateReq struct {
	Name   string `json:"name" validate:"required" label:"频道名称"`
	Url    string `json:"url"`
	Secret string `json:"secret"`
}

type UpdateReq struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Secret string `json:"secret"`
}

type QueryReq struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"keywords"`
}

func (q *QueryReq) offset() int {
	return (q.PageNumber - 1) * q.PageSize
}

type Channels struct {
	Total int64      `json:"total"`
	Items []*Channel `json:"items"`
}
