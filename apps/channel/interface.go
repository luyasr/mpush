package channel

type Service interface {
	// CreateChannel 创建频道
	CreateChannel(channel *Channel) error
}
