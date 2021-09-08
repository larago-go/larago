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

type PasswordValidation struct {
	Name     string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

type LoginValidation struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password"json:"password" binding:"required,min=8,max=255"`
}

func Auth(router *gin.RouterGroup) {
	router.POST("/signup", UsersRegistration)
	router.POST("/signin", UsersLogin)
	router.GET("/signout", Loginout)
	router.GET("/login", ViewUsersLogin)
	router.GET("/register", ViewUsersRegistration)
}

func UsersRegistration(c *gin.Context) {
	// Validate input
	var input PasswordValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytePassword := []byte(input.Password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	input.Password = string(passwordHash)

	// Create user
	user := Model.UserModel{Name: input.Name, Email: input.Email, Password: input.Password}

	//Gorm_SQL
	config.DB.Save(&user)
	//end Gorm_SQL

	// c.JSON(http.StatusOK, gin.H{"data": insertResult.InsertedID, "data1": user})
	c.Redirect(http.StatusFound, "/home")
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
	//end Gorm_SQL

	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(model.Password)
	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "Password mismatch",
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

		//c.JSON(http.StatusOK, gin.H{"message": "User signed in", "user": model.Name, "id": model.ID})

		c.Redirect(http.StatusFound, "/home")
	}

}

func Loginout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/")
	//c.JSON(http.StatusOK, gin.H{"message": "Signed out..."})
}

func ViewUsersLogin(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("user_id")

	if sessionID == nil {
		//c.JSON(http.StatusForbidden, gin.H{
		//	"message": "not authed",
		//})
		//c.Redirect(http.StatusFound, "/auth/login")
		//c.Abort()
		c.HTML(http.StatusOK, "login.html", gin.H{"csrf": csrf.GetToken(c)})
	} else {
		c.Redirect(http.StatusFound, "/home")
	}

}

func ViewUsersRegistration(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("user_id")

	if sessionID == nil {
		//c.JSON(http.StatusForbidden, gin.H{
		//	"message": "not authed",
		//})
		//c.Redirect(http.StatusFound, "/auth/login")
		//c.Abort()
		c.HTML(http.StatusOK, "register.html", gin.H{"csrf": csrf.GetToken(c)})
	} else {
		c.Redirect(http.StatusFound, "/home")
	}

}

