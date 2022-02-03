package rooms

import (
	"chatapp/models"
)

type UseCase interface {
	NewGroupRoom(title string, owner string, participants []string) (string, error)
	NewCommonRoom(contact1 string, contact2 string) (string, error)
	GetRoom(id string) models.Room
	GetAllRoomsList(user_id string) ([]models.Room, error)
	DeleteRoom(id string) bool
	AddParticipants(room_id string, users_id []string) bool
	DeleteParticipants()
}
