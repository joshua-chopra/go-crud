package api

import (
	"github.com/joshua-chopra/go-crud/database"
	"github.com/joshua-chopra/go-crud/internal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	rtr, _ := setupRouter()

	testReq, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	rtr.ServeHTTP(writer, testReq)

	assert.Equal(t, 200, writer.Code)
	expectedResp := `{"message":"pong"}`
	assert.Equal(t, expectedResp, writer.Body.String())
}

func TestGetAllBooks(t *testing.T) {
	internal.Setup()
	database.InitializeDatabase()
	rtr, _ := setupRouter()
	testReq, err := http.NewRequest("GET", "/api/book/", nil)
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	rtr.ServeHTTP(writer, testReq)
	assert.Equal(t, 200, writer.Code)
}
