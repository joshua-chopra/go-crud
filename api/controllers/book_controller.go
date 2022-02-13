package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joshua-chopra/go-crud/api/controllers/helpers"
	"github.com/joshua-chopra/go-crud/api/repository"
	"github.com/joshua-chopra/go-crud/database"
	"log"
	"net/http"
	"strconv"
)

func GetAllBooks(c *gin.Context) {

	allBooks, err := repository.GetBooks()
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"data": helpers.NoBooksErr})
		return
	}
	c.IndentedJSON(
		http.StatusOK, gin.H{"data": allBooks},
	)
}

// GetOneBook /*
/*
First ensure that we can convert the incoming ID to a valid integer. If we can't
the notify the caller that there's an issue w/ the request. Otherwise, try to
fetch the book using the id, and return it if we don't encounter any errors
like book not found, etc.
*/
func GetOneBook(c *gin.Context) {
	// modifies context response if there are issues.
	bookId, err := helpers.ParamToInt(c, "id")
	if err != nil {
		helpers.HandleBadRequest(c, err)
		return
	}
	book, err := repository.GetBook(bookId)
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

func CreateBook(c *gin.Context) {
	var bookToCreate database.Book
	bookCreatePtr := &bookToCreate
	if err := c.BindJSON(bookCreatePtr); err != nil {
		log.Println("Incoming request was not valid w.r.t expected bookToCreate struct..")
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"data": err},
		)
		return
	}
	log.Printf("Incoming request body %v for bookToCreate creation: \n", bookToCreate)
	// returns pointer to created book, i.e., ptr to bookToCreate
	// which will be initialized w/ ID.
	_, err := repository.CreateBook(bookCreatePtr)
	if err != nil {
		log.Printf("Issue creating bookToCreate: %s", err)
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"data": fmt.Sprintf("%v bookToCreate could not be created.\n", bookToCreate)})
		return
	}

	c.IndentedJSON(
		http.StatusCreated,
		gin.H{"data": bookToCreate},
	)
}

func UpdateBook(c *gin.Context) {
	// retrieve ID
	bookId, err := helpers.ParamToInt(c, "id")
	if err != nil {
		helpers.HandleBadRequest(c, err)
		return
	}
	newGenre := c.Query("genre")
	newRating, err := strconv.Atoi(c.Query("rating"))
	if err != nil {
		helpers.HandleBadRequest(c, err)
		return
	}
	if _, err := repository.UpdateBook(bookId, newGenre, newRating); err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": err})
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"data": "Updated book successfully."},
	)
}

func DeleteBook(c *gin.Context) {
	// ensure proper id was passed in as path param
	bookId, err := helpers.ParamToInt(c, "id")
	if err != nil {
		helpers.HandleBadRequest(c, err)
		return
	}
	if _, err := repository.DeleteBook(bookId); err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"data": err})
		return
	}
	c.IndentedJSON(
		http.StatusNoContent,
		gin.H{"data": fmt.Sprintf("Successfully deleted book with id: %d", bookId)},
	)

}
