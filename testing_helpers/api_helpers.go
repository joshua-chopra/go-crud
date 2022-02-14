package testing_helpers

import (
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
	respWriter := SendReqGetRespWriter(rtr, req)
	return respWriter
}
