package authhttp

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
	token, err := c.Cookie("")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, http.ErrNoCookie)
		return
	}

	user_id, err := m.UseCase.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("user_id", user_id)
}
