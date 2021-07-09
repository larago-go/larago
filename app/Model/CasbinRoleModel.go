package Model

import (
	_ "fmt"
	//database_SQL
	"gorm.io/gorm"
	//end_database_SQL
	//MongoDB
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//end MongoDB
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

//MongoDB
//type CasbinRoleModel struct {
//primitive
//ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
//	ID           string  `json:"id" bson:"_id,omitempty"`
//	RoleName     string  `json:"v0" bson:"v0,omitempty"`
//	Path         string  `json:"v1" bson:"v1,omitempty"`
//	Method       string  `json:"v2" bson:"v2,omitempty"`

//}
//end_MongoDB

