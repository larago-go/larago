package Controllers

import (
	"larago/app/Model"
	"larago/config"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

func Res_pass(router *gin.RouterGroup) {
	router.POST("/post_add", PostForgotPassword)
	router.GET("/forgot_password", ViewForgotPassword)
	router.POST("/pass/:url/post", ViewRes_passListPost)
	router.GET("/pass/:url", ViewRes_passListPrev)
	router.GET("/api/pass/:url", ApiViewRes_passListPrev)
	router.GET("/api/forgot_password", ApiViewForgotPassword)
}

type Res_passValidation struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

type Res_passPasswordValidation struct {
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

func PostForgotPassword(c *gin.Context) {
	// Validate input
	var input Res_passValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var model Model.UserModel

	if err := config.DB.Where("email = ?", input.Email).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found!"})
		return
	}

	rand_urls := config.RandomString(90)

	//smtp - forgot_password

	toList := []string{input.Email}

	body := []byte("From:" + config.EnvFunc("MAIL_USERNAME") + "\r\n" +
		"To:" + input.Email + "\r\n" +
		"Subject: Password recovery\r\n\r\n" +
		"Link to create a new password" + " " + config.EnvFunc("WWWROOT") + "/login/pass/" + rand_urls + "\r\n")

	auth := smtp.PlainAuth("", config.EnvFunc("MAIL_USERNAME"), config.EnvFunc("MAIL_PASSWORD"), config.EnvFunc("MAIL_HOST"))

	smtp.SendMail(config.EnvFunc("MAIL_HOST")+":"+config.EnvFunc("MAIL_PORT"), auth, config.EnvFunc("MAIL_USERNAME"), toList, body)

	//err := smtp.SendMail(config.EnvFunc("MAIL_HOST")+":"+config.EnvFunc("MAIL_PORT"), auth, config.EnvFunc("MAIL_USERNAME"), toList, body)

	// handling the errors
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//Gorm_SQL

	url_res := Model.ResPassUserModel{Email: input.Email, Url_full: config.EnvFunc("WWWROOT") + "/login/pass/" + rand_urls, Url: rand_urls}
	config.DB.Save(&url_res)

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {

		c.Redirect(http.StatusFound, "/")

	} else {

		c.IndentedJSON(http.StatusOK, gin.H{"data": true})

	}

	//remove link password recovery after 30 minutes

	time.AfterFunc(30*time.Minute, func() {

		var model_url_del []Model.ResPassUserModel

		config.DB.Where("email = ?", input.Email).Find(&model_url_del)
		config.DB.Delete(&model_url_del)

	})
}

func ViewRes_passListPrev(c *gin.Context) { // Get model if exist

	var model Model.ResPassUserModel

	//Gorm_SQL
	if err := config.DB.Where("url = ?", c.Param("url")).First(&model).Error; err != nil {
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
		c.HTML(http.StatusOK, "admin_auth_forgot_password_new.html", gin.H{"csrf": csrf.GetToken(c), "url": model.Url})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	}

}

func ViewRes_passListPost(c *gin.Context) { // Get model if exist

	var model Model.ResPassUserModel
	var user_model Model.UserModel

	//Gorm_SQL
	if err := config.DB.Where("url = ?", c.Param("url")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//end Gorm_SQL

	config.DB.Where("email = ?", model.Email).Find(&user_model)

	var input Res_passPasswordValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bytePassword := []byte(input.Password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	input.Password = string(passwordHash)

	//Gorm_SQL
	config.DB.Model(&user_model).Updates(Model.UserModel{Password: input.Password})

	//c.JSON(http.StatusOK, gin.H{"data": model })

	headerContentTtype := c.Request.Header.Get("Content-Type")

	if headerContentTtype != "application/json" {

		c.Redirect(http.StatusFound, "/auth/login")

	} else {

		c.IndentedJSON(http.StatusOK, gin.H{"data": true})

	}
}

func ViewForgotPassword(c *gin.Context) { // Get model if exist

	template := config.EnvFunc("TEMPLATE")

	switch {

	case template == "vue":

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	case template == "html":

		//HTML template
		c.HTML(http.StatusOK, "admin_auth_forgot_password.html", gin.H{"csrf": csrf.GetToken(c)})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Larago"})

	}

}

func ApiViewForgotPassword(c *gin.Context) { // Get model if exist

	c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c)})
}

func ApiViewRes_passListPrev(c *gin.Context) { // Get model if exist

	var model Model.ResPassUserModel

	//Gorm_SQL
	if err := config.DB.Where("url = ?", c.Param("url")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//end Gorm_SQL

	//c.JSON(http.StatusOK, gin.H{"data": model })
	c.IndentedJSON(http.StatusOK, gin.H{"csrf": csrf.GetToken(c), "url": model.Url})
}
