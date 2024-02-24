package user

const (
	name = "user"
)

type UserService interface {
	Create(u *User) error
	GetUserById(id int) (*User, error)
}