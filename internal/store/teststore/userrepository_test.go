package teststore_test

import (
	"bookLibrary/internal/store/teststore"
	"testing"

	"bookLibrary/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.NewTestStore()

	newUser := model.TestUser(t)
	err := s.User().Create(newUser)

	assert.NoError(t, err)
	assert.NotNil(t, newUser)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.NewTestStore()

	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	user, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "p.corepanow@gmail.com")
	//assert.Equal(t, user.password, "123456789")
}

func TestUserRepository_UpdatePasswordByEmail(t *testing.T) {
	s := teststore.NewTestStore()

	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	_, err = s.User().UpdatePasswordByEmail(
		"p.corepanow@gmail.com",
		"lol1lol2123345")
	assert.NoError(t, err)

	user, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "p.corepanow@gmail.com")
	//assert.Equal(t, user.password, "lol1lol2123345")
}

func TestUserRepository_DeleteByEmail(t *testing.T) {
	s := teststore.NewTestStore()

	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	err = s.User().DeleteByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	u, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.Nil(t, u)
	assert.Error(t, err)
}
