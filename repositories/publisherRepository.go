package repositories

import (
	"database/sql"
	"log"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
)

type publisherRepository struct {
	logger *log.Logger
	conn   *sql.DB
}

//NewPublisherRepository create a new publisher repository
func NewPublisherRepository(conn *sql.DB, logger *log.Logger) domains.PublisherRepository {
	return &publisherRepository{logger, conn}
}

//GetPublisherByID gets the publisher with the provided id
func (publisherRepo *publisherRepository) GetPublisherByID(id string) (publisher data.Publisher, err error) {
	publisher = data.Publisher{}
	err = publisherRepo.conn.QueryRow("SELECT id, name, date_founded, created_at, updated_at, deleted_at, is_deleted FROM publishers WHERE id = $1", id).
		Scan(&publisher.ID, &publisher.Name, &publisher.DateFounded, &publisher.CreatedAt, &publisher.UpdatedAt, &publisher.DeletedAt, &publisher.IsDeleted)

	if err != nil {
		publisherRepo.logger.Printf("method GetPublisherByID | %s", err)
	}

	return
}

func (publisherRepo *publisherRepository) GetPublisherBooks(publisherID string) (books []data.Book, err error) {
	rows, err := publisherRepo.conn.Query("SELECT title, author, release_date FROM books WHERE publisher_id = $1", publisherID)

	for rows.Next() {
		book := data.Book{}

		err = rows.Scan(&book.Title, &book.Author, &book.ReleaseDate)

		if err != nil {
			publisherRepo.logger.Printf("method GetPublisherBooks | %s", err)
			return
		}

		books = append(books, book)
	}

	rows.Close()
	return
}
