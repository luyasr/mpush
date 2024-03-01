package message

import (
	"context"
	kg "github.com/segmentio/kafka-go"
)

func (c *Controller) producer(ctx context.Context, message *Message) error {
	return c.writer.WriteMessages(ctx, kg.Message{
		Key:   []byte(Name),
		Value: []byte(message.String()),
	})
}

func (c *Controller) create(ctx context.Context, message *Message) error {
	return c.db.WithContext(ctx).Create(message).Error
}
