package domains

import (
	"github.com/DenisKnez/simpleWebGolang/data"
)

//PublisherRepository publisher repository interface
type PublisherRepository interface {
	GetPublisherByID(id string) (publisher data.Publisher, err error)
	GetPublisherBooks(publisherID string) (books []data.Book, err error)
}

//PublisherUseCase publisher useCase interface
type PublisherUseCase interface {
	GetPublisherByID(id string) (publisher data.Publisher, err error)
	GetPublisherBooks(publisherID string) (books []data.Book, err error)
}
