package repositories

import (
	"database/sql"
	"log"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
)

type authRepository struct {
	logger *log.Logger
	conn   *sql.DB
}

//NewAuthRepository create a new auth repository
func NewAuthRepository(conn *sql.DB, logger *log.Logger) domains.AuthRepository {
	return &authRepository{logger, conn}
}


//GetUserPassByEmail gets the user password with the provided email
func(repo *authRepository) GetUserPassByEmail(email string) (password string, err error){
	err = repo.conn.QueryRow("SELECT password FROM users WHERE email = $1", email).
						Scan(&password)
	if err != nil {
		repo.logger.Printf("method GetUserPassByEmail | %s", err)
		return
		
	}

	return 
}

