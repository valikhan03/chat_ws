package delivery

import (
	"chatapp/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	UseCase auth.UseCase
}

func NewAuthMiddleware(usecase auth.UseCase) *AuthMiddleware {
	return &AuthMiddleware{
		UseCase: usecase,
	}
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	cookie, err := c.Request.Cookie("access-token-chat-eltaev")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, http.ErrNoCookie)
		return
	}
	token := cookie.Value

	fmt.Println("Token in cookie:", token)

	_, err = m.UseCase.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Next()
}
