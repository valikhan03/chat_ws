package auth

type UseCase interface {
	SignUp(username, email, password string) error
	ParseToken(accessToken string) (string, error)
	GenerateAuthToken(email string, password string) (string, error)
}
