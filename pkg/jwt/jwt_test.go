package jwt

import "testing"

var secret = "123456"

func TestGenerateJwt(t *testing.T) {
	jwt, err := GenerateJwt("admin", secret)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jwt)
}

func TestParseJwt(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjk5MzQ0MDg5LCJuYmYiOjE2OTkyNTc2ODksImlhdCI6MTY5OTI1NzY4OX0.MEsvjbUnJQwtYmIpks7NRRzsY2rmAmswW1jQtZVRip8"
	jwt, err := ParseJwt(token, secret)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jwt)
}
