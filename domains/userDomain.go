package domains

import (
	"github.com/DenisKnez/simpleWebGolang/data"
)

//UserUseCase user useCase interface
type UserUseCase interface {
	GetUserByID(id string) (user data.User, err error)
	Users() (users []data.User, err error)
	PagedUsers(pageSize int, pageNumber int) (users []data.User, err error)
	CreateUser(user *data.User) (err error)
	DeleteUser(id string) (err error)
	UpdateUser(user *data.User) (err error)
}

//UserRepository user repository interface
type UserRepository interface {
	GetUserByID(id string) (user data.User, err error)
	Users() (users []data.User, err error)
	PagedUsers(pageSize int, pageNumber int) (users []data.User, err error)
	CreateUser(user data.User) (err error)
	DeleteUser(id string) (err error)
	UpdateUser(user data.User) (err error)
}
