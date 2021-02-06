package repositories

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/DenisKnez/simpleWebGolang/util"
)

//userRepository user repository
type userRepository struct {
	conn   *sql.DB
	logger *log.Logger
}

//NewUserRepository create an new instance of the user repository
func NewUserRepository(conn *sql.DB, logger *log.Logger) domains.UserRepository {
	return &userRepository{conn, logger}
}

//GetUserByID get the user by its id
func (userRepo *userRepository) GetUserByID(id string) (user data.User, err error) {
	user = data.User{}
	err = util.Db.QueryRow("SELECT id, name, lastname, age, email, password, created_at, updated_at, deleted_at, is_deleted FROM users WHERE id = $1 AND is_deleted = false", id).
		Scan(&user.ID, &user.Name, &user.Lastname, &user.Age, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.IsDeleted)

	fmt.Println(err)
	return
}

//Users returns all users
func (userRepo *userRepository) Users() (users []data.User, err error) {
	rows, err := util.Db.Query("SELECT id, name, lastname, age, email, password, created_at, updated_at, deleted_at, is_deleted FROM users WHERE is_deleted = false")

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		user := data.User{}
		if err = rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Age, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.IsDeleted); err != nil {
			fmt.Println(err)
			return
		}

		users = append(users, user)
	}
	rows.Close()

	return
}

//PagedUsers returns all users in a limited range and by a certain offset
func (userRepo *userRepository) PagedUsers(pageSize int, pageNumber int) (users []data.User, err error) {

	// calculates the offset used in the sql based on the pageSize and pageNumber
	offset := (pageSize*pageNumber - pageSize)

	rows, err := util.Db.Query(
		`SELECT id, name, lastname, age, email, password, 
				created_at, updated_at, deleted_at, is_deleted
		FROM users
		ORDER BY id
		LIMIT $1
		OFFSET $2`, pageSize, offset)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		user := data.User{}
		fmt.Println(user)
		if err = rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Age, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.IsDeleted); err != nil {
			fmt.Println(err)
			return
		}

		users = append(users, user)
	}
	fmt.Print(users)
	rows.Close()
	return
}

//CreateUser create a new user
func (userRepo *userRepository) CreateUser(user data.User) (err error) {
	stmt, err := util.Db.Prepare(`INSERT INTO users 
					(id, name, lastname, age, email, password, created_at, updated_at, deleted_at, is_deleted) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
					RETURNING id, name, lastname, age, email, password, created_at, updated_at, deleted_at, is_deleted
					`)
	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	result, stmtErr := stmt.Exec(user.ID, user.Name, user.Lastname, user.Age, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.IsDeleted)

	if stmtErr != nil {
		fmt.Println(stmtErr)
		return
	}

	userRepo.logger.Println(result)
	return nil
}

//DeleteUser deletes the user with the provided id
func (userRepo *userRepository) DeleteUser(id string) (err error) {

	stmt, err := util.Db.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

//UpdateUser deletes the user with the provided id
func (userRepo *userRepository) UpdateUser(user data.User) (err error) {

	stmt, err := util.Db.Prepare(`UPDATE users SET 
		name = $2,
		lastname = $3,
		age = $4,
		email = $5,
		password = $6,
		updated_at = $7
		WHERE id = $1`)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.Lastname, user.Age, user.Email, user.Password, user.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
