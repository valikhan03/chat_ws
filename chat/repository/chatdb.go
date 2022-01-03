package repository

import (
	"chatapp/models"
	"context"
	"fmt"
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

func (rep *ChatRepository) SaveMessage(msg *models.Message) error {
	messageStorage := rep.DB.Collection(messages)
	message := bson.D{{"sender", msg.Sender}, {"receiver", msg.Receiver}, {"payload", msg.Payload}}
	res, err := messageStorage.InsertOne(context.TODO(), message)
	if err != nil{
		log.Println(err)
		return err
	}
	msg_id := res.InsertedID.(string)
	fmt.Println(msg_id)
	return nil
}



