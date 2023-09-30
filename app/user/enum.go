package user

type Role int

const (
	RoleMember Role = iota + 1
	RoleAdmin
)
