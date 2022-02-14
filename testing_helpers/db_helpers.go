package testing_helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/database"
	"net/http/httptest"
)

func GetBookInDB(rtr *gin.Engine, id uint) (database.Book, error, *httptest.ResponseRecorder) {
	SetupProject(true, true)
	req := GetBookRequest(id)
	// simulates client making call to server
	// which makes DB call to get book by id.
	mockResp := ExecuteRequest(rtr, req)
	book, err := unmarshalBookJSON(mockResp)
	return book, err, mockResp
}

func unmarshalBookJSON(resp *httptest.ResponseRecorder) (database.Book, error) {
	var parsedResp map[string]database.Book
	err := json.Unmarshal([]byte(resp.Body.String()), &parsedResp)
	book := parsedResp["data"]
	return book, err
}
