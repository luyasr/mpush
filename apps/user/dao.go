package user

import (
	"context"
	"github.com/luyasr/gaia/crypto/bcrypt"
	"github.com/luyasr/gaia/errors"
	"gorm.io/gorm"
)

func (c *Controller) create(ctx context.Context, u *User) (*User, error) {
	tx := c.db.WithContext(ctx).Create(u)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (c *Controller) queryById(ctx context.Context, id int64) (*User, error) {
	user := new(User)
	tx := c.db.WithContext(ctx).Where("id = ?", id).First(user)

	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound(invalid, userNotFound, id)
		}
		return nil, err
	}

	return user, nil
}

func (c *Controller) queryByUsername(ctx context.Context, username string) (*User, error) {
	user := new(User)
	tx := c.db.WithContext(ctx).Where("username = ?", username).First(user)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound(invalid, userNotFound, username)
		}
		return nil, err
	}

	return user, nil
}

func (c *Controller) queryByUsernameAndPassword(ctx context.Context, username, password string) (*User, error) {
	user := new(User)
	tx := c.db.WithContext(ctx).Where("username = ?", username).First(user)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound(invalid, userNotFound, username)
		}
		return nil, err
	}

	if err := bcrypt.ComparePassword(password, user.Password); err != nil {
		return nil, errors.BadRequest(invalid, invalid)
	}

	return user, nil
}

func (c *Controller) delete(ctx context.Context, id int64, deletedAt int64) error {
	tx := c.db.WithContext(ctx).Model(User{}).Where("id = ?", id).Update("deleted_at", deletedAt)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
