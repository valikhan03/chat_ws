package delivery

import(
	"net/http"
	"chatapp/accounts"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	UseCase accounts.UseCase
}

func NewHandler(uc accounts.UseCase) *Handler{
	return &Handler{
		UseCase: uc,
	}
}

func(h *Handler) FindUser(c *gin.Context){
	username := c.Param("user")

	user_id, err := h.UseCase.FindUser(username)
	if(user_id == "" && err != nil){
		c.AbortWithStatusJSON(404, gin.H{"error":"User not found"});
		return
	}

	c.JSON(http.StatusFound, user_id)
}