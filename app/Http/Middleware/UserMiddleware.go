package Middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")

		if sessionID == nil {
			//	c.JSON(http.StatusForbidden, gin.H{
			//		"message": "not authed",
			//	})
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
		}

	}
}

