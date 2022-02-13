package api

import (
	"encoding/json"
	"github.com/joshua-chopra/go-crud/database"
	"github.com/joshua-chopra/go-crud/internal"
	"github.com/stretchr/testify/assert"
	"log"
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
	log.Println(writer.Body.String())

	// response is of the form "data": [ array of book objects ]
	var resp map[string][]database.Book
	err = json.Unmarshal([]byte(writer.Body.String()), &resp)

	books := resp["data"]
	//fmt.Println(books)

	assert.Equal(t, 200, writer.Code)
	assert.Nil(t, err)
	// 2 books after initial seeding phase.
	assert.Equal(t, len(books), 2)
}
