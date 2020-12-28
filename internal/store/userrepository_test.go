package store_test

import (
	"errors"
	"testing"

	"bookLibrary/internal/model"
	"bookLibrary/internal/store"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	user, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	user, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "p.corepanow@gmail.com")
	//assert.Equal(t, user.Password, "123456789")
}

func TestUserRepository_UpdatePasswordByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	_, err = s.User().UpdatePasswordByEmail(
		"p.corepanow@gmail.com",
		"lol1lol2123345")
	assert.NoError(t, err)

	user, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "p.corepanow@gmail.com")
	//assert.Equal(t, user.Password, "lol1lol2123345")
}

func TestUserRepository_DeleteByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	err = s.User().DeleteByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	u, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.Nil(t, u)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
