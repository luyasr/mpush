package token

import (
	"context"
	"testing"
)

var service = NewService()

func TestService_Login(t *testing.T) {
	req := &LoginRequest{Username: "test", Password: "123456"}
	token, err := service.Login(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestService_Logout(t *testing.T) {
	req := &LogoutRequest{UserId: 5, AccessToken: "ckchgmhqd2tr26d2uql0", RefreshToken: "ckchgmhqd2tr26d2uqlg"}
	err := service.Logout(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}
