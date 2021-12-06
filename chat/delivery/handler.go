package wsdelivery

import (
	"log"
	"net/http"

	"chatapp/chat"
	"chatapp/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	UseCase chat.UseCase
}

func NewHandler(uc chat.UseCase)*Handler{
	return &Handler{
		UseCase: uc,
	}
}

var upgrader = websocket.Upgrader{}

func (h *Handler) WSEndpoint(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusForbidden, "connection error")
	}

	h.messageReader(conn)

}

func (h *Handler) messageReader(c *websocket.Conn) {
	defer c.Close()
	var msg models.Message
	for {
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
		} else {
			h.SaveMessage(&msg)
		}

	}
}

func (h *Handler) SaveMessage(msg *models.Message) {
	err := h.UseCase.SaveMessage(msg)
	if err != nil {
		log.Println(err)
	}
}
