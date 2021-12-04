package auth

type UseCase interface {
	SignUp(username, email, password string) error
	SignIn(email, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	GenerateAuthToken(user_id string) (string, error)
}
