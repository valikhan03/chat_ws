package usecase

import (
	"chatapp/models"
	"chatapp/rooms"

)

type UseCase struct{
	repository rooms.Repository
}

func NewRoomsUseCase(rep rooms.Repository) *UseCase{
	return &UseCase{
		repository: rep,
	}
}


func (u *UseCase) NewRoom(title string, owner string, participants []string) (string, error){
	room_id, err := u.repository.NewRoom(title, owner, participants)
	if err != nil{
		return "", err
	}
	return room_id, nil
}

func (u *UseCase) GetRoom(id string) (models.Room){
	room := u.repository.GetRoom(id)
	return room
}

func (u *UseCase) GetAllRoomsList(user_id string) ([]models.Room, error){
	rooms, err :=u.repository.GetAllRoomsList(user_id)
	if err != nil{
		return nil, err
	}
	return rooms, nil
}

func (u *UseCase) DeleteRoom(id string) bool{
	res := u.repository.DeleteRoom(id)
	return res
}

func (u *UseCase) AddParticipants(room_id string, users_id []string) bool {
	res, _ := u.repository.AddParticipants(room_id, users_id)
	
	return res
}

func (u *UseCase) DeleteParticipants(){

}
