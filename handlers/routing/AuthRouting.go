package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/handlers"
)

//AuthRouting setup user routes
func AuthRouting(sm *http.ServeMux, authHandler *handlers.AuthHandler) {

	// POST auth/refresh
	sm.HandleFunc("/auth/refresh", authHandler.Refresh)

	// GET auth/test/token
	sm.HandleFunc("/auth/test/token", authHandler.TestToken)
}
