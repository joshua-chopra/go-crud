package repository

import (
	"fmt"
	"github.com/joshua-chopra/go-crud/database"
	"github.com/joshua-chopra/go-crud/internal"
	"log"
)

func GetBook(bookId int) (database.Book, error) {
	var book database.Book
	// try to fetch first book (by primary key, ID) and store it in the
	// struct if found.
	if result := database.DB.First(&book, bookId); result.Error != nil {
		msg := fmt.Sprintf("Could not locate book with id: %d\n", bookId)
		log.Printf(msg)
		return book, internal.NewError(msg)
	}
	return book, nil
}

func CreateBook(book *database.Book) (*database.Book, error) {
	// attempt to create a book given the pointer to a book struct
	if result := database.DB.Create(book); result.Error != nil {
		log.Printf("Error creating book object: $v\n", result.Error)
		return book, internal.NewError("Book object could not be created.")
	} else {
		return book, result.Error
	}
}

func GetBooks() ([]database.Book, error) {
	// retrieve all books in DB if they exist.
	var allBooks []database.Book
	if res := database.DB.Find(&allBooks); res.Error != nil {
		log.Printf("Could not retrieve all books. Encountered error: %v\n", res.Error)
		return allBooks, internal.NewError("The server encountered an issue retrieving all books.")
	}
	return allBooks, nil
}

func UpdateBook(bookId int, genre string, rating int) (bool, error) {
	if book, err := GetBook(bookId); err != nil {
		msg := fmt.Sprintf("Encountered error when retrieving book with id: %d error: %v\n", bookId, err)
		log.Printf(msg)
		return true, internal.NewError(msg)
	} else {
		database.DB.Model(&book).Updates(database.Book{Genre: genre, Rating: rating})
		return false, nil
	}
}

func DeleteBook(bookId int) (bool, error) {
	if book, err := GetBook(bookId); err != nil {
		log.Println(err)
		return false, internal.NewError(fmt.Sprintf("Was unable to fetch book with id: [%d]", bookId))
	} else {
		database.DB.Delete(&book)
		return true, nil
	}
}
