package domains

import (
	"github.com/DenisKnez/simpleWebGolang/data"
)

//BookRepository book repository interface
type BookRepository interface {
	GetBookByID(id string) (book data.Book, err error)
	Books() (books []data.Book, err error)
	PagedBooks(pageSize int, pageNumber int) (books []data.Book, err error)
	CreateBook(book data.Book) (err error)
	DeleteBook(id string) (err error)
	UpdateBook(book data.Book) (err error)
}

//BookUseCase book useCase interface
type BookUseCase interface {
	GetBookByID(id string) (book data.Book, err error)
	Books() (books []data.Book, err error)
	PagedBooks(pageSize int, pageNumber int) (books []data.Book, err error)
	CreateBook(book *data.Book) (err error)
	DeleteBook(id string) (err error)
	UpdateBook(book *data.Book) (err error)
}
