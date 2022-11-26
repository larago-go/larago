package Controllers

import (
	"larago/app/Model"
	"larago/config"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	csrf "github.com/utrack/gin-csrf"
)

func CasbinRole(router *gin.RouterGroup) {

	router.POST("/post_add", AddPostCasbinRole)
	router.GET("/list/:id/delete", DeleteCasbinRole)
	router.GET("/list", ViewCasbinRole)
	router.GET("/add", AddCasbinRole)
	router.GET("/api/add", ApiAddCasbinRole)
	router.GET("/api/list", ApiViewCasbinRole)
	router.GET("/api/list/:id/delete", ApiDeleteCasbinRole)

}

type CasbinRoleAddValidation struct {
	RoleName string `form:"rolename" json:"rolename" binding:"required"`
	Path     string `form:"path" json:"path" binding:"required"`
	Method   string `form:"method" json:"method" binding:"required"`
}

func AddPostCasbinRole(c *gin.Context) {
	// Validate input
	var input CasbinRoleAddValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := config.CasbinRole()

	e.AddPolicy(input.RoleName, input.Path, input.Method)

	// Create role
	role := Model.CasbinRoleModel{RoleName: input.RoleName, Path: input.Path, Method: input.Method}
	//Gorm_SQL
	config.DB.Save(&role)
	//end_Gorm_SQL

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {

		c.Redirect(http.StatusFound, "/role/list")

	} else {

		c.IndentedJSON(http.StatusCreated, role)

	}
}

func ViewCasbinRole(c *gin.Context) {

	//Gorm_SQL
	var model []Model.CasbinRoleModel
	//end_Gorm_SQL
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
	//end_Gorm_SQL

	//env
	env := godotenv.Load()

	if env != nil {

		panic("Error loading .env file")

	}
	//end_env

	template := os.Getenv("TEMPLATE")

	switch {

	case template == "vue":

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	case template == "html":

		//HTML template
		c.HTML(http.StatusOK, "admin_views_casbin_role.html", gin.H{"session_id": sessionID, "session_name": sessionName, "list": model})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	}

}

func AddCasbinRole(c *gin.Context) {

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

	//env
	env := godotenv.Load()

	if env != nil {

		panic("Error loading .env file")

	}
	//end_env

	template := os.Getenv("TEMPLATE")

	switch {

	case template == "vue":

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	case template == "html":

		//HTML template
		c.HTML(http.StatusOK, "admin_views_casbin_role_add.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

	}

}

func DeleteCasbinRole(c *gin.Context) {
	// Get model if exist

	var model Model.CasbinRoleModel

	//Gorm_SQL
	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	e := config.CasbinRole()

	e.RemovePolicy(model.RoleName, model.Path, model.Method)

	config.DB.Delete(&model)
	//end_Gorm_SQL

	//c.JSON(http.StatusOK, gin.H{"data": true})
	c.Redirect(http.StatusFound, "/role/list")
}

func ApiViewCasbinRole(c *gin.Context) {

	//Gorm_SQL
	var model []Model.CasbinRoleModel
	//end_Gorm_SQL

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {
		//c.JSON(http.StatusForbidden, gin.H{
		//	"message": "not authed",
		//})

		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})

		c.Abort()

	}

	//Gorm_SQL
	config.DB.Find(&model)
	//end_Gorm_SQL

	c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName, "list": model})

}

func ApiAddCasbinRole(c *gin.Context) {

	session := sessions.Default(c)
	sessionID := session.Get("user_id")
	sessionName := session.Get("user_name")

	if sessionID == nil {
		//c.JSON(http.StatusForbidden, gin.H{
		//	"message": "not authed",
		//})

		c.IndentedJSON(http.StatusOK, gin.H{"csrf": "redirect_auth_login"})

		c.Abort()

	}

	//c.JSON(http.StatusOK, gin.H{"data": model})
	c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName})

}

func ApiDeleteCasbinRole(c *gin.Context) {
	// Get model if exist

	var model Model.CasbinRoleModel

	//Gorm_SQL
	if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	e := config.CasbinRole()

	e.RemovePolicy(model.RoleName, model.Path, model.Method)

	config.DB.Delete(&model)
	//end_Gorm_SQL

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
