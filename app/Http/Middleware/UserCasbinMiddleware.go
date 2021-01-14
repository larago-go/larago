package Middleware

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
	"larago/config"

)



func AuthCasbinMiddleware(bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		//Casbinrole
		e := config.CasbinRole()
		
		if sessionID == nil {
			//c.JSON(http.StatusForbidden, gin.H{
			//	"message": "not authed",
			//})
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
		}
       sub := session.Get("user_role")
       obj := c.Request.URL.Path
       act := c.Request.Method

		res, err := e.Enforce(sub, obj, act)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "wrong information" + err.Error(),
			})
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "Sorry you do not have this permission",
			})
			c.Abort()
			return
		}


	}
}

