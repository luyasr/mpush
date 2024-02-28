package message

import (
	"context"
	kg "github.com/segmentio/kafka-go"
)

func (c *Controller) clientSend(ctx context.Context, message *Message) error {
	return c.writer.WriteMessages(ctx, kg.Message{
		Key:   []byte(Name),
		Value: []byte(message.String()),
	})
}
