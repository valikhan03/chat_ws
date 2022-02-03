package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type AccountsRepository struct {
	DB *sqlx.DB
}

func NewAccountsRepository(db *sqlx.DB) *AccountsRepository {
	return &AccountsRepository{
		DB: db,
	}
}

type userID struct{
	Id string `json:"id"`
}

func (rep *AccountsRepository) FindUser(username string) (string, error) {
	var user_id userID
	err := rep.DB.Get(&user_id, "select id from chat_users where username=$1 LIMIT 1", username)
	if err != nil{
		log.Println(err)
		return "", err
	}

	return user_id.Id, nil
}
