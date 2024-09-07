package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/naren/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //Automatically migrate your schema, to keep your schema up to date. NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
}

//functions and methods to talk to the database

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	if db.RecordNotFound() {
		fmt.Printf("No book found with ID: %d\n", Id)
	}
	if db.Error != nil {
		fmt.Printf("Error finding book with ID: %d, error: %v\n", Id, db.Error)
	}
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}

//Model gives the structure to store the information in the database
