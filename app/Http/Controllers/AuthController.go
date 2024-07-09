package Controllers

import (
	"larago/app/Model"
	"larago/config"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

func Auth(router *gin.RouterGroup) {

	router.POST("/signup", UsersRegistration)
	router.POST("/signin", UsersLogin)
	router.GET("/signout", Loginout)
	router.GET("/login", ViewUsersLogin)
	router.GET("/register", ViewUsersRegistration)
	router.GET("/api/register", ApiViewUsersRegistration)
	router.GET("/api/login", ApiViewUsersLogin)
	router.GET("/api/session", ViewUserSession)
	router.GET("/api/signout", ApiLoginout)
}

type PasswordValidation struct {
	Name     string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

type LoginValidation struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password"json:"password" binding:"required,min=8,max=255"`
}

func UsersRegistration(c *gin.Context) {

	// Validate input
	var input PasswordValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytePassword := []byte(input.Password)

	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	input.Password = string(passwordHash)

	// Create user
	user := Model.UserModel{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	//Gorm_SQL
	config.DB.Save(&user)

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {
		c.Redirect(http.StatusFound, "/home")
	} else {
		c.IndentedJSON(http.StatusCreated, user)
	}

}

func UsersLogin(c *gin.Context) {

	// Validate input
	var input LoginValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var model Model.UserModel

	//Gorm_SQL
	config.DB.Where("email = ?", input.Email).First(&model)

	bytePassword := []byte(input.Password)

	byteHashedPassword := []byte(model.Password)

	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Password mismatch",
		})
		return
	} else {
		session := sessions.Default(c)
		session.Set("user_id", model.ID)
		session.Set("user_email", model.Email)
		session.Set("user_name", model.Name)
		//Casbinrole
		session.Set("user_role", model.Role)
		session.Save()

		headerContentTtype := c.Request.Header.Get("Content-Type")

		if headerContentTtype != "application/json" {
			c.Redirect(http.StatusFound, "/home")
		} else {
			c.IndentedJSON(http.StatusCreated, gin.H{
				"message": "User signed in",
				"user":    model.Name,
				"id":      model.ID,
			})
		}

	}

}

func Loginout(c *gin.Context) {

	session := sessions.Default(c)

	session.Clear()

	session.Save()

	c.Redirect(http.StatusFound, "/")

}

func ViewUsersLogin(c *gin.Context) {

	session := sessions.Default(c)

	sessionID := session.Get("user_id")

	if sessionID == nil {
		//env
		template := config.EnvFunc("TEMPLATE")

		switch {
		case template == "vue":
			//VUE template
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
		case template == "html":
			//HTML template
			c.HTML(http.StatusOK, "admin_auth_login.html", gin.H{"csrf": csrf.GetToken(c)})
		default:
			//VUE template
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
		}

	} else {
		c.Redirect(http.StatusFound, "/home")
	}

}

func ViewUsersRegistration(c *gin.Context) {

	session := sessions.Default(c)

	sessionID := session.Get("user_id")

	if sessionID == nil {
		//env
		template := config.EnvFunc("TEMPLATE")

		switch {
		case template == "vue":
			//VUE template
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
		case template == "html":
			//HTML template
			c.HTML(http.StatusOK, "admin_auth_register.html", gin.H{"csrf": csrf.GetToken(c)})
		default:
			//VUE template
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})
		}

	} else {
		c.Redirect(http.StatusFound, "/home")
	}

}

func ApiViewUsersRegistration(c *gin.Context) {

	session := sessions.Default(c)

	sessionID := session.Get("user_id")

	if sessionID == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_home"})
	}

	// RegisterAuth.vue
}

func ApiViewUsersLogin(c *gin.Context) {

	session := sessions.Default(c)

	sessionID := session.Get("user_id")

	if sessionID == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_home"})
	}

	//LoginAuth.vue

}

func ViewUserSession(c *gin.Context) {

	session := sessions.Default(c)

	sessionID := session.Get("user_id")

	if sessionID == nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"userid_session_id": "no_auth",
			"userid_session":    "no_auth",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"userid_session_id": sessionID,
			"userid_session":    "auth",
		})
	}

}

func ApiLoginout(c *gin.Context) {

	session := sessions.Default(c)

	session.Clear()

	session.Save()

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Signed out..."})

}
