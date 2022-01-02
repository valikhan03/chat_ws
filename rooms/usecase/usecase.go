package usecase

import(
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


