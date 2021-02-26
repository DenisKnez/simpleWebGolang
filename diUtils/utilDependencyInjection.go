package diUtils

import (
	"log"
	"os"

	setup "github.com/DenisKnez/simpleWebGolang/util/setuputil"
	"github.com/spf13/viper"
)

var config = setup.SetupConfig()
var file, logger = setup.SetupLogger(config)

//GetLogger get logger and the log file
func GetLogger() (file *os.File, log *log.Logger) {
	return file, logger
}

//GetConfig get config file
func GetConfig() viper.Viper {
	return *config
}
