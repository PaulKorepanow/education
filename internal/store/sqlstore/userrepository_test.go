package sqlstore_test

import (
	"bookLibrary/internal/store/sqlstore"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"bookLibrary/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestStore(t, databaseURL, logPath)
	defer teardown("users")

	newUser := model.TestUser(t)
	err := s.User().Create(newUser)

	assert.NoError(t, err)
	assert.NotNil(t, newUser)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := sqlstore.TestStore(t, databaseURL, logPath)
	defer teardown("users")

	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	user, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "p.corepanow@gmail.com")
}

func TestUserRepository_UpdatePasswordByEmail(t *testing.T) {
	s, teardown := sqlstore.TestStore(t, databaseURL, logPath)
	defer teardown("users")

	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	_, err = s.User().UpdatePassword(
		"p.corepanow@gmail.com",
		"lol1lol2123345")
	assert.NoError(t, err)

	user, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	assert.Equal(t, user.Email, "p.corepanow@gmail.com")
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte("lol1lol2123345")))
}

func TestUserRepository_DeleteByEmail(t *testing.T) {
	s, teardown := sqlstore.TestStore(t, databaseURL, logPath)
	defer teardown("users")

	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	err = s.User().DeleteByEmail("p.corepanow@gmail.com")
	assert.NoError(t, err)

	u, err := s.User().FindByEmail("p.corepanow@gmail.com")
	assert.Nil(t, u)
	assert.EqualError(t, err, fmt.Sprintf("user with email(%s) not found", "p.corepanow@gmail.com"))
}

func TestUserRep_AddBookByEmail(t *testing.T) {
	ts, teardown := sqlstore.TestStore(t, databaseURL, logPath)
	defer teardown("users", "books")

	err := ts.User().Create(model.TestUser(t))
	assert.NoError(t, err)

	user, err := ts.User().(*sqlstore.UserRep).AddBookByEmail(model.TestUser(t).Email, "SimpleBook")
	assert.NoError(t, err)
	assert.Len(t, user.Books, 1)
}
