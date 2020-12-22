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
type UserModel struct {
        gorm.Model
	ID           uint    `gorm:"column:id;primary_key"`
	Name         string  `gorm:"column:name;unique_index"`
	Email        string  `gorm:"column:email;unique_index"`
	Password     string  `gorm:"column:password;not null"`
//	//Casbinrole
	Role         string  `gorm:"column:role"`

} 

//func (UserModel) TableName() string {
//	return "custom_table_users"
// }
//end_database_SQL

//MongoDB
//type UserModel struct {
	//primitive
    //ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
//	ID           string  `json:"id" bson:"_id,omitempty"`
//	Name         string  `json:"name" bson:"name,omitempty"`
//	Email        string  `json:"email" bson:"email,omitempty"`
//	Password     string  `json:"password" bson:"password,omitempty"`
	//Casbinrole
//	Role         string  `json:"role" bson:"role,omitempty"`
//	Url         string  `json:"url" bson:"url,omitempty"`

//} 
//end MongoDB
