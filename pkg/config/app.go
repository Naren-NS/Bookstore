package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db *gorm.DB
)

func Connect(){
	d,err := gorm.Open("mysql", "root:Naren@123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		panic(err)//used for aborting function if it returns an error that we donno how to handle
	}
	db=d
}//used to open a connection with our database 

func GetDB() *gorm.DB{
	return db
}


//Hence, app.go returns a db variable that helps to talk to the database.