package repository

import (
	"chatapp/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct{
	DB *mongo.Database
}

func NewChatRepository(db *mongo.Database) *ChatRepository{
	return &ChatRepository{
		DB: db,
	}
}

const(
	messages = "messages"
)

func (rep *ChatRepository) SaveMessage(msg models.Message) error {
	messageStorage := rep.DB.Collection(messages)
	message := bson.D{{"sender", msg.Sender}, {"chat", msg.ChatID}, {"payload", msg.Payload}}
	_, err := messageStorage.InsertOne(context.TODO(), message)
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}



