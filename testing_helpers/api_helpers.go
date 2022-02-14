package testing_helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
)

func SendReqGetRespWriter(rtr *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	writer := httptest.NewRecorder()
	rtr.ServeHTTP(writer, req)
	log.Println(writer.Body.String())
	return writer
}

func ExecuteRequest(rtr *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	log.Printf("Request URL is: %v", req.URL)
	respWriter := SendReqGetRespWriter(rtr, req)
	return respWriter
}

func GetBookRequest(id uint) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/book/%d", id), nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}
