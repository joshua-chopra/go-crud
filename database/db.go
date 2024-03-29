package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitializeDatabase() {
	dbConn := connectDB()
	seedDB(dbConn)
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	log.Printf("DB URL: %s", dsn)
	// create a new connection w/ gorm.Open method
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("Failed to connect to database!")
	}
	// run migrations i.e., load table for given model if not exists
	// or any schema changes, i.e., if new field is added to book model
	// then the migration will be run and table is updated.
	// if we had multiple database, we'd call it on each model.
	if err = db.AutoMigrate(&Book{}); err != nil {
		log.Fatalf("encountered error running migrations for Book model: %v\n", err)
	}
	// otherwise, assign DB object, and we can use this as needed
	// in controllers or routes when exported from this file.
	DB = db
	return db
}

func seedDB(db *gorm.DB) {
	// seed database with 2 books to begin with
	db.Create(
		&Book{
			ID:     1,
			Title:  "Harry Potter",
			Author: "JK Rowling",
			Genre:  "Fiction",
			Rating: 9,
		},
	)
	db.Create(
		&Book{
			ID:     2,
			Title:  "Tom Sawyer",
			Author: "Mark Twain",
			Genre:  "Folklore",
			Rating: 8,
		},
	)
	log.Println("No errors, initialized DB and applied migrations.")
}
