package diservices

import (
	"github.com/DenisKnez/simpleWebGolang/diutils"
	"github.com/DenisKnez/simpleWebGolang/handlers"
	"github.com/DenisKnez/simpleWebGolang/repositories"
	"github.com/DenisKnez/simpleWebGolang/usecase"
	"github.com/DenisKnez/simpleWebGolang/util"
)

var userHandler *handlers.UserHandler
var bookHandler *handlers.BookHandler

func init() {

	_, logger := diutils.GetLogger()
	//config := diutils.GetConfig()

	// user DI
	userRepo := repositories.NewUserRepository(util.Db, logger)
	userUseCase := usecase.NewUserUseCase(userRepo, logger)
	userHandler = handlers.NewUserHandler(userUseCase)

	// book DI
	bookRepo := repositories.NewBookRepository(util.Db, logger)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	bookHandler = handlers.NewBookHandler(bookUseCase)
}

//GetUserHandler Returns the user handler
func GetUserHandler() *handlers.UserHandler {
	return userHandler
}

//GetBookHandler Returns the book handler
func GetBookHandler() *handlers.BookHandler {
	return bookHandler
}
