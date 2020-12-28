package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "p.corepanow@gmail.com",
		Password: "123456789",
	}
}
