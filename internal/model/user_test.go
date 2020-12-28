package model_test

import (
	"bookLibrary/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreation(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreation())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name  string
		user  func() model.User
		isErr bool
	}{
		{"invalid email",
			func() model.User {
				u := model.TestUser(t)
				u.Email = "McPaha"
				return *u
			},
			true,
		},
		{"invalid password",
			func() model.User {
				u := model.TestUser(t)
				u.Password = "1234"
				return *u
			},
			true,
		},
		{
			"right case",
			func() model.User {
				return *model.TestUser(t)
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isErr {
				assert.Error(t, tt.user().Validate())
			} else {
				assert.NoError(t, tt.user().Validate())
			}
		})
	}
}
