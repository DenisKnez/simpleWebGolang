package usecase

import (
	"log"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
)

//PublisherUseCase publisher use case
type PublisherUseCase struct {
	publisherRepo domains.PublisherRepository
	logger        *log.Logger
}

//NewPublisherUseCase create new publisher use case
func NewPublisherUseCase(publisherRepo domains.PublisherRepository, logger *log.Logger) domains.PublisherUseCase {
	return &PublisherUseCase{publisherRepo, logger}
}

//GetPublisherByID get publisher by provided id
func (publisherUC *PublisherUseCase) GetPublisherByID(id string) (publisher data.Publisher, err error) {
	publisher, err = publisherUC.publisherRepo.GetPublisherByID(id)

	if err != nil {
		publisherUC.logger.Printf("method GetPublisherByID | %s", err)
	}

	return
}

//GetPublisherBooks gets the books for the provided publisher
func (publisherUC *PublisherUseCase) GetPublisherBooks(publisherID string) (books []data.Book, err error) {
	books, err = publisherUC.publisherRepo.GetPublisherBooks(publisherID)

	if err != nil {
		publisherUC.logger.Printf("method GetPublisherBooks | %s", err)
	}

	return
}
