package delivery

import(
	"github.com/gin-gonic/gin"

	"chatapp/rooms"
)

func RegisterRoomsAPI(router *gin.RouterGroup, usecase rooms.UseCase){
	h := NewHandler(usecase)

	router.POST("/new-chat", h.CreateRoom)
	router.GET("/my-chats", h.GetAllRoomsList)
}