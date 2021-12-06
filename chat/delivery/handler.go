package wsdelivery

import (
	"fmt"
	"html/template"
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
	defer c.Close()
	var msg models.Message
	for {
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		//h.SaveMessage(&msg)
		fmt.Println(msg)
		/*
		msg_json, err := json.Marshal(msg)
		if err != nil{
			log.Println(err)
		}*/
		err = c.WriteJSON(msg)
		if err != nil{
			log.Println(err)
		}
	}
}

func (h *Handler) SaveMessage(msg *models.Message) {
	err := h.UseCase.SaveMessage(msg)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) ChatPage(c *gin.Context) {
	tmp, err := template.ParseFiles("client/templates/chat/index.htm")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	err = tmp.Execute(c.Writer, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
