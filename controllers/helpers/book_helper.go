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

func BookIdToInt(c *gin.Context) (int, error) {
	paramId := c.Param("id")
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

func BooksNotFoundErr(invalidBookId string) string {
	return fmt.Sprintf("Could not retrieve book with id: %d", invalidBookId)
}

const NoBooksErr = "Could not retrieve any books from DB."
