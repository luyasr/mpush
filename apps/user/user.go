package user

import "github.com/luyasr/gaia/ioc"

var _ UserService = (*UserController)(nil)

type UserController struct {

}

func init() {
	ioc.Container.Registry(name, &UserController{})
}

func (c *UserController) Init() error {
	return nil
}

func (c *UserController) Name() string {
	return name
}

func (c *UserController) Create(u *User) error {
	return nil

}

func (c *UserController) GetUserById(id int) (*User, error) {
	return nil, nil
}