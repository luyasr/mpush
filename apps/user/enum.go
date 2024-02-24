package user

type Status int

const (
	StatusNormal Status = iota
	StatusDeleted
)

type Role int

const (
	RoleAnonymous Role = iota
	RoleAdmin
)
