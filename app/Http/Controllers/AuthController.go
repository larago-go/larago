package Controllers

import (

	"larago/config"
  "github.com/gin-gonic/gin"
  "larago/app/Model"
  "golang.org/x/crypto/bcrypt"
  "github.com/utrack/gin-csrf"
  "github.com/gin-contrib/sessions"
  "net/http"
  //MongoDB
  //"go.mongodb.org/mongo-driver/bson"
  //"go.mongodb.org/mongo-driver/mongo"
  //"go.mongodb.org/mongo-driver/mongo/options"
  //"time"
  //"log"
  //"context"
  //"os"
	//"github.com/joho/godotenv"
  //end MongoDB

)


type PasswordValidation struct {
	Name  string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
	Email string `form:"email" json:"email" binding:"required,email"`
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
  user := Model.UserModel{ Name: input.Name, Email: input.Email, Password: input.Password }

  //Gorm_SQL
  config.DB.Save(&user)
  //end Gorm_SQL


  //MongoDB

  ////env
  //errenv := godotenv.Load()
  //  if errenv != nil {
  //      panic("Error loading .env file")
  //    }
  //DB_DATABASE := os.Getenv("DB_DATABASE")

  //collection := config.MongoClient.Database(DB_DATABASE).Collection("usermodels")
  //ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

  //defer cancel()

  //_, err_post := collection.InsertOne(ctx, user)

  //if err_post != nil {
  //  c.JSON(http.StatusBadRequest, gin.H{
  //    "msg": "A user with the same name already exists",
  //  })
  //}
  //opt := options.Index()
  //opt.SetUnique(true)
  //index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}
  //if _, err := collection.Indexes().CreateOne(ctx, index); err != nil {
  //log.Println("Could not create index:", err)
  //}
  //end MongoDB

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

  //MongoDB

  ////env
  //errenv := godotenv.Load()
  //  if errenv != nil {
  //      panic("Error loading .env file")
  //    }
  //DB_DATABASE := os.Getenv("DB_DATABASE")

  //collection := config.MongoClient.Database(DB_DATABASE).Collection("usermodels")

  //ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

  //defer cancel()

  //filter := bson.M{"email": input.Email}
  
  //errmongo := collection.FindOne(ctx, filter).Decode(&model)

  //if errmongo != nil {

  //log.Fatal("err collections users")

  //}
  //end MongoDB

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
    c.HTML(http.StatusOK, "login.html", gin.H{ "csrf": csrf.GetToken(c), })
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
    c.HTML(http.StatusOK, "register.html", gin.H{ "csrf": csrf.GetToken(c), })
  } else {
    c.Redirect(http.StatusFound, "/home")
  }

}
