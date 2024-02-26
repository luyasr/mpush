package token

import (
	"context"
	"github.com/luyasr/gaia/errors"
	"gorm.io/gorm"
)

func (c *Controller) findByUserId(ctx context.Context, userId int64) (*Token, error) {
	token := new(Token)
	tx := c.db.WithContext(ctx).Where("user_id = ?", userId).First(token)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound(tokenInvalid, tokenNotFound, userId)
		}
		return nil, err
	}
	return token, nil
}

func (c *Controller) findByToken(ctx context.Context, token string) (*Token, error) {
	tk := new(Token)
	tx := c.db.WithContext(ctx).Where("token = ?", token).First(tk)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound(tokenInvalid, tokenNotFound, token)
		}
		return nil, err
	}
	return tk, nil
}

func (c *Controller) update(ctx context.Context, token *Token) error {
	return c.db.WithContext(ctx).Model(token).Updates(token).Error
}

func (c *Controller) login(ctx context.Context, token *Token) (*Tk, error) {
	tx := c.db.WithContext(ctx).Create(token)
	if err := tx.Error; err != nil {
		return nil, err
	}

	tk := &Tk{
		Token:        token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return tk, nil
}

func (c *Controller) refresh(ctx context.Context, token *Token) (string, error) {
	return "", nil
}
