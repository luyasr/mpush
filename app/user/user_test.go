package user

import (
	"context"
	"testing"
)

var service = NewService()

func TestService_CreateUser(t *testing.T) {
	req := CreateUserRequest{
		Username:   "test",
		Password:   "123456",
		RePassword: "123456",
		Email:      "test@gmail.com",
	}

	user, err := service.CreateUser(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user)
}

func TestService_DeleteUser(t *testing.T) {
	req := &DeleteUserRequest{
		ID: 9,
	}
	err := service.DeleteUser(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_UpdateUser(t *testing.T) {
	req := &UpdateUserRequest{
		ID:       9,
		Nickname: "test",
	}
	err := service.UpdateUser(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}
