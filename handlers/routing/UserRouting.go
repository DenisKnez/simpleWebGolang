package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/handlers"
	"github.com/DenisKnez/simpleWebGolang/util/middleware"
)

//UserRouting setup user routes
func UserRouting (sm *http.ServeMux, userHandler *handlers.UserHandler) {

	// GET /users
	sm.HandleFunc("/users", middleware.Adapt(userHandler.GetUsers, middleware.NotifyMiddleware()))
	// GET /users/id?id=
	sm.HandleFunc("/users/id", userHandler.GetUser)
	// POST /users/createuser
	sm.HandleFunc("/users/createuser", userHandler.CreateUser)
	// GET /users/paged?pagesize={1}&pagenumber={2}
	sm.HandleFunc("/users/paged", userHandler.PagedUsers)
	// DELETE /users/delete/id?id={1}
	sm.HandleFunc("/users/delete/id", userHandler.DeleteUser)
	// PUT /users/update/id?id={1}
	sm.HandleFunc("/users/update/id", userHandler.UpdateUser)
}

