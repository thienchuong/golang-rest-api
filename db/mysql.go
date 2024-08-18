package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thienchuong/golang-rest-api/config"
	"github.com/thienchuong/golang-rest-api/models"
)

type mysql struct {
	db *sql.DB
}

func NewMysqlDb() IDatabase {
	mysqlConfig := config.Get().Database.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}
	return &mysql{
		db: db,
	}
}

func (m *mysql) GetAllBooks() ([]models.Book, error) {
	rows, err := m.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (m *mysql) GetBookByID(id int) (models.Book, error) {
	var book models.Book
	err := m.db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (m *mysql) CreateBook(book models.Book) (models.Book, error) {
	result, err := m.db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", book.Title, book.Author, book.Year)
	if err != nil {
		return models.Book{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Book{}, err
	}

	book.ID = int(id)
	return book, nil
}

func (m *mysql) UpdateBook(id int, book models.Book) (models.Book, error) {
	_, err := m.db.Exec("UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?", book.Title, book.Author, book.Year, id)
	if err != nil {
		return models.Book{}, err
	}

	book.ID = id
	return book, nil
}

func (m *mysql) DeleteBook(id int) error {
	_, err := m.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

// Close closes the database connection
func (m *mysql) Close() {
	m.db.Close()
}
