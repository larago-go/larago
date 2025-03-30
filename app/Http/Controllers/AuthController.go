package Controllers

import (
	"larago/app/Model"
	"larago/config"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"user_email": user.Email,
		"user_name":  user.Name,
		"user_role":  user.Role,
		//session time
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.EnvFunc("APP_KEYS")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {
		c.Redirect(http.StatusFound, "/home")
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"user_name":  user.Name,
			"user_email": user.Email,
			"user_id":    user.ID,
			"token":      tokenString,
		})

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

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":    model.ID,
			"user_email": model.Email,
			"user_name":  model.Name,
			"user_role":  model.Role,
			//session time
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})

		tokenString, err := token.SignedString([]byte(config.EnvFunc("APP_KEYS")))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		headerContentTtype := c.Request.Header.Get("Content-Type")

		if headerContentTtype != "application/json" {
			c.Redirect(http.StatusFound, "/home")
		} else {
			c.IndentedJSON(http.StatusCreated, gin.H{
				"user_name":  model.Name,
				"user_email": model.Email,
				"user_id":    model.ID,
				"token":      tokenString,
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

	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok..."})

	// RegisterAuth.vue
}

func ApiViewUsersLogin(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok..."})

	//LoginAuth.vue

}

func ApiLoginout(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Signed out..."})

}
