package accounts

import(
)

type UseCase interface{
	FindUser(username string) (string, error)
}


