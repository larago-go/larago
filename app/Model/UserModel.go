package Model

import (
	_ "fmt"
	//database_SQL
	"gorm.io/gorm"
	//end_database_SQL
)

//database_SQL
type UserModel struct {
	gorm.Model
	ID       uint   `gorm:"column:id;primary_key"`
	Name     string `gorm:"column:name;unique_index"`
	Email    string `gorm:"column:email;unique_index"`
	Password string `gorm:"column:password;not null"`
	//	//Casbinrole
	Role string `gorm:"column:role"`
}

//func (UserModel) TableName() string {
//	return "custom_table_users"
// }
//end_database_SQL
