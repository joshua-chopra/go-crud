package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/internal"
	"github.com/joshua-chopra/go-crud/models"
	"log"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	var allBooks []models.Book
	internal.DB.Find(&allBooks)
	c.IndentedJSON(
		http.StatusOK, gin.H{"data": allBooks},
	)
}

// GetBook /*
/*
First ensure that we can convert the incoming ID to a valid integer. If we can't
the notify the caller that there's an issue w/ the request. Otherwise, try to
fetch the book using the id, and return it if we don't encounter any errors
like book not found, etc.
*/
func GetBook(c *gin.Context) {
	idParam := c.Param("id")
	bookId, err := bookIdToInt(idParam)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"data": fmt.Sprintf(
				"Malformed id not in integer form passed: [%s] and error: \n", idParam),
			},
		)
		return
	}
	book, err := getOne(bookId)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"data": fmt.Sprintf("Unable to locate book with id: %d", bookId)},
		)
		return
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"data": book},
	)
}

func getOne(bookId int) (models.Book, error) {
	var book models.Book
	// pass in destination struct to avoid referring to model. If we encounter an
	// error searching for the book, we'll return the error and the uninitialized
	// struct to the caller.
	if result := internal.DB.First(&book, bookId); result.Error != nil {
		fmt.Printf("Could not locate book with id: %d\n", bookId)
		return book, result.Error
	}
	return book, nil
}

func bookIdToInt(id string) (int, error) {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error converting %s to integer, will not execute search for book\n Error: %v\n",
			id, err,
		)
		return -1, err
	}
	return bookId, nil
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		log.Println("Incoming request was not valid w.r.t expected book struct..")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Printf("Incoming request body for book creation: \n", book)
	if result := internal.DB.Create(&book); result.Error != nil {
		log.Printf("Error creating book object: $v\n", result.Error)
		return
	}
	c.IndentedJSON(
		http.StatusCreated,
		gin.H{"data": book},
	)
}

func UpdateBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
