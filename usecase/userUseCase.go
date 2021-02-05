package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/domains"
	"github.com/DenisKnez/simpleWebGolang/util"
)

type userUseCase struct {
	repo   domains.UserRepository
	logger *log.Logger
}

//NewUserUseCase returns new user useCase
func NewUserUseCase(repo domains.UserRepository, logger *log.Logger) domains.UserUseCase {

	return &userUseCase{repo, logger}
}

func (userUC *userUseCase) GetUserByID(id string) (user data.User, err error) {
	user, err = userUC.repo.GetUserByID(id)
	return
}

func (userUC *userUseCase) Users() (users []data.User, err error) {
	users, err = userUC.repo.Users()
	return
}

func (userUC *userUseCase) PagedUsers(pageSize int, pageNumber int) (users []data.User, err error) {
	users, err = userUC.repo.PagedUsers(pageSize, pageNumber)
	return
}

func (userUC *userUseCase) CreateUser(user *data.User) (err error) {

	user.ID = util.CreateUUID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsDeleted = false

	err = userUC.repo.CreateUser(*user)
	return
}


func (userUC *userUseCase) DeleteUser(id string) (err error){
	err = userUC.repo.DeleteUser(id)

	if err != nil {
		fmt.Println(err)
	}

	return
}