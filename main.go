package main

import (
	"net/http"

	services "github.com/DenisKnez/simpleWebGolang/diServices"
	utils "github.com/DenisKnez/simpleWebGolang/diUtils"
	"github.com/DenisKnez/simpleWebGolang/handlers/routing"
	_ "github.com/golang-migrate/migrate"
	_ "github.com/spf13/viper"
)

func main() {

	file, logger := utils.GetLogger()

	logger.Println("Program started")

	sm := http.NewServeMux()

	// Routing
	routing.AuthRouting(sm, services.GetAuthHandler())
	routing.PublisherRouting(sm, services.GetPublisherHandler())
	routing.UserRouting(sm, services.GetUserHandler())
	routing.BookRouting(sm, services.GetBookHandler())

	http.ListenAndServe(":9090", sm)
	//Close the log file
	file.Close()
}
