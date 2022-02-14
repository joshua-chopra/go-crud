package testing_helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/database"
	"github.com/joshua-chopra/go-crud/internal"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CheckTestErrHandle(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func SendReqGetRespWriter(rtr *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	writer := httptest.NewRecorder()
	rtr.ServeHTTP(writer, req)
	log.Println(writer.Body.String())
	return writer
}

func SetupProject(envSetup bool, dbSetup bool) {
	if envSetup {
		internal.Setup()
	}
	if dbSetup {
		database.InitializeDatabase()
	}
}

func ExecuteRequest(rtr *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	respWriter := SendReqGetRespWriter(rtr, req)
	return respWriter
}

func GetBookInDB(rtr *gin.Engine, req *http.Request) (database.Book, error, *httptest.ResponseRecorder) {
	SetupProject(true, true)
	mockResp := ExecuteRequest(rtr, req)
	var parsedResp map[string]database.Book
	err := json.Unmarshal([]byte(mockResp.Body.String()), &parsedResp)
	book := parsedResp["data"]
	return book, err, mockResp
}
