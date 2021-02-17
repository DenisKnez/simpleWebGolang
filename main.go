package main

import (
	"net/http"

	"github.com/DenisKnez/simpleWebGolang/diservices"
	"github.com/DenisKnez/simpleWebGolang/diutils"
	"github.com/DenisKnez/simpleWebGolang/handlers/routing"
	_ "github.com/golang-migrate/migrate"
	_ "github.com/spf13/viper"
)

func main() {

	file, logger := diutils.GetLogger()

	logger.Println("Program started")

	sm := http.NewServeMux()

	// Routing
	routing.AuthRouting(sm, diservices.GetAuthHandler())
	routing.PublisherRouting(sm, diservices.GetPublisherHandler())
	routing.UserRouting(sm, diservices.GetUserHandler())
	routing.BookRouting(sm, diservices.GetBookHandler())

	http.ListenAndServe(":9090", sm)
	//Close the log file
	file.Close()
}
