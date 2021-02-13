package domains


type AuthRepository interface {
	GetUserPassByEmail(email string) (password string, err error)
}

type AuthUseCase interface {
	GetUserPassByEmail(email string) (password string, err error)
}