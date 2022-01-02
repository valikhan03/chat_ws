package delivery

import(
	"chatapp/rooms"
)

type Handler struct{
	usecase rooms.UseCase
}

func NewHandler(uc rooms.UseCase) *Handler{
	return &Handler{
		usecase: uc,
	}
}


