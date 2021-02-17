package routing

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/diservices"
	"github.com/DenisKnez/simpleWebGolang/diutils"
	"github.com/DenisKnez/simpleWebGolang/handlers"
	"github.com/DenisKnez/simpleWebGolang/util/middleware"
)

//UserRouting setup user routes
func UserRouting(sm *http.ServeMux, userHandler *handlers.UserHandler) {

	config := diutils.GetConfig()

	// GET /users
	sm.HandleFunc("/users", middleware.Adapt(userHandler.GetUsers,
		middleware.AuthMiddleware(diservices.AuthUseCase, config)))
	// GET /users/id?id=
	sm.HandleFunc("/users/id", middleware.Adapt(userHandler.GetUser,
		middleware.AuthMiddleware(diservices.AuthUseCase, config)))
	// POST /users/createuser
	sm.HandleFunc("/users/createuser", middleware.Adapt(userHandler.CreateUser,
		middleware.AuthMiddleware(diservices.AuthUseCase, config)))
	// GET /users/paged?pagesize={1}&pagenumber={2}
	sm.HandleFunc("/users/paged", middleware.Adapt(userHandler.PagedUsers,
		middleware.AuthMiddleware(diservices.AuthUseCase, config)))
	// DELETE /users/delete/id?id={1}
	sm.HandleFunc("/users/delete/id", middleware.Adapt(userHandler.DeleteUser,
		middleware.AuthMiddleware(diservices.AuthUseCase, config)))
	// PUT /users/update/id?id={1}
	sm.HandleFunc("/users/update/id", middleware.Adapt(userHandler.UpdateUser,
		middleware.AuthMiddleware(diservices.AuthUseCase, config)))
}
