package rooms

type UseCase interface{
	NewRoom()
	GetRoom()
	GetAllRoomsList()
	DeleteRoom()
	AddParticipants()
	DeleteParticipants()
}