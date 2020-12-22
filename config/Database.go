package config

import (

	//mysql
	"gorm.io/driver/mysql"
	//postgres
	//"gorm.io/driver/postgres"
	//sqlite
	//"gorm.io/driver/sqlite"
	//mssql
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
	"github.com/joho/godotenv"
)


var DB *gorm.DB

func Init() {
	
	//env
	errenv := godotenv.Load()
	if errenv != nil {
		panic("Error loading .env file")
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE")

	//mysql
	dsn := DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//postgres
	//dsn := "user="+DB_USERNAME+" password="+DB_PASSWORD+" dbname="+DB_DATABASE+" port="+DB_PORT+" sslmode=disable TimeZone=Asia/Shanghai"
    //db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    //sqlite
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
    //mssql
	//dsn := "sqlserver://"+DB_USERNAME+":"+DB_PASSWORD+"@"+DB_HOST+":"+DB_PORT+"?database="+DB_DATABASE
    //db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

 // Migrations.
  
  db.AutoMigrate(&UserModel{})
  db.AutoMigrate(&CasbinRoleModel{})

 //End Migrations.


  DB = db 
}

// Run the migrations.

type UserModel struct {
	
	gorm.Model
	ID           uint    `gorm:"column:id;primary_key"`
	Name         string  `gorm:"column:name;unique_index"`
	Email        string  `gorm:"column:email;unique_index"`
	Password     string  `gorm:"column:password;not null"`
	//Casbinrole
	Role         string  `gorm:"column:role"`

}

//func (UserModel) TableName() string {
//	return "custom_table_users"
// }

type CasbinRoleModel struct {
	
        gorm.Model
	ID       uint    `gorm:"column:id;primary_key"`
	RoleName string  `gorm:"column:v0"`
	Path     string  `gorm:"column:v1"`
	Method   string  `gorm:"column:v2"`

}

//func (CasbinRoleModel) TableName() string {
//	return "casbin_rule"
// }
