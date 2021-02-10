package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/handlers"
)

//PublisherRouting setup user routes
func PublisherRouting(sm *http.ServeMux, publisherHandler *handlers.PublisherHandler) {

	// GET /books
	//sm.HandleFunc("/books", bookHandler.Books)
	// GET /publishers/id?id=
	sm.HandleFunc("/publishers/id", publisherHandler.GetPublisher)
	// // POST /books/createbook
	// sm.HandleFunc("/books/createbook", bookHandler.CreateBook)
	// // GET /books/paged?pagesize={1}&pagenumber={2}
	// sm.HandleFunc("/books/paged", bookHandler.PagedBooks)
	// // GET /books/delete/id?id={1}
	// sm.HandleFunc("/books/delete/id", bookHandler.DeleteBook)
	// // PUT /books/update/id?id={1}
	// sm.HandleFunc("/books/update/id", bookHandler.UpdateBook)

}
