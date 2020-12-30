package app_test

import (
	"bookLibrary/internal/app"
	"bookLibrary/internal/store/teststore"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_handleUsersCreate(t *testing.T) {
	var conf = app.Config{
		BindAddr: ":8080",
		LogLevel: "info",
		Store:    "host=localhost dbname=postgres user=postgres password=12345678 sslmode=disable",
	}

	ts := app.NewServer(&conf, teststore.NewTestStore())

	rr := httptest.NewRecorder()

	var jsonData string = "{\n  \"email\": \"p.corepanow@gmail.com\",\n  \"password\": \"123456789\"\n}\n"

	request, err := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(jsonData)))
	assert.NoError(t, err)

	ts.ServerHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
}
