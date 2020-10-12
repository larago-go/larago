package Controllers

import (

	"github.com/gin-gonic/gin"
	"net/http"

	
)


func GetWelcome(c *gin.Context) {

	c.HTML(http.StatusOK, "welcome.html", gin.H{ "title": "Larago", })
}
