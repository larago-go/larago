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

func UsersRegister(router *gin.RouterGroup) {

	router.POST("/post_add", UsersAddPost)
	router.POST("/list/:id/edit", UpdateUsers)
	router.PUT("/api/list/:id/edit", UpdateUsers)
	router.GET("/list/:id/delete", DeleteUsers)
	router.GET("/add", ViewAddUsers)
	router.GET("/list", ViewUsersList)
	router.GET("/list/:id", ViewUsersListPrev)
	router.GET("/api/list", ApiViewUsersList)
	router.GET("/api/add", ApiViewAddUsers)
	router.GET("/api/list/:id", ApiViewUsersListPrev)
	router.DELETE("/api/list/:id/delete", ApiDeleteUsers)

}

type UsersValidation struct {
	Name     string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Role     string `form:"role" json:"role"`
	Password string `form:"password" json:"password"`
}

func UsersAddPost(c *gin.Context) {
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

	// Create user
	user := Model.UserModel{
		Name:     input.Name,
		Role:     input.Role,
		Email:    input.Email,
		Password: input.Password,
	}

	//Gorm_SQL
	config.DB.Save(&user)
	//end Gorm_SQL

	//c.JSON(http.StatusOK, gin.H{"data": user})

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {

		c.Redirect(http.StatusFound, "/users/list")

	} else {

		c.IndentedJSON(http.StatusCreated, user)

	}

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

	if len(input.Password) > 0 {

		bytePassword := []byte(input.Password)

		passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

		input.Password = string(passwordHash)

		//Gorm_SQL
		config.DB.Model(&model).Select(
			"name",
			"email",
			"role",
			"password",
		).Updates(Model.UserModel{
			Name:     input.Name,
			Email:    input.Email,
			Role:     input.Role,
			Password: input.Password,
		})
		//end Gorm_SQL

	} else {

		//Gorm_SQL
		config.DB.Model(&model).Select(
			"name",
			"email",
			"role",
			"password",
		).Updates(Model.UserModel{
			Name:  input.Name,
			Email: input.Email,
			Role:  input.Role,
		})
		//end Gorm_SQL

	}

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {

		c.Redirect(http.StatusFound, "/users/list")

	} else {

		c.IndentedJSON(http.StatusOK, model)

	}

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

	c.Redirect(http.StatusFound, "/users/list")
}

func ViewUsersList(c *gin.Context) {

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

		//Gorm_SQL
		var model []Model.UserModel

		config.DB.Find(&model)
		//end Gorm_SQL

		//HTML template
		c.HTML(http.StatusOK, "admin_views_users_list.html", gin.H{
			"csrf":         csrf.GetToken(c),
			"session_id":   sessionID,
			"session_name": sessionName,
			"list":         model})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	}

}

func ViewUsersListPrev(c *gin.Context) { // Get model if exist

	var model Model.UserModel

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {

		c.Redirect(http.StatusFound, "/auth/login")

		c.Abort()
	}
	//Gorm_SQL
	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}
	//end Gorm_SQL

	template := config.EnvFunc("TEMPLATE")

	switch {

	case template == "vue":

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	case template == "html":

		//HTML template
		c.HTML(http.StatusOK, "admin_views_users_list_prev.html", gin.H{
			"csrf":         csrf.GetToken(c),
			"session_id":   sessionID,
			"session_name": sessionName,
			"id":           model.ID,
			"name":         model.Name,
			"email":        model.Email,
			"role":         model.Role,
		})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	}

}

func ViewAddUsers(c *gin.Context) { // Get model if exist

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
		c.HTML(http.StatusOK, "admin_views_users_add.html", gin.H{
			"csrf":         csrf.GetToken(c),
			"session_id":   sessionID,
			"session_name": sessionName})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	}

}

func ApiViewUsersList(c *gin.Context) {

	//Gorm_SQL
	var model []Model.UserModel
	//end Gorm_SQL

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {

		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})

		c.Abort()

	}

	//Gorm_SQL
	config.DB.Find(&model)
	//end Gorm_SQL

	c.IndentedJSON(http.StatusOK, gin.H{
		"csrf":         csrf.GetToken(c),
		"session_id":   sessionID,
		"session_name": sessionName,
		"list":         model})

}

func ApiViewAddUsers(c *gin.Context) { // Get model if exist

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {

		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})

		c.Abort()

	}

	//c.JSON(http.StatusOK, gin.H{"data": model})
	c.IndentedJSON(http.StatusOK, gin.H{
		"csrf":         csrf.GetToken(c),
		"session_id":   sessionID,
		"session_name": sessionName})

}

func ApiViewUsersListPrev(c *gin.Context) { // Get model if exist

	var model Model.UserModel

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {

		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})

		c.Abort()
	}
	//Gorm_SQL
	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {

		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})

		return
	}
	//end Gorm_SQL

	c.IndentedJSON(http.StatusOK, gin.H{
		"csrf":         csrf.GetToken(c),
		"session_id":   sessionID,
		"session_name": sessionName,
		"id":           model.ID,
		"name":         model.Name,
		"email":        model.Email,
		"role":         model.Role,
	})

}

func ApiDeleteUsers(c *gin.Context) {

	//Gorm_SQL
	var model Model.UserModel

	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&model)
	//end Gorm_SQL

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
