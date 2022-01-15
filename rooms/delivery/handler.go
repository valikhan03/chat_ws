package delivery

import (
	"chatapp/auth"
	"chatapp/rooms"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase      rooms.UseCase
	auth_usecase auth.UseCase
}

func NewHandler(uc rooms.UseCase, auth_uc auth.UseCase) *Handler {
	return &Handler{
		usecase:      uc,
		auth_usecase: auth_uc,
	}
}

type NewRoomInput struct {
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var room NewRoomInput
	cookie, err := c.Request.Cookie("access-token-chat-eltaev")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, http.ErrNoCookie)
		return
	}
	token := cookie.Value

	fmt.Println(token)

	user_id, err := h.auth_usecase.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.BindJSON(&room)

	room.Participants = append(room.Participants, user_id)

	room_id, err := h.usecase.NewRoom(room.Title, user_id, room.Participants)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Writer.Write([]byte(room_id))
}

func (h *Handler) GetRoom(c *gin.Context) {

	room_id := c.Param("chat_id")

	fmt.Println(room_id)

	room := h.usecase.GetRoom(room_id)

	c.JSON(200, room)
}

func (h *Handler) GetAllRoomsList(c *gin.Context) {
	cookie, err := c.Request.Cookie("access-token-chat-eltaev")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, http.ErrNoCookie)
		return
	}
	token := cookie.Value

	user_id, err := h.auth_usecase.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	rooms, err := h.usecase.GetAllRoomsList(user_id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, rooms)
}

type NewParticipantsInput struct {
	Payload []string `json:"participants"`
}

func (h *Handler) AddParticipants(c *gin.Context) {
	room_id := c.Param("chat_id")
	var participants NewParticipantsInput
	c.BindJSON(&participants)

	fmt.Println(room_id+"\n", participants.Payload)

	if h.usecase.AddParticipants(room_id, participants.Payload) {
		c.Status(http.StatusOK)
		return
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
