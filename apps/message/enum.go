package message

type Channel int

const (
	// ChannelDingTalk 钉钉消息
	ChannelDingTalk Channel = iota
	// ChannelWeChat 微信消息
	ChannelWeChat
	// ChannelEmail 邮件消息
	ChannelEmail
	// ChannelSlack Slack消息
	ChannelSlack
	// ChannelSMS 短信消息
	ChannelSMS
)

type Status int

const (
	// StatusUnsent 未发送
	StatusUnsent Status = iota
	// StatusSent 已发送
	StatusSent
	// StatusFailed 发送失败
	StatusFailed
)
