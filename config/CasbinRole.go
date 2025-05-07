package config

import (
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CasbinRole() *casbin.Enforcer {
	// Load environment variables
	errenv := godotenv.Load()
	if errenv != nil {
		panic("Error loading .env file")
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE") // Use the existing database name

	// Create a new adapter for the existing database, specifying the database name and dbSpecified
	dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_DATABASE + " port=" + DB_PORT + " sslmode=disable TimeZone=Europe/Moscow"
	db, errcasbindb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	gormadapter.TurnOffAutoMigrate(db)
	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, nil, "casbin_rule")
	if errcasbindb != nil {
		panic("Failed to connect databases: " + errcasbindb.Error())
	}

	if err != nil {
		panic("Failed to create adapter: " + err.Error())
	}
	// Create a new Casbin enforcer
	e, errcasbin := casbin.NewEnforcer("config/Casbin_role_model.conf", a)
	if errcasbin != nil {
		panic("Failed to create Casbin enforcer: " + errcasbin.Error())
	}

	// Load the policy from the database
	if err := e.LoadPolicy(); err != nil {
		panic("Failed to load policy: " + err.Error())
	}
	// Close the adapter when done
	defer a.Close()

	return e
}
