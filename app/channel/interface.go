package channel

import "context"

type Interface interface {
	CreateChannel(context.Context, *Request) (*Channel, error)
}
