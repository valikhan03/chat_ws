package wsdelivery

import (
	"chatapp/chat"

	"github.com/gin-gonic/gin"
)

func RegisterChatHTTPWSEndpoints(router *gin.Engine, uc chat.UseCase) {
	h := NewHandler(uc)

	chat := router.Group("/chat")
	{
		chat.POST("/demochat", h.WSEndpoint)
	}

}
