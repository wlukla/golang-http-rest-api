package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	database := os.Getenv("DATABASE_URL")

	if database == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
