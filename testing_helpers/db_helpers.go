package testing_helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/database"
	"net/http"
	"net/http/httptest"
)

func GetBookInDB(rtr *gin.Engine, req *http.Request) (database.Book, error, *httptest.ResponseRecorder) {
	SetupProject(true, true)
	mockResp := ExecuteRequest(rtr, req)
	var parsedResp map[string]database.Book
	err := json.Unmarshal([]byte(mockResp.Body.String()), &parsedResp)
	book := parsedResp["data"]
	return book, err, mockResp
}
