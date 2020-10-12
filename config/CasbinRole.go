package config

import (

	"github.com/casbin/casbin/v2"
	"os"
	"github.com/joho/godotenv"
	//database_SQL
	gormadapter "github.com/casbin/gorm-adapter/v3"
	//end_database_SQL
	//mongodb
	//"github.com/casbin/mongodb-adapter/v2"
	//end_mongodb
)



func CasbinRole() *casbin.Enforcer {
	
		//env
	errenv := godotenv.Load()
	 if errenv != nil {
	    	panic("Error loading .env file")
		}
	
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	//database_SQL
	DB_DATABASE := os.Getenv("DB_DATABASE")
	//end_database_SQL
	
	//mysql	
	a, errcasbindb := gormadapter.NewAdapter("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE, true)
    //postgres
	//a, errcasbindb := gormadapter.NewAdapter("postgres", "host="+DB_HOST+" port="+DB_PORT+" user="+DB_USERNAME+" dbname="+DB_DATABASE+" password="+DB_PASSWORD, true)
    //sqlite
	//a, errcasbindb := gormadapter.NewAdapter("sqlite3", "/tmp/gorm.db", true)
    //mssql
	//a, errcasbindb := gormadapter.NewAdapter("mssql", "sqlserver://"+DB_USERNAME+":"+DB_PASSWORD+"@"+DB_HOST+":"+DB_PORT+"?database="+DB_DATABASE, true)
	//mongodb
    //a, errcasbindb := mongodbadapter.NewAdapter("mongodb://"+DB_USERNAME+":"+DB_PASSWORD+"@"+DB_HOST+":"+DB_PORT)


	if errcasbindb != nil {
		panic("Failed to connect to database!")
	  }

	  e, errcasbin := casbin.NewEnforcer("config/Casbin_role_model.conf", a)

	if errcasbin != nil {
		panic("Failed to casbin!")
	  }

	e.LoadPolicy()


	return e
}