package util

import (
	"database/sql"

	"github.com/google/uuid"

	utils "github.com/DenisKnez/simpleWebGolang/diUtils"
	//used to provide a driver for the postgresql database
	_ "github.com/jackc/pgx/v4/stdlib"
)

//CreateUUID create a new uuid
func CreateUUID() uuid.UUID {
	return uuid.Must(uuid.NewRandom())
}

//Db database connection
var Db *sql.DB

func init() {
	config := utils.GetConfig()
	_, logger := utils.GetLogger()
	connString := config.GetString("Databases.PostgresConnection")

	var err error

	Db, err = sql.Open("pgx", connString)

	if err != nil {
		logger.Printf("method init | %s", err)
		panic("Database connection failed")
	}

	return
}
