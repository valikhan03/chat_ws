package delivery

import (
	"chatapp/auth"
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
		c.AbortWithError(http.StatusUnauthorized, http.ErrNoCookie)
		return
	}
	token := cookie.Value


	_, err = m.UseCase.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
