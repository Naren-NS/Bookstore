package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	"github.com/naren/go-bookstore/pkg/models"
	"github.com/naren/go-bookstore/pkg/utils"
)

var NewBook models.Book
//Waitgroups help with synchronizing a group of go routines

func GetBook(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup // Initialize a wait group
	bookChannel := make(chan []models.Book) // Create a channel for sending books

	wg.Add(1) // Increment the wait group counter
	go func() {
		defer wg.Done()  // Ensure wg.Done() is called when the function exits
		newBooks := models.GetAllBooks() // Fetch all books
		bookChannel <- newBooks // Send the books to the channel
	}()

	go func() {
		wg.Wait() // Wait for all go routines to finish
		close(bookChannel)  // Close the channel when done
	}()

	var allBooks []models.Book
	for books := range bookChannel { // Receive books from the channel
		allBooks = append(allBooks, books...) // Append books to the list
	}

	res, err := json.Marshal(allBooks)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
