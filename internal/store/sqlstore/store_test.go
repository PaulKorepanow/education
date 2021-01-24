package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
	logPath     string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASEURL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=postgres user=postgres password=12345678 sslmode=disable"
	}

	logPath = "../../../logs"

	os.Exit(m.Run())
}
