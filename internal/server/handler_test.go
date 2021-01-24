package server_test

import (
	"bookLibrary/internal/model"
	"bookLibrary/internal/server"
	"bookLibrary/internal/store/teststore"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_handleUsersCreate(t *testing.T) {
	ts := server.NewServer(teststore.NewTestStore())

	rr := httptest.NewRecorder()

	user, err := model.TestUser(t).Marshal()
	assert.NoError(t, err)

	request, err := http.NewRequest(
		http.MethodPost,
		"/api/user/new",
		bytes.NewBuffer(user),
	)
	assert.NoError(t, err)

	ts.ServeHttp(rr, request)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestHandler_handleSessions(t *testing.T) {
	testDB := teststore.NewTestStore()
	testDB.User().Create(model.TestUser(t))

	ts := server.NewServer(testDB)

	rr := httptest.NewRecorder()

	user, err := model.TestUser(t).Marshal()
	assert.NoError(t, err)

	request, err := http.NewRequest(
		http.MethodPost,
		"/api/user/login",
		bytes.NewBuffer(user),
	)
	assert.NoError(t, err)

	ts.ServeHttp(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_UpdatePassword(t *testing.T) {
	testDB := teststore.NewTestStore()
	testUser := model.TestUser(t)
	testDB.User().Create(testUser)

	testServer := server.NewServer(testDB)

	w := httptest.NewRecorder()

	newUser := model.NewUser(testUser.Email, "asdasczzvvdasaaqw1231231")
	user, err := newUser.Marshal()
	assert.NoError(t, err)

	r, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/api/user/%d/password", testUser.ID),
		bytes.NewBuffer(user),
	)
	assert.NoError(t, err)

	testServer.ServeHttp(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandler_AddBook(t *testing.T) {
	testDB := teststore.NewTestStore()
	testUser := model.TestUser(t)
	testDB.User().Create(testUser)

	testServer := server.NewServer(testDB)

	w := httptest.NewRecorder()

	user, err := testUser.Marshal()
	assert.NoError(t, err)

	r, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/api/user/%d/book", testUser.ID),
		bytes.NewBuffer(user),
	)
	assert.NoError(t, err)

	testServer.ServeHttp(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
