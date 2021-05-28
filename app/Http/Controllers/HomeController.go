package Controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Home(router *gin.RouterGroup) {

	router.GET("/", ViewHome)

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

	c.HTML(http.StatusOK, "home.html", gin.H{"session_id": sessionID, "session_name": sessionName})
}

