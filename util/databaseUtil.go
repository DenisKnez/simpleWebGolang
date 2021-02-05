package util

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

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
	var err error

	Db, err = sql.Open("pgx", "user=postgres password=rootPassword dbname=simplewebgolang sslmode=disable")

	if err != nil {
		fmt.Println("ERROR: ")
		fmt.Println(err)
	}

	return
}
