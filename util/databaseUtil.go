package util

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/DenisKnez/simpleWebGolang/diutils"
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
	config := diutils.GetConfig()
	_, logger := diutils.GetLogger()
	connString := config.GetString("Databases.PostgresConnection")

	var err error

	Db, err = sql.Open("pgx", connString)

	if err != nil {
		logger.Printf("method init | %s", err)
		panic("Database connection failed")
	}

	return
}
