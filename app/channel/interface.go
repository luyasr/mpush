package channel

import "context"

type Interface interface {
	Create(context.Context, *CreateChannelRequest) (*Channel, error)
	Validate(context.Context, *CreateChannelRequest) error
}
