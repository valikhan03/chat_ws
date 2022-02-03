package clienthandler

import(
	"github.com/gin-gonic/gin"
)


func CacheContolMiddleware(c *gin.Context){
	c.Writer.Header().Add("Cache-Control", "no-cache, must-revalidate")
	c.Next()
}