package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book represents a book in the library.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var nextID = 1

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler to get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Handler to get a single book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	book.ID = nextID
	nextID++
	books = append(books, book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// Handler to update a book by ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// Handler to delete a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
