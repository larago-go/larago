package Controllers

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	csrf "github.com/utrack/gin-csrf"

	"net/http"
)

func Home(router *gin.RouterGroup) {

	router.GET("/", ViewHome)
	router.GET("/api", ApiViewHome)

}

func ViewHome(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")
	if sessionID == nil {
		//c.JSON(http.StatusForbidden, gin.H{
		//	"message": "not authed",
		//})
		c.Redirect(http.StatusFound, "/auth/login")
		c.Abort()
	}

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
		c.HTML(http.StatusOK, "home.html", gin.H{"session_id": sessionID, "session_name": sessionName})
	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	}

}

func ApiViewHome(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")
	if sessionID == nil {
		//c.JSON(http.StatusForbidden, gin.H{
		//	"message": "not authed",
		//})
		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})

		c.Abort()

	}

	//c.JSON(http.StatusOK, gin.H{"data": model})
	c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName})

}
