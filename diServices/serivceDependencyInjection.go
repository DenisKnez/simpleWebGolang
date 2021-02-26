package diServices

import (
	"github.com/DenisKnez/simpleWebGolang/diUtils"
	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/DenisKnez/simpleWebGolang/handlers"
	"github.com/DenisKnez/simpleWebGolang/repositories"
	"github.com/DenisKnez/simpleWebGolang/usecase"
	"github.com/DenisKnez/simpleWebGolang/util"
)

//usecase
var authUseCase domains.AuthUseCase
var userUseCase domains.UserUseCase
var bookUseCase domains.BookUseCase
var publisherUseCase domains.PublisherUseCase

//handlers
var authHandler *handlers.AuthHandler
var userHandler *handlers.UserHandler
var bookHandler *handlers.BookHandler
var publisherHandler *handlers.PublisherHandler

func init() {

	_, logger := diUtils.GetLogger()
	config := diUtils.GetConfig()

	// user DI
	userRepo := repositories.NewUserRepository(util.Db, logger)
	userUseCase := usecase.NewUserUseCase(userRepo, logger)
	userHandler = handlers.NewUserHandler(userUseCase)

	// book DI
	bookRepo := repositories.NewBookRepository(util.Db, logger)
	bookUseCase := usecase.NewBookUseCase(bookRepo, logger)
	bookHandler = handlers.NewBookHandler(bookUseCase)

	// publisher DI
	publisherRepo := repositories.NewPublisherRepository(util.Db, logger)
	publisherUseCase := usecase.NewPublisherUseCase(publisherRepo, logger)
	publisherHandler = handlers.NewPublisherHandler(publisherUseCase)

	// auth DI
	authRepo := repositories.NewAuthRepository(util.Db, logger)
	authUseCase := usecase.NewAuthUseCase(authRepo, logger)
	authHandler = handlers.NewAuthHandler(authUseCase, config)
}

// HANDLERS

//GetUserHandler Returns the user handler
func GetUserHandler() *handlers.UserHandler {
	return userHandler
}

//GetBookHandler Returns the book handler
func GetBookHandler() *handlers.BookHandler {
	return bookHandler
}

//GetPublisherHandler Returns the publisher handler
func GetPublisherHandler() *handlers.PublisherHandler {
	return publisherHandler
}

//GetAuthHandler Returns the auth handler
func GetAuthHandler() *handlers.AuthHandler {
	return authHandler
}

// USECASES

//GetAuthUseCase Returns the auth UseCase
func GetAuthUseCase() domains.AuthUseCase {
	return authUseCase
}

//GetUserUseCase Returns the user UseCase
func GetUserUseCase() domains.UserUseCase {
	return userUseCase
}

//GetBookUseCase Returns the book UseCase
func GetBookUseCase() domains.BookUseCase {
	return bookUseCase
}

//GetPublisherUseCase Returns the publisher UseCase
func GetPublisherUseCase() domains.PublisherUseCase {
	return publisherUseCase
}
