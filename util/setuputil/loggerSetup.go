package setuputil

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//SetupLogger creates the file and the directory for logging and returns
// the logger
func SetupLogger(config *viper.Viper) (file *os.File, logger *log.Logger) {

	dirLocation := config.GetString("Logging.DirLocation")
	fileLocation := config.GetString("Logging.FileLocation")

	err := os.Mkdir(dirLocation, os.ModeDir)

	if os.IsNotExist(err) && err != nil {
		fmt.Println(err)
	}

	file, err = os.OpenFile(fileLocation, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err != nil && err != os.ErrExist {
		fmt.Println(err)
		return
	}

	logger = log.New(file, "LOG: ", log.Lshortfile|log.Ldate|log.Ltime)

	return
}
