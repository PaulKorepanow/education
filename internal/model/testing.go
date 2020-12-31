package model

import (
	"testing"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "p.corepanow@gmail.com",
		Password: "123456789",
	}
}

func TestBook(t *testing.T) *Book {
	return &Book{
		Title: "SimpleBook",
	}
}
