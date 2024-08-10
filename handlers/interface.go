package handlers

import "net/http"

type IHandler interface {
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetBookByID(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
