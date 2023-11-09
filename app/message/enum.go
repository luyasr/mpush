package message

type Status int

const (
	StatusPending Status = iota
	StatusSent
	StatusFailed
)
