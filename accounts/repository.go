package accounts

type Repository interface{
	FindUser(username string) (string, error)
}