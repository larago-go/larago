package Controllers

import (
	"net/http"
	"os"

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

		//HTML template
		c.HTML(http.StatusOK, "welcome.html", gin.H{"title": "Larago"})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	}

}
