package repository

import (
	"fmt"
	"github.com/joshua-chopra/go-crud/internal"
	"github.com/joshua-chopra/go-crud/models"
	"log"
)

func GetBook(bookId int) (models.Book, error) {
	var book models.Book
	// try to fetch first book (by primary key, ID) and store it in the
	// struct if found.
	if result := internal.DB.First(&book, bookId); result.Error != nil {
		fmt.Printf("Could not locate book with id: %d\n", bookId)
		return book, result.Error
	}
	return book, nil
}

func CreateBook(book *models.Book) error {
	// attempt to create a book given the pointer to a book struct
	if result := internal.DB.Create(book); result.Error != nil {
		log.Printf("Error creating book object: $v\n", result.Error)
		return nil
	} else {
		return result.Error
	}
}

func GetBooks() ([]models.Book, error) {
	// retrieve all books in DB if they exist.
	var allBooks []models.Book
	if res := internal.DB.Find(&allBooks); res.Error != nil {
		log.Printf("Could not retrieve all books. Encountered error: %v\n", res.Error)
		return allBooks, res.Error
	}
	return allBooks, nil
}

func UpdateBook() {

}

func DeleteBook(bookId int) error {
	// retrieve book first, if we are unable to then
	// return error to caller, otherwise return
	// nil to indicate deleted.
	if book, err := GetBook(bookId); err != nil {
		return err
	} else {
		internal.DB.Delete(&book)
	}
	return nil
}
