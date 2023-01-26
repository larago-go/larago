package Model

import (
	_ "fmt"
	//database_SQL
	"gorm.io/gorm"
	//end_database_SQL
)

//database_SQL
type ResPassUserModel struct {
	gorm.Model
	ID       uint   `gorm:"column:id;primary_key"`
	Email    string `gorm:"column:email"`
	Url      string `gorm:"column:url"`
	Url_full string `gorm:"column:url_full"`
}

//func (ResPassUserModel) TableName() string {
//	return "custom_res_pass_users"
//}

//end_database_SQL
