package clienthandler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatPage(c *gin.Context) {
	tmp, err := template.ParseFiles("client/templates/chat/index.htm")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(c.Writer, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func ChatListsPage(c *gin.Context) {
	tmp, err := template.ParseFiles("client/templates/chatslist/index.htm")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(c.Writer, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func SignUpPage(c *gin.Context) {
	tmp, err := template.ParseFiles("client/templates/sign-up/sign_up.htm")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(c.Writer, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func SignInPage(c *gin.Context) {
	tmp, err := template.ParseFiles("client/templates/sign-in/sign_in.htm")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(c.Writer, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
