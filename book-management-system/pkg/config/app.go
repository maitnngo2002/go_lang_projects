package config

// purpose of this file: return a db variable that helps us to talk to the database
import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "github_username:github_password/db_name?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}