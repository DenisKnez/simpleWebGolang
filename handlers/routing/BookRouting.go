package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/handlers"
)

//BookRouting setup user routes
func BookRouting(sm *http.ServeMux, bookHandler *handlers.BookHandler) {

	// GET /books
	sm.HandleFunc("/books", bookHandler.Books)
	// GET /books/id?id=
	sm.HandleFunc("/books/id", bookHandler.GetBook)
	// POST /books/createbook
	sm.HandleFunc("/books/createbook", bookHandler.CreateBook)
	// GET /books/paged?pagesize={1}&pagenumber={2}
	sm.HandleFunc("/books/paged", bookHandler.PagedBooks)
	// GET /books/delete/id?id={1}
	sm.HandleFunc("/books/delete/id", bookHandler.DeleteBook)
	// PUT /books/update/id?id={1}
	sm.HandleFunc("/books/update/id", bookHandler.UpdateBook)

}
