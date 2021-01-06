package test

import (
	"os"
	"testing"

	"github.com/black-lotus/go-test/app"
	"github.com/black-lotus/go-test/utils/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
