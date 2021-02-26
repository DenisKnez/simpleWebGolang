package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/diServices"
	"github.com/DenisKnez/simpleWebGolang/diUtils"
	"github.com/DenisKnez/simpleWebGolang/handlers"
	"github.com/DenisKnez/simpleWebGolang/util/middleware"
)

//UserRouting setup user routes
func UserRouting(sm *http.ServeMux, userHandler *handlers.UserHandler) {

	config := diUtils.GetConfig()

	// GET /users
	sm.HandleFunc("/users", middleware.Adapt(userHandler.GetUsers,
		middleware.AuthMiddleware(diServices.GetAuthUseCase(), config)))
	// GET /users/id?id=
	sm.HandleFunc("/users/id", middleware.Adapt(userHandler.GetUser,
		middleware.AuthMiddleware(diServices.GetAuthUseCase(), config)))
	// POST /users/createuser
	sm.HandleFunc("/users/createuser", middleware.Adapt(userHandler.CreateUser,
		middleware.AuthMiddleware(diServices.GetAuthUseCase(), config)))
	// GET /users/paged?pagesize={1}&pagenumber={2}
	sm.HandleFunc("/users/paged", middleware.Adapt(userHandler.PagedUsers,
		middleware.AuthMiddleware(diServices.GetAuthUseCase(), config)))
	// DELETE /users/delete/id?id={1}
	sm.HandleFunc("/users/delete/id", middleware.Adapt(userHandler.DeleteUser,
		middleware.AuthMiddleware(diServices.GetAuthUseCase(), config)))
	// PUT /users/update/id?id={1}
	sm.HandleFunc("/users/update/id", middleware.Adapt(userHandler.UpdateUser,
		middleware.AuthMiddleware(diServices.GetAuthUseCase(), config)))
}
