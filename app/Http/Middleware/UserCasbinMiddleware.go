package Middleware

import (
	"larago/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func JWTAuthCasbinMiddleware(bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, exists := c.Get("claims")

		if !exists {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
		}

		userClaims := claims.(*jwt.MapClaims)

		user_role := (*userClaims)["user_role"].(string)
		//Casbinrole
		e := config.CasbinRole()

		sub := user_role
		obj := c.Request.URL.Path
		act := c.Request.Method

		res, err := e.Enforce(sub, obj, act)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"error":  "wrong information" + err.Error(),
			})
			c.Abort()
			return
		}

		if res {
			c.Next()
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"status": 0,
				"error":  "Sorry you do not have this permission",
			})
			c.Abort()
			return
		}

	}
}

func AuthCasbinMiddleware(bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_id")
		//Casbinrole
		e := config.CasbinRole()

		if sessionID == nil {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
		}

		sub := session.Get("user_role")
		obj := c.Request.URL.Path
		act := c.Request.Method

		res, err := e.Enforce(sub, obj, act)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"error":  "wrong information" + err.Error(),
			})
			c.Abort()
			return
		}

		if res {
			c.Next()
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"status": 0,
				"error":  "Sorry you do not have this permission",
			})
			c.Abort()
			return
		}

	}
}
