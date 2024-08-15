package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/thienchuong/go-rest-api/config"
	"github.com/thienchuong/go-rest-api/models"
)

type postgres struct {
	conn *pgx.Conn
}

// NewPostgresDB creates a new postgres database
func NewPostgresDB() IDatabase {
	postgresConfig := config.Get().Database.Postgresql

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", postgresConfig.Username, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(fmt.Sprintf("Error opening database: %v", err))
	}

	return &postgres{
		conn: conn,
	}
}

func (p *postgres) GetAllBooks() ([]models.Book, error) {
	rows, err := p.conn.Query(context.Background(), "SELECT * FROM books")
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

func (p *postgres) GetBookByID(id int) (models.Book, error) {
	var book models.Book
	err := p.conn.QueryRow(context.Background(), "SELECT * FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (p *postgres) CreateBook(book models.Book) (models.Book, error) {
	_, err := p.conn.Exec(context.Background(), "INSERT INTO books (title, author, year) VALUES ($1, $2, $3)", book.Title, book.Author, book.Year)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (p *postgres) UpdateBook(id int, book models.Book) (models.Book, error) {
	_, err := p.conn.Exec(context.Background(), "UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4", book.Title, book.Author, book.Year, id)
	if err != nil {
		return models.Book{}, err
	}

	book.ID = id
	return book, nil
}

func (p *postgres) DeleteBook(id int) error {
	_, err := p.conn.Exec(context.Background(), "DELETE FROM books WHERE id = $1", id)
	return err
}

func (p *postgres) Close() {
	p.conn.Close(context.Background())
}
