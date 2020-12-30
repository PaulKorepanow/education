package server_test

import (
	"bookLibrary/internal/model"
	"bookLibrary/internal/server"
	"bookLibrary/internal/store/teststore"
	"bytes"
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
		"/users",
		bytes.NewBuffer(user),
	)
	assert.NoError(t, err)

	ts.ServeHttp(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
}
