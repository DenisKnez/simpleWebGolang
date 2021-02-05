package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DenisKnez/simpleWebGolang/handlers"
	"github.com/DenisKnez/simpleWebGolang/repositories"
	"github.com/DenisKnez/simpleWebGolang/usecase"
	"github.com/DenisKnez/simpleWebGolang/util"
)

func main() {

	sm := http.NewServeMux()

	logger := log.New(os.Stdout, "LOG: ", log.Flags())

	// user DI
	userRepo := repositories.NewUserRepository(util.Db, logger)
	userUseCase := usecase.NewUserUseCase(userRepo, logger)
	userHandler := handlers.NewUserHandler(userUseCase)

	// book DI
	bookRepo := repositories.NewBookRepository(util.Db, logger)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	bookHandler := handlers.NewBookHandler(bookUseCase)

	//Users

	// GET /users
	sm.HandleFunc("/users", userHandler.GetUsers)
	// GET /users/id?id=
	sm.HandleFunc("/users/id", userHandler.GetUser)
	// POST /users/createuser
	sm.HandleFunc("/users/createuser", userHandler.CreateUser)
	// GET /users/paged?pagesize={1}&pagenumber={2}
	sm.HandleFunc("/users/paged", userHandler.PagedUsers)
	// GET /users/delete/id?id={1}
	sm.HandleFunc("/users/delete/id", userHandler.DeleteUser)

	//Books

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

	http.ListenAndServe(":9090", sm)

}
