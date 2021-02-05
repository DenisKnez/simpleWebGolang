package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/DenisKnez/simpleWebGolang/util"
)

//bookRepository book repository
type bookRepository struct {
	conn   *sql.DB
	logger *log.Logger
}

//NewBookRepository create an new instance of the book repository
func NewBookRepository(conn *sql.DB, logger *log.Logger) domains.BookRepository {
	return &bookRepository{conn, logger}
}

//GetBookByID get the book by its id
func (bookRepo *bookRepository) GetBookByID(id string) (book data.Book, err error) {
	book = data.Book{}
	err = util.Db.QueryRow("SELECT id, title, author, release_date, created_at, updated_at, deleted_at FROM books WHERE id = $1 AND is_deleted = false", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.ReleaseDate, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)

	return
}

//Books return all the books
func (bookRepo *bookRepository) Books() (books []data.Book, err error) {
	rows, err := bookRepo.conn.Query(
		`SELECT 
		id, title, author, release_date, created_at, updated_at, deleted_at, is_deleted
		FROM books
			WHERE is_deleted = false
	`)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		book := data.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ReleaseDate, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.IsDeleted)

		if err != nil {
			return
		}

		books = append(books, book)
	}

	return
}

//PagedBooks return books based on parameters
func (bookRepo *bookRepository) PagedBooks(pageSize int, pageNumber int) (books []data.Book, err error) {

	offset := (pageSize * pageNumber - pageSize)

	rows, err := bookRepo.conn.Query(
		`SELECT 
			id, title, author, release_date, created_at, updated_at, deleted_at, is_deleted
		FROM books
			WHERE is_deleted = false
		LIMIT $1
		OFFSET $2
		`, pageSize, offset)


	fmt.Println(rows)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		book := data.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ReleaseDate, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.IsDeleted)

		if err != nil {
			fmt.Println(err)
			return
		}

		books = append(books, book)
	}

	return
}

//CreateBook create a new book
func (bookRepo *bookRepository) CreateBook(book data.Book) (err error) {

	stmt, err := bookRepo.conn.Prepare(
		`INSERT INTO books 
			(id, title, author, release_date, created_at, updated_at, deleted_at, is_deleted )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title, author, release_date, created_at, updated_at, deleted_at, is_deleted
		`)

	if err != nil {
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(book.ID, book.Title, book.Author, book.ReleaseDate, book.CreatedAt, book.UpdatedAt, book.DeletedAt, book.IsDeleted)

	if err != nil {
		return
	}

	bookRepo.logger.Println(result)

	return
}

//DeleteBook delete the book with the provided id
func (bookRepo *bookRepository) DeleteBook(id string) (err error) {

	stmt, err := bookRepo.conn.Prepare("DELETE FROM books WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

