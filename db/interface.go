package db

import "github.com/thienchuong/golang-rest-api/models"

type IDatabase interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id int) (models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	UpdateBook(id int, book models.Book) (models.Book, error)
	DeleteBook(id int) error
}
