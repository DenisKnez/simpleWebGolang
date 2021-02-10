package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/domains"
)

//PublisherHandler publisher handler
type PublisherHandler struct {
	publisherUC domains.PublisherUseCase
}

//NewPublisherHandler create new publisher handler
func NewPublisherHandler(publisherUC domains.PublisherUseCase) *PublisherHandler {
	return &PublisherHandler{publisherUC}
}

//GetPublisher get a publisher with the provided parameter id
func (handler *PublisherHandler) GetPublisher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Required parameter id was not provided", http.StatusBadRequest)
		return
	}

	publisher, err := handler.publisherUC.GetPublisherByID(id)

	if err != nil {
		http.Error(w, "Publisher with the provided id was not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(publisher)

	if err != nil {
		http.Error(w, "Couldn't encode the json", http.StatusInternalServerError)
	}

	return
}

//GetPublisherBooks gets the books by provided publisher
func (handler *PublisherHandler) GetPublisherBooks(w http.ResponseWriter, r *http.Request) {
	publisher := r.URL.Query().Get("publisherid")

	if publisher == "" {
		http.Error(w, "Required parameter publisher was not provided", http.StatusBadRequest)
		return
	}

	books, err := handler.publisherUC.GetPublisherBooks(publisher)

	if err != nil {
		http.Error(w, "Publisher with the provided id was not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		http.Error(w, "Couldn't encode the json", http.StatusInternalServerError)
	}

	return
}
