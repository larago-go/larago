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


type UsersValidation struct {
  Name  string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
  Email string `form:"email" json:"email" binding:"required,email"`
  Role  string `form:"role" json:"role" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

type UsersPasswordValidation struct {
  Name  string `form:"name" json:"name" binding:"required,alphanum,min=4,max=255"`
  Email string `form:"email" json:"email" binding:"required,email"`
  Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

func UsersRegister(router *gin.RouterGroup) {
  router.POST("/post_add", UsersAddPost)
  router.POST("/list/:name/edit", UpdateUsers)
  router.GET("/list/:name/delete", DeleteUsers)
  router.GET("/add", ViewAddUsers)
  router.GET("/list", ViewUsersList)
  router.GET("/list/:name", ViewUsersListPrev)
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

  //c.JSON(http.StatusOK, gin.H{"data": user})
  c.Redirect(http.StatusFound, "/users/list")
}



func UpdateUsers(c *gin.Context) {
  // Get model if exist

  //Gorm_SQL
  var model Model.UserModel

  if err := config.DB.Where("name = ?", c.Param("name")).First(&model).Error; err != nil {
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
  config.DB.Model(&model).Updates(Model.UserModel{ Name: input.Name, Email: input.Email, Role: input.Role, Password: input.Password })
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

  //filter := bson.M{"name": c.Param("name")}

  //update := bson.D{
  //  {"$set", bson.D{
  //      {"name", input.Name},
  //      {"email", input.Email},
  //      {"role", input.Role},
  //      {"password", input.Password},
  //  }},
  //}
  
  //_, err := collection.UpdateOne(ctx, filter, update)

  //if err != nil {
  //  c.JSON(http.StatusBadRequest, gin.H{
  //   "msg": "err collections find one",
  //  })
  //}  
  //end MongoDB


  c.Redirect(http.StatusFound, "/users/list")

}




func DeleteUsers(c *gin.Context) {
  // Get model if exist
  
  //Gorm_SQL
  var model Model.UserModel
  
  if err := config.DB.Where("name = ?", c.Param("name")).First(&model).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  config.DB.Delete(&model)
  //end Gorm_SQL

  //MongoDB
  ////env
  // errenv := godotenv.Load()
  // if errenv != nil {
  //     panic("Error loading .env file")
  //   }
  //DB_DATABASE := os.Getenv("DB_DATABASE")

  //collection := config.MongoClient.Database(DB_DATABASE).Collection("usermodels")

  //ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

  //defer cancel()

  //filter := bson.M{"name": c.Param("name")}
  
  //_, err := collection.DeleteMany(ctx, filter) 

  //if err != nil {
  //  c.JSON(http.StatusBadRequest, gin.H{
  //    "msg": "err collections find one",
  //  })
  //}  
  //end MongoDB

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

    //MongoDB
    //filter := bson.M{}
    
    //// Here's an array in which you can store the decoded documents
    //var model []*Model.UserModel
    
    //// Passing nil as the filter matches all documents in the collection
    ////env
    //errenv := godotenv.Load()
    //if errenv != nil {
    //    panic("Error loading .env file")
    //  }
    //DB_DATABASE := os.Getenv("DB_DATABASE")

    //collection := config.MongoClient.Database(DB_DATABASE).Collection("usermodels")

    //ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

    //defer cancel()

    //cur, err := collection.Find(ctx, filter)
    //if err != nil {
    //    log.Fatal(err)
    //}
    
    //// Finding multiple documents returns a cursor
    //// Iterating through the cursor allows us to decode documents one at a time
    //for cur.Next(ctx) {
    
        //// create a value into which the single document can be decoded
    //    var elem Model.UserModel
    //    err := cur.Decode(&elem)
    //    if err != nil {
    //        log.Fatal(err)
    //    }
    
    //    model = append(model, &elem)
    //}
    
    //if err := cur.Err(); err != nil {
    //    log.Fatal(err)
    //}
    
    //// Close the cursor once finished
    //cur.Close(ctx)
    //end MongoDB
 
  
	  c.HTML(http.StatusOK, "users_list.html", gin.H{ "csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName, "list": model })
}


func ViewUsersListPrev(c *gin.Context) {  // Get model if exist

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
  if err := config.DB.Where("name = ?", c.Param("name")).First(&model).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }
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

  //filter := bson.M{"name": c.Param("name")}
  
  //res := collection.FindOne(ctx, filter).Decode(&model)
  //errmongo := collection.Find(filter)

  //if res != nil {
  //  c.JSON(http.StatusBadRequest, gin.H{
  //    "msg": "err collections find one",
  //  })
  //}  

  //end MongoDB
  //c.JSON(http.StatusOK, gin.H{"data": model })
  c.HTML(http.StatusOK, "users_list_prev.html", gin.H{ "csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName,  "id": model.ID, "name": model.Name,
  "email": model.Email, "role": model.Role, })
}


func ViewAddUsers(c *gin.Context) {  // Get model if exist


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
  c.HTML(http.StatusOK, "users_add.html", gin.H{ "csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName })
}
