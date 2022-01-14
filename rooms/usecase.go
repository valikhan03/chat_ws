package rooms

import (
	"chatapp/models"
)

type UseCase interface {
	NewRoom(title string, owner string, participants []string) (string, error)
	GetRoom(id string) models.Room
	GetAllRoomsList(user_id string) ([]models.Room, error)
	DeleteRoom(id string) bool
	AddParticipants(room_id string, users_id []string) bool
	DeleteParticipants()
}
