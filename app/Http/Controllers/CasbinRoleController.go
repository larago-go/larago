package Controllers

import (
  "larago/config"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"
  "larago/app/Model"
  "github.com/utrack/gin-csrf"
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


type CasbinRoleAddValidation struct {
	RoleName  string `form:"rolename" json:"rolename" binding:"required"`
	Path string `form:"path" json:"path" binding:"required"`
	Method string `form:"method" json:"method" binding:"required"`
}

func CasbinRole(router *gin.RouterGroup) {

  router.POST("/post_add", AddPostCasbinRole)
  router.GET("/list/:v0/delete", DeleteCasbinRole)
  router.GET("/list", ViewCasbinRole)
  router.GET("/add", AddCasbinRole)
  
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
	role := Model.CasbinRoleModel{ RoleName: input.RoleName, Path: input.Path, Method: input.Method }
	//Gorm_SQL
	config.DB.Save(&role)
	//end_Gorm_SQL
	
    //MongoDB
    ////env
    //errenv := godotenv.Load()
    //  if errenv != nil {
    //      panic("Error loading .env file")
    //    }
    //DB_DATABASE := os.Getenv("DB_DATABASE")

    //collection := config.MongoClient.Database(DB_DATABASE).Collection("casbinrolemodels")
    //ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

    //defer cancel()

    //_, err_post := collection.InsertOne(ctx, role)

    //if err_post != nil {
    //  c.JSON(http.StatusBadRequest, gin.H{
    //  "msg": "A role with the same name already exists",
    //})
    //}
    //opt := options.Index()
    //opt.SetUnique(true)
    //index := mongo.IndexModel{Keys: bson.M{"v0": 1}, Options: opt}
    //if _, err := collection.Indexes().CreateOne(ctx, index); err != nil {
    //log.Println("Could not create index:", err)
    //}
    //end MongoDB
  
	//c.JSON(http.StatusOK, gin.H{"data": role})
	c.Redirect(http.StatusFound, "/role/list")
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

	//MongoDB
	//	filter := bson.M{}
    
	//	// Here's an array in which you can store the decoded documents
	//	var model []*Model.CasbinRoleModel
		
	//	// Passing nil as the filter matches all documents in the collection
	//	 //env
	//  errenv := godotenv.Load()
	//	if errenv != nil {
	//		panic("Error loading .env file")
	//	  }
	//  DB_DATABASE := os.Getenv("DB_DATABASE")
	
	//  collection := config.MongoClient.Database(DB_DATABASE).Collection("casbinrolemodels")
	
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	
	//	defer cancel()
	
	//	cur, err := collection.Find(ctx, filter)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
		
	//	// Finding multiple documents returns a cursor
	//	// Iterating through the cursor allows us to decode documents one at a time
	//	for cur.Next(ctx) {
		
	//		// create a value into which the single document can be decoded
	//		var elem Model.CasbinRoleModel
	//		err := cur.Decode(&elem)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
		
	//		model = append(model, &elem)
	//	}
		
	//	if err := cur.Err(); err != nil {
	//		log.Fatal(err)
	//	}
		
	//	// Close the cursor once finished
	//	cur.Close(ctx)
	//end MongoDB
  
	  c.HTML(http.StatusOK, "casbin_role.html", gin.H{ "session_id": sessionID, "session_name": sessionName, "list": model })
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
  
	  c.HTML(http.StatusOK, "casbin_role_add.html", gin.H{ "csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName })
  }

  func DeleteCasbinRole(c *gin.Context) {
	// Get model if exist

	var model Model.CasbinRoleModel
	
	//Gorm_SQL
	if err := config.DB.Where("v0 = ?", c.Param("v0")).First(&model).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
	
	e := config.CasbinRole()

	e.RemovePolicy(model.RoleName, model.Path, model.Method)

	config.DB.Delete(&model)
	//end_Gorm_SQL

   //MongoDB
   ////env
   //errenv := godotenv.Load()
   //if errenv != nil {
   //    panic("Error loading .env file")
   //  }
   //DB_DATABASE := os.Getenv("DB_DATABASE")

   //collection := config.MongoClient.Database(DB_DATABASE).Collection("casbinrolemodels")

   //ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

   //defer cancel()

   //filter := bson.M{"v0": c.Param("v0")}

   //res := collection.FindOne(ctx, filter).Decode(&model)

   //if res != nil {
   //  c.JSON(http.StatusBadRequest, gin.H{
   //   "msg": "err collections find one",
   // })
   //} 
  
   //e := config.CasbinRole()

   //e.RemovePolicy(model.RoleName, model.Path, model.Method)

   //_, err := collection.DeleteMany(ctx, filter) 

   //if err != nil {
   //  c.JSON(http.StatusBadRequest, gin.H{
   //    "msg": "err collections find one",
   // })
   //}  
   //end MongoDB
  
	//c.JSON(http.StatusOK, gin.H{"data": true})
	c.Redirect(http.StatusFound, "/role/list")
  }
