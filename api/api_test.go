package api

import (
	"encoding/json"
	"github.com/joshua-chopra/go-crud/database"
	tHelp "github.com/joshua-chopra/go-crud/testing_helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	rtr, _ := SetupRouter()

	testReq, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	rtr.ServeHTTP(writer, testReq)

	assert.Equal(t, http.StatusOK, writer.Code)
	expectedResp := `{"message":"pong"}`
	assert.Equal(t, expectedResp, writer.Body.String())
}

func TestGetAllBooks(t *testing.T) {
	tHelp.SetupProject(true, true)
	rtr, _ := SetupRouter()
	testReq, err := http.NewRequest("GET", "/api/book/", nil)
	tHelp.CheckTestErrHandle(t, err)

	respWriter := tHelp.SendReqGetRespWriter(rtr, testReq)
	// response is of the form "data": [ array of book objects ]
	var resp map[string][]database.Book
	err = json.Unmarshal([]byte(respWriter.Body.String()), &resp)

	books := resp["data"]
	//fmt.Println(books)
	assert.Equal(t, http.StatusOK, respWriter.Code)
	assert.Nil(t, err)
	// 2 books after initial seeding phase.
	assert.Equal(t, len(books), 2)
}

func TestGetBookByIDInDB(t *testing.T) {
	// id field is of type unsigned int
	var bookID uint = 1
	rtr, _ := SetupRouter()
	book, err, resp := tHelp.GetBookInDB(rtr, bookID)

	assert.Nil(t, err)
	assert.Equal(t, bookID, book.ID)
	assert.Equal(t, resp.Code, http.StatusOK)
}

func TestGetBookByIDNotInDB(t *testing.T) {
	var bookID uint = 999
	rtr, _ := SetupRouter()
	book, err, resp := tHelp.GetBookInDB(rtr, bookID)

	assert.NotNil(t, err)
	assert.True(t, book.IsEmpty())
	assert.Equal(t, resp.Code, http.StatusNotFound)
}
