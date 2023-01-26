package Model

import (
	_ "fmt"
	//database_SQL
	"gorm.io/gorm"
	//end_database_SQL
)

//database_SQL
type CasbinRoleModel struct {
	gorm.Model
	ID       uint   `gorm:"column:id;primary_key"`
	RoleName string `gorm:"column:v0"`
	Path     string `gorm:"column:v1"`
	Method   string `gorm:"column:v2"`
}

//func (CasbinRoleModel) TableName() string {
//	return "casbin_user_role"
// }
//end_database_SQL

