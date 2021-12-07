package auth

import(
	"chatapp/models"
)

type UserRepository interface{
	CreateUser(user *models.User) error
	GetUser(email, password string) (models.User, error)
}