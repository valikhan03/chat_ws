package repository

import (
	"chatapp/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomsRepository struct {
	DB *mongo.Database
}

func NewRoomRepository(db *mongo.Database) *RoomsRepository {
	return &RoomsRepository{
		DB: db,
	}
}

const (
	roomsCollection = "rooms"
)

func (r *RoomsRepository) NewRoom(title string, owner string, participants []string) (string, error) {
	collection := r.DB.Collection(roomsCollection)

	res, err := collection.InsertOne(context.Background(), bson.M{"title": title, "owner": owner, "participants": bson.A{participants}})
	if err != nil {
		log.Println(err)
	}
	return res.InsertedID.(string), err
}

func (r *RoomsRepository) GetRoom(room_id string) models.Room {
	collection := r.DB.Collection(roomsCollection)
	var room models.Room
	res := collection.FindOne(context.Background(), bson.D{{"id", room_id}})
	res.Decode(&room)

	return room
}

func (r *RoomsRepository) GetAllRoomsList(user_id string) ([]models.Room, error) {
	collection := r.DB.Collection(roomsCollection)
	cur, err := collection.Find(context.Background(), bson.M{"participants": bson.M{"$in": user_id}})
	if err != nil {
		log.Println(err)
	}

	var rooms []models.Room
	err = cur.All(context.Background(), &rooms)
	if err != nil {
		log.Println(err)
	}

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
	collection := r.DB.Collection(roomsCollection)
	filer := options.ArrayFilters{Filters: bson.A{bson.M{"id": room_id}}}
	update := bson.M{
		"$push": bson.M{"participants": users_id},
	}

	_, err := collection.UpdateOne(context.Background(), filer, update)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (r *RoomsRepository) DeleteParticipants() {

}
