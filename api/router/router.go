package router

import (
	"myapp/api/resource/book"
	"myapp/api/resource/health"
	"net/http"
)

func NewRouter() *http.ServeMux {
	// Create a new router
	router := http.NewServeMux()

	router.HandleFunc("GET /health/", health.GetHealth)

	bookAPI := &book.API{}
	router.HandleFunc("/books/", bookAPI.GetBooks)
	router.HandleFunc("/books/{id}", bookAPI.GetBookById)
	router.HandleFunc("/books/update", bookAPI.UpdateBook)
	router.HandleFunc("/books/delete", bookAPI.DeleteBook)
	router.HandleFunc("/books/create", bookAPI.CreateBook)

	return router
}
