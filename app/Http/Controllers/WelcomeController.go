package Controllers

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Welcome(router *gin.RouterGroup) {

	router.GET("/", GetWelcome)
	router.GET("/welcome", GetWelcome)

}

func GetWelcome(c *gin.Context) {

	//env
	env := godotenv.Load()

	if env != nil {

		panic("Error loading .env file")

	}
	//end_env

	template := os.Getenv("TEMPLATE")

	switch {

	case template == "vue":

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	case template == "html":

		session := sessions.Default(c)

		sessionID := session.Get("user_id")

		//HTML template
		c.HTML(http.StatusOK, "public_welcome.html", gin.H{"title": "Larago", "session_id": sessionID})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	}

}
