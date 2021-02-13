package usecase

import (
	"log"
	"github.com/DenisKnez/simpleWebGolang/domains"
)

type authUseCase struct {
	authRepo domains.authRepository
	logger *log.Logger
}

//NewAuthUseCase creates a new auth usecase
func NewAuthUseCase(authRepo domains.authRepository, logger *log.Logger) domains.AuthUseCase {
	return &authUseCase{authRepo, logger}
}

//GetUserPassByEmail get the user password with the provided email
func(usecase *authUseCase) GetUserPassByEmail(email string) (password string, err error) {
	pass, err = usecase.authRepo.GetUserPassByEmail(email)

	if err != nil {
		usecase.logger.Printf("method GetUserPassByEmail | %s", err)
		return
	}

	return
}