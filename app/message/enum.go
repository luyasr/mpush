package message

import "github.com/luyasr/mpush/pkg/errorhandler"

type Status int

const (
	StatusPending Status = iota
	StatusSent
	StatusFailed
)

var (
	ErrUpdateMessageFailed = errorhandler.NewUpdateFailed("update message failed")
	ErrDeleteMessageFailed = errorhandler.NewDeleteFailed("delete message failed")
)
