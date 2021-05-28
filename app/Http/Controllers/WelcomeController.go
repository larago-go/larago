package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWelcome(c *gin.Context) {

	c.HTML(http.StatusOK, "welcome.html", gin.H{"title": "Larago"})
}
