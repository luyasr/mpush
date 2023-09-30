package user

import (
	"context"
	"testing"
)

var service = NewService()

func TestService_Create(t *testing.T) {
	req := CreateUserRequest{
		Username:   "test",
		Password:   "123456",
		RePassword: "123456",
		Email:      "test@gmail.com",
	}

	user, err := service.Create(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}

func TestService_DeleteById(t *testing.T) {
	err := service.DeleteById(context.Background(), 9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_Update(t *testing.T) {
	req := UpdateUserRequest{
		Nickname: "test",
	}
	err := service.Update(context.Background(), 9, &req)
	if err != nil {
		t.Fatal(err)
	}
}
