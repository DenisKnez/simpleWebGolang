package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/handlers"
)

//AuthRouting setup user routes
func AuthRouting(sm *http.ServeMux, bookHandler *handlers.BookHandler) {

	// GET auth/signup
	sm.HandleFunc("/auth/signup", bookHandler.Books)

	// GET auth/login
	sm.HandleFunc("/auth/login", bookHandler.Books)

	// GET auth/refresh
	sm.HandleFunc("auth/refresh", bookHandler.Books)
}
