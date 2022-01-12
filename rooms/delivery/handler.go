package delivery

import (
	"chatapp/models"
	"chatapp/rooms"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase rooms.UseCase
}

func NewHandler(uc rooms.UseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var room models.Room
	c.BindJSON(&room)
	room_id, err := h.usecase.NewRoom(room)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Writer.Write([]byte(room_id))
}

func (h *Handler) GetRoom(c *gin.Context) {
	var room_id string
	c.Bind(&room_id)
	room := h.usecase.GetRoom(room_id)

	c.JSON(200, room)
}

func (h *Handler) GetAllRoomsList(c *gin.Context) {
	user_id := c.GetString("user_id")

	fmt.Println(user_id)

	rooms, err := h.usecase.GetAllRoomsList(user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, rooms)
}
