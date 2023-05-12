package Controllers

import (
	"larago/config"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Welcome(router *gin.RouterGroup) {

	router.GET("/", GetWelcome)
	router.GET("/welcome", GetWelcome)

}

func GetWelcome(c *gin.Context) {

	//env

	template := config.EnvFunc("TEMPLATE")

	switch {
	case template == "vue":
		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
	case template == "html":
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		//HTML template
		c.HTML(http.StatusOK, "public_welcome.html", gin.H{
			"title":      "Larago",
			"session_id": sessionID})
	default:
		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	}

}
