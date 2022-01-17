package delivery

import (
	"fmt"
	"log"
	"net/http"

	"chatapp/auth"
	"chatapp/chat"
	"chatapp/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	UseCase     chat.UseCase
	AuthUseCase auth.UseCase
}

func NewHandler(uc chat.UseCase, auth_uc auth.UseCase) *Handler {
	return &Handler{
		UseCase:     uc,
		AuthUseCase: auth_uc,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan models.Message)

//var msg_saver = make(chan models.Message)

func (h *Handler) HandleConnections(c *gin.Context) {
	cookie, err := c.Request.Cookie("access-token-chat-eltaev")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, http.ErrNoCookie)
		return
	}
	token := cookie.Value

	user_id, err := h.AuthUseCase.ParseToken(token)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusForbidden, err)
	}

	clients[ws] = true

	for {
		var msg models.Message
		msg.Sender = user_id
		fmt.Println(c.Param("chat_id"))
		err = ws.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
		}
		broadcaster <- msg
	}

}

func (h *Handler) HandleMessages() {
	for {
		msg := <-broadcaster
		for ws_client := range clients {
			err := ws_client.WriteJSON(&msg)
			if err != nil {
				log.Println(err)
			}
			h.SaveMessage(msg)
		}
	}
}

func (h *Handler) SaveMessage(msg models.Message) {
	err := h.UseCase.SaveMessage(msg)
	if err != nil {
		log.Println(err)
	}
}
