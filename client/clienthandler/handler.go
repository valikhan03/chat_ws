package clienthandler


import(
	"log"
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
)

func ChatPage(c *gin.Context) {
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