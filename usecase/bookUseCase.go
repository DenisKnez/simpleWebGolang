package usecase

import (
	"log"
	"time"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/DenisKnez/simpleWebGolang/util"
)

type bookUseCase struct {
	bookRepo domains.BookRepository
	logger   *log.Logger
}

//NewBookUseCase returns new book useCase
func NewBookUseCase(repo domains.BookRepository, logger *log.Logger) domains.BookUseCase {
	return &bookUseCase{repo, logger}
}

func (bookUC *bookUseCase) GetBookByID(id string) (book data.Book, err error) {

	book, err = bookUC.bookRepo.GetBookByID(id)

	if err != nil {
		bookUC.logger.Printf("method GetBookByID | %s", err)
	}

	return
}

func (bookUC *bookUseCase) Books() (books []data.Book, err error) {

	books, err = bookUC.bookRepo.Books()

	if err != nil {
		bookUC.logger.Printf("method Books | %s", err)
		return
	}

	return
}

func (bookUC *bookUseCase) PagedBooks(pageSize int, pageNumber int) (books []data.Book, err error) {
	books, err = bookUC.bookRepo.PagedBooks(pageSize, pageNumber)

	if err != nil {
		bookUC.logger.Printf("method PagedBooks | %s", err)
		return
	}

	return
}

func (bookUC *bookUseCase) CreateBook(book *data.Book) (err error) {

	book.ID = util.CreateUUID()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	book.IsDeleted = false

	err = bookUC.bookRepo.CreateBook(*book)

	if err != nil {
		bookUC.logger.Printf("method CreateBook | %s", err)
		return
	}

	return
}

func (bookUC *bookUseCase) DeleteBook(id string) (err error) {

	err = bookUC.bookRepo.DeleteBook(id)

	if err != nil {
		bookUC.logger.Printf("method DeleteBook | %s", err)
	}
	return
}

func (bookUC *bookUseCase) UpdateBook(book *data.Book) (err error) {

	book.UpdatedAt = time.Now()

	err = bookUC.bookRepo.UpdateBook(*book)

	if err != nil {
		bookUC.logger.Printf("method UpdateBook | %s", err)
		return
	}

	return
}
