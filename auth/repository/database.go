package repository

import (
	"chatapp/models"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (rep *UserRepository) CreateUser(user *models.User) error {
	_, err := rep.database.Exec("insert into chat_users (id, email, username, password) values ($1, $2, $3, $4)", user.Id, user.Email, user.Username, user.Password)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (rep *UserRepository) GetUser(email, password string) (models.User, error) {
	user := models.User{}
	err := rep.database.Get(&user, "select id from chat_users where email=$1 and password=$2 LIMIT 1", email, password)
	if err != nil {
		log.Println(err)
	}

	return user, err
}
