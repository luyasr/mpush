package channel

import (
	"context"
)

func (c *Controller) create(ctx context.Context, ch *Channel) (*Channel, error) {
	if err := c.db.WithContext(ctx).Create(ch).Error; err != nil {
		return nil, err
	}

	return ch, nil
}

func (c *Controller) update(ctx context.Context, m map[string]any) error {
	return c.db.WithContext(ctx).Model(&Channel{}).Updates(m).Error
}

func (c *Controller) delete(ctx context.Context, id int64) error {
	return c.db.WithContext(ctx).Delete(&Channel{}, id).Error
}
