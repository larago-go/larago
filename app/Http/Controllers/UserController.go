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

type UsersValidation struct {
	Name     string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Role     string `form:"role" json:"role" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

type UsersPasswordValidation struct {
	Name     string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/post_add", UsersAddPost)
	router.POST("/list/:id/edit", UpdateUsers)
	router.GET("/list/:id/delete", DeleteUsers)
	router.GET("/add", ViewAddUsers)
	router.GET("/list", ViewUsersList)
	router.GET("/list/:id", ViewUsersListPrev)
}

func UsersAddPost(c *gin.Context) {
	// Validate input
	var input UsersPasswordValidation

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

	//c.JSON(http.StatusOK, gin.H{"data": user})
	c.Redirect(http.StatusFound, "/users/list")
}

func UpdateUsers(c *gin.Context) {
	// Get model if exist

	//Gorm_SQL
	var model Model.UserModel

	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//end Gorm_SQL

	// Validate input
	var input UsersValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytePassword := []byte(input.Password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	input.Password = string(passwordHash)

	//Gorm_SQL
	config.DB.Model(&model).Updates(Model.UserModel{Name: input.Name, Email: input.Email, Role: input.Role, Password: input.Password})
	//end Gorm_SQL

	c.Redirect(http.StatusFound, "/users/list")

}

func DeleteUsers(c *gin.Context) {
	// Get model if exist

	//Gorm_SQL
	var model Model.UserModel

	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&model)
	//end Gorm_SQL

	//c.JSON(http.StatusOK, gin.H{"data": true})
	c.Redirect(http.StatusFound, "/users/list")
}

func ViewUsersList(c *gin.Context) {

	//Gorm_SQL
	var model []Model.UserModel
	//end Gorm_SQL
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
	//Gorm_SQL
	config.DB.Find(&model)
	//end Gorm_SQL

	c.HTML(http.StatusOK, "users_list.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName, "list": model})
}

func ViewUsersListPrev(c *gin.Context) { // Get model if exist

	var model Model.UserModel

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
	//Gorm_SQL
	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//end Gorm_SQL

	//c.JSON(http.StatusOK, gin.H{"data": model })
	c.HTML(http.StatusOK, "users_list_prev.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName, "id": model.ID, "name": model.Name,
		"email": model.Email, "role": model.Role})
}

func ViewAddUsers(c *gin.Context) { // Get model if exist

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

	//c.JSON(http.StatusOK, gin.H{"data": model})
	c.HTML(http.StatusOK, "users_add.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName})
}
