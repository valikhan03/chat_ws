package accountsUsecase

import(
	"chatapp/accounts"

)

type AccountsUseCase struct{
	repository accounts.Repository 
}

func NewAccountsUseCase(repos accounts.Repository) *AccountsUseCase{
	return &AccountsUseCase{
		repository: repos,
	}
}

func (u *AccountsUseCase) FindUser(username string) (string, error){
	return u.repository.FindUser(username)
}
