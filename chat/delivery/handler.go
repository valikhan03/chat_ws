package delivery

import (
	"fmt"
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

func NewHandler(uc chat.UseCase) *Handler {
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
	var msgCh = make(chan models.Message)
	var errCh = make(chan error)

	var msg models.Message
	go func() {
		for{
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println(err)
				errCh <- err
			}
			msgCh <- msg
			fmt.Println(msg)
			//save to db
		}		
	}()
		

	go func(){
		for{
			msg := <- msgCh
			err := c.WriteJSON(msg)
			if err != nil{
				log.Println(err)
				errCh <- err
			}
		}
	}()
		
	err := <- errCh
	if err != nil{
		log.Println(err)
	}
}

func (h *Handler) SaveMessage(msg *models.Message) {
	err := h.UseCase.SaveMessage(msg)
	if err != nil {
		log.Println(err)
	}
}


