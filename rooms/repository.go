package rooms

import(
	"chatapp/models"
)

type Repository interface{
	NewRoom(room_id string, title string, participants []string) (string, error)
	GetRoom(room_id string) models.Room
	GetAllRoomsList(user_id string) ([]models.Room, error) 
	DeleteRoom(room_id string) bool
	AddParticipants()
	DeleteParticipants()
}