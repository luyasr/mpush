package token

import (
	"encoding/json"
	"github.com/rs/xid"
	"time"
)

type Token struct {
	Id                    int64  `json:"id" gorm:"primaryKey"`
	UserId                int64  `json:"userId" gorm:"not null;uniqueIndex"`
	AccessToken           string `json:"access_token" gorm:"not null;uniqueIndex;type:varchar(50)"`
	RefreshToken          string `json:"refresh_token" gorm:"not null;uniqueIndex;type:varchar(50)"`
	AccessTokenExpiredAt  int64  `json:"access_token_expired_at" gorm:"not null"`
	RefreshTokenExpiredAt int64  `json:"refresh_token_expired_at" gorm:"not null"`
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) String() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}

func (t *Token) Refresh() {
	t.AccessToken = xid.New().String()
}

func NewDefaultToken() *Token {
	return &Token{
		AccessToken:           xid.New().String(),
		RefreshToken:          xid.New().String(),
		AccessTokenExpiredAt:  time.Now().Add(2 * time.Hour).UnixMilli(),
		RefreshTokenExpiredAt: time.Now().Add(168 * time.Hour).UnixMilli(),
	}
}
