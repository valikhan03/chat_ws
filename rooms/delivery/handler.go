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


type NewCommonChatInput struct {
	ContactUser string `json:"contact_user"`
}

func (h *Handler) CreateCommonRoom(c *gin.Context) {
	cookie, err := c.Request.Cookie("access-token-chat-eltaev")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, http.ErrNoCookie)
		return
	}
	token := cookie.Value

	fmt.Println(token)

	user_id, err := h.auth_usecase.ParseToken(token)

	var contact NewCommonChatInput
	c.BindJSON(&contact)

	room_id, err := h.usecase.NewCommonRoom(user_id, contact.ContactUser)

	c.Writer.Write([]byte(room_id))
}

type NewRoomInput struct {
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}

func (h *Handler) CreateGroupRoom(c *gin.Context) {
	var room NewRoomInput
	cookie, err := c.Request.Cookie("access-token-chat-eltaev")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, http.ErrNoCookie)
		return
	}
	token := cookie.Value

	user_id, err := h.auth_usecase.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.BindJSON(&room)
	room.Participants = append(room.Participants, user_id)

	room_id, err := h.usecase.NewGroupRoom(room.Title, user_id, room.Participants)
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
