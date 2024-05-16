package book

import "net/http"

type API struct{}

func (api *API) GetBooks(w http.ResponseWriter, r *http.Request) {
	// Get all books
}

func (api *API) GetBookById(w http.ResponseWriter, r *http.Request) {
	// Get a book by id
}

func (api *API) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Update a book
}

func (api *API) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Delete a book
}

func (api *API) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Create a book
}
