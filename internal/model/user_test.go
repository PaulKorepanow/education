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
