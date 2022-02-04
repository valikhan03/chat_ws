package repository

import (
	"chatapp/models"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomsRepository struct {
	DB         *mongo.Database
	PostgresDB *sqlx.DB
}

func NewRoomRepository(db *mongo.Database, sql_db *sqlx.DB) *RoomsRepository {
	return &RoomsRepository{
		DB:         db,
		PostgresDB: sql_db,
	}
}

const (
	roomsCollection = "rooms"
)

func (r *RoomsRepository) NewRoom(title string, owner string, participants []string, room_type string) (string, error) {
	collection := r.DB.Collection(roomsCollection)

	if room_type == "common" {
		username1, _ := r.GetUsernameByID(participants[0])
		username2, _ := r.GetUsernameByID(participants[1])
		title = username1 + "-" + username2
	}

	id_gen, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return "", err
	}

	id := id_gen.String()
	_, err = collection.InsertOne(context.Background(), bson.M{"id": id, "title": title, "owner": owner, "participants": participants, "type": room_type})
	if err != nil {
		log.Println("Mongo New Room Error:", err)
	}

	return id, err
}

func (r *RoomsRepository) GetRoom(room_id string) models.Room {
	collection := r.DB.Collection(roomsCollection)
	var room models.Room
	res := collection.FindOne(context.Background(), bson.M{"id": room_id})
	res.Decode(&room)

	return room
}

func (r *RoomsRepository) GetAllRoomsList(user_id string) ([]models.Room, error) {
	collection := r.DB.Collection(roomsCollection)
	var participants []string
	participants = append(participants, user_id)
	fmt.Println(participants)
	cur, err := collection.Find(context.Background(), bson.M{"participants": bson.M{"$all": participants}})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rooms []models.Room
	err = cur.All(context.Background(), &rooms)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(rooms)

	return rooms, err
}

func (r *RoomsRepository) DeleteRoom(room_id string) bool {
	collection := r.DB.Collection(roomsCollection)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": room_id})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *RoomsRepository) AddParticipants(room_id string, users_id []string) (bool, error) {
	fmt.Println(room_id+"\n", users_id)

	collection := r.DB.Collection(roomsCollection)
	update := bson.M{"$push": bson.M{"participants": bson.M{"$each": users_id}}}
	_, err := collection.UpdateMany(context.Background(), bson.M{"id": room_id}, update)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (r *RoomsRepository) DeleteParticipants() {

}

type userName struct {
	username string `db:"username"`
}

func (r *RoomsRepository) GetUsernameByID(id string) (string, error) {
	var users []string
	err := r.PostgresDB.Select(&users, "select username from chat_users where id=$1 LIMIT 1", id)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return users[0], nil
}
