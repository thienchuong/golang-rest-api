package main

import (
	"net/http"

	"github.com/thienchuong/go-rest-api/db"
	"github.com/thienchuong/go-rest-api/handlers"
	"github.com/thienchuong/go-rest-api/log"
)

func main() {
	log := log.NewConsoleLogger("myapp")

	log.Info("This is an info message")

	db := db.NewPostgresDB()
	// or db := db.NewMySQLDatabase()

	handler := handlers.NewHandler(db, log)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", handler.GetAllBooks)
	mux.HandleFunc("GET /books/{id}", handler.GetBookByID)
	mux.HandleFunc("POST /books", handler.CreateBook)
	mux.HandleFunc("PUT /books/{id}", handler.UpdateBook)
	mux.HandleFunc("DELETE /books/{id}", handler.DeleteBook)

	log.Info("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
