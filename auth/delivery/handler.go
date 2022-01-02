package delivery

import (
	"chatapp/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase auth.UseCase
}

func NewHadler(usecase auth.UseCase) *Handler {
	return &Handler{
		UseCase: usecase,
	}
}

type SignUpInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	input := new(SignUpInput)

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.UseCase.SignUp( input.Username, input.Email, input.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(c *gin.Context) {
	input := new(SignInInput)
	err := c.BindJSON(&input)	
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.UseCase.SignIn(input.Email, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invilid email or password"})
		return
	}

	c.SetCookie("access-token", token, 1000000, "/", "localhost", false, true)
}
