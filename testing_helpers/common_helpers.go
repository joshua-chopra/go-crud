package testing_helpers

import (
	"github.com/joshua-chopra/go-crud/database"
	"github.com/joshua-chopra/go-crud/internal"
	"testing"
)

func CheckTestErrHandle(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func SetupProject(envSetup bool, dbSetup bool) {
	if envSetup {
		internal.Setup()
	}
	if dbSetup {
		database.InitializeDatabase()
	}
}
