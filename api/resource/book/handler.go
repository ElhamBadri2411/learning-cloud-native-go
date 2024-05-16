package book

import "net/http"

type API struct{}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Get all books
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	// Get a book by id
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// Update a book
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Delete a book
}

func createBook(w http.ResponseWriter, r *http.Request) {
	// Create a books
}
