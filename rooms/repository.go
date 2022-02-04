package rooms

import (
	"chatapp/models"
)

type Repository interface {
	NewRoom(title string, owner string, participants []string, room_type string) (string, error)
	GetRoom(room_id string) models.Room
	GetAllRoomsList(user_id string) ([]models.Room, error)
	DeleteRoom(room_id string) bool
	AddParticipants(room_id string, users_id []string) (bool, error)
	DeleteParticipants()
	GetUsernameByID(id string) (string, error)
}