package router

import (
	"net/http"
)

func NewRouter() {
	// Create a new router
	router := http.NewServeMux()

	router.HandleFunc("GET /health/", health.getHealth)

	router.HandleFunc("GET /books/", book.getBooks)
}
