package Controllers

import (
	"larago/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		c.Redirect(http.StatusFound, "/auth/login")
		c.Abort()
	}

	template := config.EnvFunc("TEMPLATE")

	switch {
	case template == "vue":
		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
	case template == "html":
		//HTML template
		c.HTML(http.StatusOK, "admin_views_home.html", gin.H{
			"session_id":   sessionID,
			"session_name": sessionName,
		})
	default:
		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
	}

}

func ApiViewHome(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})
		c.Abort()
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"csrf":         csrf.GetToken(c),
		"session_id":   sessionID,
		"session_name": sessionName,
	})

}
