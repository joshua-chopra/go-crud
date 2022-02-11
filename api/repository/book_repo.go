package repository

import (
	"fmt"
	"github.com/joshua-chopra/go-crud/database"
	"log"
)

func GetBook(bookId int) (database.Book, error) {
	var book database.Book
	// try to fetch first book (by primary key, ID) and store it in the
	// struct if found.
	if result := database.DB.First(&book, bookId); result.Error != nil {
		fmt.Printf("Could not locate book with id: %d\n", bookId)
		return book, result.Error
	}
	return book, nil
}

func CreateBook(book *database.Book) error {
	// attempt to create a book given the pointer to a book struct
	if result := database.DB.Create(book); result.Error != nil {
		log.Printf("Error creating book object: $v\n", result.Error)
		return nil
	} else {
		return result.Error
	}
}

func GetBooks() ([]database.Book, error) {
	// retrieve all books in DB if they exist.
	var allBooks []database.Book
	if res := database.DB.Find(&allBooks); res.Error != nil {
		log.Printf("Could not retrieve all books. Encountered error: %v\n", res.Error)
		return allBooks, res.Error
	}
	return allBooks, nil
}

func UpdateBook(bookId int, genre string, rating int) {
	if book, err := GetBook(bookId); err != nil {
		log.Printf("Encountered error when retrieving book with id: %d error: %v\n", bookId, err)
	} else {
		database.DB.Model(&book).Updates(database.Book{Genre: genre, Rating: rating})
	}
}

func DeleteBook(bookId int) error {
	// retrieve book first, if we are unable to then
	// return error to caller, otherwise return
	// nil to indicate deleted.
	if book, err := GetBook(bookId); err != nil {
		return err
	} else {
		database.DB.Delete(&book)
	}
	return nil
}
