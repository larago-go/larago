package Controllers

import (
	"larago/app/Model"
	"larago/config"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/crypto/bcrypt"
)

type Res_passValidation struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

type Res_passPasswordValidation struct {
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

func Res_pass(router *gin.RouterGroup) {
	router.POST("/post_add", PostForgotPassword)
	router.GET("/forgot_password", ViewForgotPassword)
	router.POST("/pass/:url/post", ViewRes_passListPost)
	router.GET("/pass/:url", ViewRes_passListPrev)
	router.GET("/api/pass/:url", ApiViewRes_passListPrev)
	router.GET("/api/forgot_password", ApiViewForgotPassword)
}

func PostForgotPassword(c *gin.Context) {
	// Validate input
	var input Res_passValidation

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var model Model.UserModel

	//env
	errenv := godotenv.Load()
	if errenv != nil {
		panic("Error loading .env file")
	}

	//Gorm_SQL

	if err := config.DB.Where("email = ?", input.Email).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found!"})
		return
	}

	rand_urls := config.RandomString(90)

	//smtp - forgot_password

	toList := []string{input.Email}

	body := []byte("From:" + os.Getenv("MAIL_USERNAME") + "\r\n" +
		"To:" + input.Email + "\r\n" +
		"Subject: Password recovery\r\n\r\n" +
		"Link to create a new password" + " " + os.Getenv("WWWROOT") + "/login/pass/" + rand_urls + "\r\n")

	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_USERNAME"), toList, body)

	//err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_USERNAME"), toList, body)

	// handling the errors
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//Gorm_SQL

	url_res := Model.ResPassUserModel{Email: input.Email, Url_full: os.Getenv("WWWROOT") + "/login/pass/" + rand_urls, Url: rand_urls}
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
		c.HTML(http.StatusOK, "forgot_password_new.html", gin.H{"csrf": csrf.GetToken(c), "url": model.Url})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

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
		c.HTML(http.StatusOK, "forgot_password.html", gin.H{"csrf": csrf.GetToken(c)})

	default:

		//VUE template
		c.HTML(http.StatusOK, "index_vue.html", gin.H{"title": "Larago"})

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
