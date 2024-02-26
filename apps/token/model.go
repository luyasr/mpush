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
	AccessToken string `json:"access_token"`
	// RefreshToken 刷新token
	RefreshToken string `json:"refresh_token"`
	// ExpireTime 过期时间
	AccessTokenExpiredAt int64 `json:"access_token_expired_at"`
	// RefreshTime 刷新时间
	RefreshTokenExpiredAt int64 `json:"refresh_token_expired_at"`
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

func (t *Token) IsExpired() bool {
	return time.Now().Unix() > t.AccessTokenExpiredAt
}

func (t *Token) IsRefreshExpired() bool {
	return time.Now().Unix() > t.RefreshTokenExpiredAt
}

func NewToken(userId int64) *Token {
	now := time.Now().Unix()
	return &Token{
		UserId:                userId,
		AccessToken:           xid.New().String(),
		RefreshToken:          xid.New().String(),
		AccessTokenExpiredAt:  now + 7200,
		RefreshTokenExpiredAt: now + 604800,
		CreatedAt:             now,
		UpdatedAt:             now,
	}
}
