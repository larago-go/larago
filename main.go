package main

import (
    "os"
	"github.com/gin-gonic/gin"
	"larago/config"
	"larago/app/Http/Controllers"
	"larago/app/Http/Middleware"
	"github.com/utrack/gin-csrf"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/sessions"
	
	//sessions_redis
	//"github.com/gin-contrib/sessions/redis"
	//end_sessions_redis
	//sessions_cookie
	"github.com/gin-contrib/sessions/cookie"
	//end_sessions_cookie
	//Memcached
	//"github.com/bradfitz/gomemcache/memcache"
	//"github.com/gin-contrib/sessions/memcached"
	//end_Memcached

)



func main() {

    //database_SQL
	config.Init()
	//end_database_SQL
	
	//Mongodb
	//config.Init_Mongo()
	//end_Mongodb

	//env
	errenv := godotenv.Load()
	if errenv != nil {
		panic("Error loading .env file")
	}
    //end_env

	APP_KEYS := os.Getenv("APP_KEYS")

	//gin
	//switch to "release" mode in production
	
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	//sessions

	//sessions_cookie
	store := cookie.NewStore([]byte(APP_KEYS))
	//end_sessions_cookie
	
	//redis_sessions
	//REDIS_HOST := os.Getenv("REDIS_HOST")
	//REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	//REDIS_PORT := os.Getenv("REDIS_PORT")
	//REDIS_SECRET := os.Getenv("REDIS_SECRET")

	//store, err := redis.NewStore(10, "tcp", REDIS_HOST+":"+REDIS_PORT, REDIS_PASSWORD, []byte(REDIS_SECRET))
		
	//if err != nil {
	//		panic("Failed to connect to redis_sessions!")
	//	  }
	//redis_sessions

	//Memcached
	//store := memcached.NewStore(memcache.New("localhost:11211"), "", []byte("APP_KEYS"))
	//end_Memcached

    //sessions_use
	r.Use(sessions.Sessions("larago", store))
    //end_sessions
    
	//CSRF_Middleware
	r.Use(csrf.Middleware(csrf.Options{
        Secret: APP_KEYS,
        ErrorFunc: func(c *gin.Context) {
            c.String(400, "CSRF token mismatch")
            c.Abort()                                                                                                                                                                             
        },
	}))
	//end_CSRF_Middleware


    //gin_html_and_static
	r.Static("/public", "./public")
	r.Static("/node_modules", "./node_modules")
	r.LoadHTMLGlob("resources/views/*.html")
	//end_gin_html_and_static


    //gin_route_middleware
    
	r.GET("/", Controllers.GetWelcome)

    auth := r.Group("/auth")
	Controllers.Auth(auth.Group("/"))

	//Casbin_Role_Middleware
	//r.Use(Middleware.AuthCasbinMiddleware(true))
    //end_Casbin_Role_Middleware

    //Auth_Middleware
	r.Use(Middleware.AuthMiddleware(true))
	//end_Auth_Middleware

	users := r.Group("/users")
	Controllers.UsersRegister(users.Group("/"))
	
	home := r.Group("/home")
	Controllers.Home(home.Group("/"))

	role := r.Group("/role")
	Controllers.CasbinRole(role.Group("/"))
	//end_gin_route_middleware

	//test
	//test := r.Group("/api/ping")
	//test.Use(Middleware.AuthMiddleware(true))

	//test.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//end_test

    PORT := os.Getenv("PORT")
	r.Run(PORT) // listen and serve on 0.0.0.0:8080
}

