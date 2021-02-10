package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/handlers"
)

//PublisherRouting setup user routes
func PublisherRouting(sm *http.ServeMux, publisherHandler *handlers.PublisherHandler) {

	// GET /publishers/id?id=
	sm.HandleFunc("/publishers/id", publisherHandler.GetPublisher)

	// GET publishers/id/books/?publisherid={1}
	sm.HandleFunc("/publishers/id/books", publisherHandler.GetPublisherBooks)

}
