package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleBadRequest(c *gin.Context, err error) {
	c.AbortWithError(http.StatusBadRequest, err)
}

func ParamToInt(c *gin.Context, paramName string) (int, error) {
	paramId := c.Param(paramName)
	bookId, err := strconv.Atoi(paramId)
	if err != nil {
		msg := fmt.Sprintf(
			"Malformed id not in integer form passed: [%s] and error: %s\n", paramId, err,
		)
		log.Println(msg)
		return bookId, err
	}
	return bookId, nil
}

const NoBooksErr = "Could not retrieve any books from DB."
