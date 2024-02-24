package user

import (
	"encoding/json"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   Status `json:"status"`
	Role     Role   `json:"role"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) String() string {
	bytes, _ := json.Marshal(u)
	return string(bytes)
}
