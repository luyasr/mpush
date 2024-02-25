package token

import (
	"encoding/json"
	"github.com/rs/xid"
	"time"
)

const (
	Name = "token"
)

type Token struct {
	// ID token ID
	Id int64 `json:"id"`
	// UserId 用户ID
	UserId int64 `json:"user_id"`
	// Token token
	Token string `json:"token"`
	// RefreshToken 刷新token
	RefreshToken string `json:"refresh_token"`
	// ExpireTime 过期时间
	ExpireTime int64 `json:"expire_time"`
	// RefreshTime 刷新时间
	RefreshTime int64 `json:"refresh_time"`
	// CreatedAt 创建时间
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt int64 `json:"updated_at"`
	// DeletedAt 删除时间
	DeletedAt int64 `json:"deleted_at"`
}

func (Token) TableName() string {
	return Name
}

func (t *Token) String() string {
	bytes, _ := json.Marshal(t)

	return string(bytes)
}

func NewToken(userId int64) *Token {
	now := time.Now().Unix()
	return &Token{
		UserId:       userId,
		Token:        xid.New().String(),
		RefreshToken: xid.New().String(),
		ExpireTime:   now + 7200,
		RefreshTime:  now + 604800,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
