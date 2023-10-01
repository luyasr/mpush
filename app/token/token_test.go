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
	req := &Request{UserId: 5, AccessToken: "ckcjcj9qd2ttspogiuk0", RefreshToken: "ckcjcj9qd2ttspogiukg"}
	err := service.Logout(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_Refresh(t *testing.T) {
	req := &Request{UserId: 5, AccessToken: "ckcjimpqd2tpehnm7ejg", RefreshToken: "ckcjhnhqd2ts95c6vqj0"}
	err := service.Refresh(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestService_Validate(t *testing.T) {
	req := &Request{UserId: 5, AccessToken: "ckcjcj9qd2ttspogiuk0", RefreshToken: "ckcjcj9qd2ttspogiukg"}
	err := service.Validate(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}
